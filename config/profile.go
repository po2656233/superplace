package config

import (
	"path/filepath"
	cerror "github/po2656233/superplace/logger/error"

	cfile "github/po2656233/superplace/extend/file"
	cjson "github/po2656233/superplace/extend/json"
	cstring "github/po2656233/superplace/extend/string"
	face "github/po2656233/superplace/facade"
)

var (
	cfg = &struct {
		profilePath string  // config root dir
		profileName string  // config name
		jsonConfig  *Config // config-x.json parse to json object
		env         string  // env name
		debug       bool    // debug default is true
		printLevel  string  // extend log print level
	}{}
)

func Path() string {
	return cfg.profilePath
}

func Name() string {
	return cfg.profileName
}

func Env() string {
	return cfg.env
}

func Debug() bool {
	return cfg.debug
}

func PrintLevel() string {
	return cfg.printLevel
}

func Init(filePath, nodeId string) (face.INode, error) {
	if filePath == "" {
		return nil, cerror.Error("File path is nil.")
	}

	if nodeId == "" {
		return nil, cerror.Error("NodeId is nil.")
	}

	judgePath, ok := cfile.JudgeFile(filePath)
	if !ok {
		return nil, cerror.Errorf("File path error. filePath = %s", filePath)
	}

	p, f := filepath.Split(judgePath)
	jsonConfig, err := loadFile(p, f)
	if err != nil || jsonConfig.Any == nil || jsonConfig.LastError() != nil {
		return nil, cerror.Errorf("Load config file error. [err = %v]", err)
	}

	node, err := GetNodeWithConfig(jsonConfig, nodeId)
	if err != nil {
		return nil, cerror.Errorf("Failed to get node config from config file. [err = %v]", err)
	}

	// init cfg
	cfg.profilePath = p
	cfg.profileName = f
	cfg.jsonConfig = jsonConfig
	cfg.env = jsonConfig.GetString("env", "default")
	cfg.debug = jsonConfig.GetBool("debug", true)
	cfg.printLevel = jsonConfig.GetString("print_level", "debug")

	return node, nil
}

func GetConfig(path ...interface{}) face.ProfileJSON {
	return cfg.jsonConfig.GetConfig(path...)
}

func loadFile(filePath, fileName string) (*Config, error) {
	// merge include json file
	var maps = make(map[string]interface{})

	// read master json file
	fileNamePath := filepath.Join(filePath, fileName)
	if err := cjson.ReadMaps(fileNamePath, maps); err != nil {
		return nil, err
	}

	// read include json file
	if v, found := maps["include"].([]interface{}); found {
		paths := cstring.ToStringSlice(v)
		for _, p := range paths {
			includePath := filepath.Join(filePath, p)
			if err := cjson.ReadMaps(includePath, maps); err != nil {
				return nil, err
			}
		}
	}

	return Wrap(maps), nil
}

//func judgeNameList(path, name string) ([]string, error) {
//	var list []string
//
//	if name != "" {
//		fileName := mergeProfileName(name)
//		list = append(list, fileName)
//
//	} else {
//		// find path
//		filesPath, err := cfile.ReadDir(path, "config-", ".json")
//		if err != nil {
//			return nil, err
//		}
//
//		if len(filesPath) < 1 {
//			return nil, cerror.Errorf("[path = %s] config file not found.", path)
//		}
//
//		for _, fp := range filesPath {
//			list = append(list, fp)
//		}
//	}
//
//	return list, nil
//}

//func mergeProfileName(name string) string {
//	return fmt.Sprintf("%s%s%s", profilePrefix, name, profileSuffix)
//}
