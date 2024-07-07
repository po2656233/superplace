package exFile

import (
	"fmt"
	"testing"
)

func TestWalkFiles(t *testing.T) {
	files := WalkFiles("/config/", ".json")

	for _, file := range files {
		fmt.Println(file)
	}
}

func TestReadDir(t *testing.T) {
	files, err := ReadDir("/config/", "config-", ".json")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
