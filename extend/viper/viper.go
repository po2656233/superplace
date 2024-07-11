package exViper

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"io"
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

type SuperViper struct {
	*viper.Viper
	filePath string
	data     interface{}
}
type ConfData map[string]interface{}

func NewViper(filePath string) *SuperViper {
	viperCfg := viper.New()
	strDir, fileName := filepath.Split(filePath)
	if fileName == "" {
		panic(fmt.Errorf("no include file:%v", filePath))
	}
	if strDir == "" {
		strDir = "."
	}
	suffix := filepath.Ext(fileName)
	confName := strings.TrimSuffix(fileName, suffix)

	viperCfg.AddConfigPath(strDir)
	viperCfg.SetConfigName(confName)
	viperCfg.SetConfigFile(fileName)
	if 1 < len(suffix) {
		viperCfg.SetConfigType(suffix[1:])
	}

	sv := &SuperViper{
		Viper:    viperCfg,
		filePath: filePath,
		data:     make(map[string]interface{}),
	}
	err := sv.Viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = sv.Viper.Unmarshal(&sv.data)
	if err != nil {
		fmt.Printf("SuperViper Unmarshal err:%v", err)
	}

	sv.Viper.WatchConfig()
	sv.Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("OnConfigChange file changed:", e.Name)
		if err := sv.Viper.Unmarshal(&sv.data); err != nil {
			fmt.Println("OnConfigChange err:", err)
			return
		}
	})
	return sv
}

func (sv *SuperViper) Unmarshal(rawVal any, opts ...viper.DecoderConfigOption) error {
	if sv.Viper == nil {
		return fmt.Errorf("no superviper obj! ")
	}
	return sv.Viper.Unmarshal(rawVal, opts...)
}

func (sv *SuperViper) ToJson() error {
	return sv.ToFile(JSONSuffix, func(file io.Writer, info interface{}) error {
		return json.NewEncoder(file).Encode(info)
	})
}
func (sv *SuperViper) ToYaml() error {
	return sv.ToFile(YAMLSuffix, func(file io.Writer, info interface{}) error {
		return yaml.NewEncoder(file).Encode(info)
	})
}
func (sv *SuperViper) ToYml() error {
	return sv.ToFile(YMLSuffix, func(file io.Writer, info interface{}) error {
		return yaml.NewEncoder(file).Encode(info)
	})
}
func (sv *SuperViper) ToToml() error {
	return sv.ToFile(TOMLSuffix, func(file io.Writer, info interface{}) error {
		return toml.NewEncoder(file).Encode(info)
	})
}
func (sv *SuperViper) ToXml() error {
	return sv.ToFile(XMLSuffix, func(file io.Writer, info interface{}) error {
		return xml.NewEncoder(file).Encode(info)
	})
}
func (sv *SuperViper) ToIni() error {
	return sv.ToFile(INISuffix, func(file io.Writer, info interface{}) error {
		// 创建一个INI文件对象
		iniCfg := ini.Empty()
		switch info.(type) {
		case map[string]interface{}:
			if err := parseMap(iniCfg, info.(map[string]interface{}), ""); err != nil {
				return err
			}
		case string:
			iniCfg.Section("").NewKey(info.(string), fmt.Sprintf("%v", info))
		case map[interface{}]interface{}:
			if err := parseMap2(iniCfg, info.(map[interface{}]interface{}), ""); err != nil {
				return err
			}
		}

		// 将INI内容写入缓冲区
		_, err := iniCfg.WriteTo(file)
		return err
	})
}

func (sv *SuperViper) ToFile(ext string, f func(file io.Writer, info interface{}) error) error {
	extension := path.Ext(sv.filePath)
	if extension == ext {
		return nil
	}
	sv.filePath = strings.TrimSuffix(sv.filePath, extension) + ext
	file, err := os.Create(sv.filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	if err = f(file, sv.data); err != nil {
		log.Printf("%v encode[%#v] failed: %v!!!!!!!!!!", ext, sv.data, err)
		return err
	}

	return nil
}

// ///////////////////////////////////////

func parseMap2(cfg *ini.File, m map[interface{}]any, parentSection interface{}) error {
	for k, v := range m {
		switch v := v.(type) {
		case map[interface{}]interface{}:
			_ = parseMap2(cfg, v, parentSection)
		case map[string]interface{}:
			_ = parseMap(cfg, v, k.(string))
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
			_ = parseMap2(cfg, v, parentSection)
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
					_ = parseMap(cfg, valMap, parentSection)
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
