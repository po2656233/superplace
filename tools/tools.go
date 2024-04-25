package tools

import (
	exFile "github.com/po2656233/superplace/extend/file"
	"github.com/po2656233/superplace/extend/proc"
	"path/filepath"
)

func GetNatsSHFile() string {
	abs := exFile.GetCurrentPath()
	return filepath.Join(abs, "nats-server", proc.GetScriptName("run_nats"))
}
