package config

import (
	"fmt"
	"github.com/po2656233/superplace/extend/viper"
	"testing"
)

func TestLoadFile(t *testing.T) {
	path := "profile-gc.json"
	node, err := Init(path, "game-1")
	fmt.Println(node, err)
	viper.InitViper(path)
	viper.ToJson()
	viper.ToToml()
	//ToXml()
	viper.ToYaml()
	viper.ToYml()
	viper.ToIni()

}
