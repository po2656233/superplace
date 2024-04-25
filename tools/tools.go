package tools

import (
	exFile "github.com/po2656233/superplace/extend/file"
	"path/filepath"
)

func GetNatsSHFile() string {
	abs := exFile.GetCurrentPath()
	return filepath.Join(abs, "nats-server", exFile.GetScriptName("run_nats"))
}
