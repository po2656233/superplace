package config

import (
	"fmt"
	"testing"
)

func TestLoadFile(t *testing.T) {
	path := "config/config-dev.json"
	node, err := Init(path, "game-1")
	fmt.Println(node, err)
	InitViper(path)
}
