package viper

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	JSONSuffix = ".json"
	XMLSuffix  = ".xml"
	YAMLSuffix = ".yaml"
	YMLSuffix  = ".yml"
	TOMLSuffix = ".toml"
	INISuffix  = ".ini"
)

type ConfData map[string]interface{}

var data = make(map[string]ConfData)
var curPath = ""

func InitViper(filePath string) {
	viperCfg := viper.New()

	strDir, fileName := filepath.Split(filePath)
	if fileName == "" {
		panic(fmt.Errorf("no include file:%v", filePath))
	}
	if strDir == "" {
		strDir = "."
	}
	viperCfg.AddConfigPath(strDir)

	suffix := filepath.Ext(fileName)
	confName := strings.TrimSuffix(fileName, suffix)
	viperCfg.SetConfigName(confName)
	if 1 < len(suffix) {
		viperCfg.SetConfigType(suffix[1:])
	}

	err := viperCfg.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var confData ConfData
	err = viperCfg.Unmarshal(&confData)
	if err != nil {
		panic(err)
	}

	viperCfg.WatchConfig()
	viperCfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("OnConfigChange file changed:", e.Name)
		if err = viperCfg.Unmarshal(&confData); err != nil {
			fmt.Println("OnConfigChange err:", err)
			return
		}
		data[e.Name] = confData
	})
	data[filePath] = confData
	curPath = filePath
}

func GetCurConfig() ConfData {
	return GetViperConfig(curPath)
}

func GetViperConfig(filePath string) ConfData {
	info, ok := data[filePath]
	if !ok {
		return nil
	}
	return info
}

func ToJson() error {
	return toFile(JSONSuffix, func(fileName string, info ConfData) error {
		file, err := os.Create(fileName)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		err = json.NewEncoder(file).Encode(info)
		_ = file.Close()
		return err
	})
}
func ToYaml() error {
	return toFile(YAMLSuffix, func(fileName string, info ConfData) error {
		file, err := os.Create(fileName)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		err = yaml.NewEncoder(file).Encode(info)
		_ = file.Close()
		return err
	})
}
func ToYml() error {
	return toFile(YMLSuffix, func(fileName string, info ConfData) error {
		file, err := os.Create(fileName)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		err = yaml.NewEncoder(file).Encode(info)
		_ = file.Close()
		return err
	})
}
func ToToml() error {
	return toFile(TOMLSuffix, func(fileName string, info ConfData) error {
		file, err := os.Create(fileName)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		err = toml.NewEncoder(file).Encode(info)
		if err != nil {
			return err
		}
		_ = file.Close()
		return err
	})
}
func ToXml() error {
	return toFile(XMLSuffix, func(fileName string, info ConfData) error {
		file, err := os.Create(fileName)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		err = xml.NewEncoder(file).Encode(info)
		if err != nil {
			return err
		}
		_ = file.Close()
		return err
	})
}
func ToIni() error {
	return toFile(INISuffix, func(fileName string, info ConfData) error {
		// 创建一个INI文件对象
		iniCfg := ini.Empty()
		if err := parseMap(iniCfg, info, ""); err != nil {
			return err
		}
		// 将INI内容写入缓冲区
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		if _, err = iniCfg.WriteTo(file); err != nil {
			return err
		}
		return file.Close()
	})
}
func toFile(ext string, f func(fileName string, info ConfData) error) error {
	for fileName, configData := range data {
		// TOMLEncodeIntoFile 将 v 对象编码成 TOML 格式配置并写入 filepath 文件
		extension := path.Ext(fileName)
		if extension == ext {
			continue
		}
		fileName = strings.TrimSuffix(fileName, extension) + ext
		err := f(fileName, configData)
		if err != nil {
			log.Printf("%v encode[%#v] failed: %v!!!!!!!!!!", ext, configData, err)
			return err
		}
	}

	return nil
}

// ///////////////////////////////////////

func parseMap2(cfg *ini.File, m map[interface{}]any, parentSection interface{}) error {
	for k, v := range m {
		switch v := v.(type) {
		case map[interface{}]interface{}:
			parseMap2(cfg, v, parentSection)
		case map[string]interface{}:
			parseMap(cfg, v, k.(string))
		default:
			// 否则，处理为键值对
			key, ok := k.(string)
			if !ok {
				return fmt.Errorf("key is not string")
			}
			section, ok1 := parentSection.(string)
			if !ok1 {
				return fmt.Errorf("parentSection is not string")
			}

			if v == nil {
				break
			}
			_, err := cfg.Section(section).NewKey(key, fmt.Sprintf("%v", v))
			if err != nil {
				return err
			}

		}
	}
	return nil
}
func parseMap(cfg *ini.File, m map[string]interface{}, parentSection string) error {
	for k, v := range m {
		switch v := v.(type) {
		case map[interface{}]any:
			parseMap2(cfg, v, parentSection)
		case map[string]interface{}:
			// 如果是嵌套的map，则处理为新的section
			sectionName := k
			if parentSection != "" {
				sectionName = parentSection + "." + sectionName
			}
			_, err := cfg.NewSection(sectionName)
			if err != nil {
				return err
			}
			err = parseMap(cfg, v, sectionName)
			if err != nil {
				return err
			}

		default:
			// 否则，处理为键值对
			key := k
			//if parentSection != "" {
			//	key = parentSection + "." + key
			//}
			if v == nil {
				break
			}
			if vals, ok := v.([]interface{}); ok {
				for _, i2 := range vals {
					valMap := make(map[string]interface{})
					valMap[key] = i2
					parseMap(cfg, valMap, parentSection)
				}
			} else {
				_, err := cfg.Section(parentSection).NewKey(key, fmt.Sprintf("%v", v))
				if err != nil {
					return err
				}
			}

		}
	}
	return nil
}
