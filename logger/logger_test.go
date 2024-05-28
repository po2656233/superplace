package logger

import (
	config2 "github.com/po2656233/superplace/config"
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	config := config2.DefaultLogConfig()
	config.EnableConsole = false
	config.EnableWriteFile = true
	config.FileLinkPath = "logs/log1.log"
	config.FilePathFormat = "logs/log1_%Y%m%d%H%M.log"

	log1 := NewConfigLogger(config)

	for i := 0; i < b.N; i++ {
		log1.Debug(1)
		log1.Info(2)
		log1.Error(3)
		log1.Panic(4)
	}
}
