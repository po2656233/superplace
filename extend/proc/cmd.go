package proc

import (
	"os/exec"
	"path"
	"runtime"
)

// StartCMD 启动指定进程
func StartCMD(shFile string) error {
	var cmd *exec.Cmd
	shFile = GetScriptName(shFile)
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe", "/C", shFile)
	case "linux":
		cmd = exec.Command("/bin/bash", "-c", shFile)
	}
	return cmd.Start()
}

func GetScriptName(name string) string {
	ext := path.Ext(name)
	switch runtime.GOOS {
	case "windows":
		if ext != ".bat" {
			name += ".bat"
		}
	case "linux":
		if ext != ".sh" {
			name += ".sh"
		}
	}
	return name
}
