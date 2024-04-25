package proc

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// StartProcess 启动指定进程
func StartProcess(processPath, processArgs string) (pid int, err error) {
	var cmd *exec.Cmd
	if processArgs == "" {
		cmd = exec.Command(processPath)
	} else {
		cmd = exec.Command(processPath, processArgs)
	}
	if err = cmd.Start(); err != nil {
		return -1, err
	}
	return cmd.Process.Pid, nil
}

// StopProcess 根据进程名称终止进程
func StopProcess(processName string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		if !strings.Contains(processName, ".exe") {
			processName += ".exe"
		}
		cmd = exec.Command("taskkill", "/f", "/im", processName)

	case "linux":
		cmd = exec.Command("killall", processName)
	default:
		return fmt.Errorf("unsupported operating system")
	}
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to kill process %s: %v", processName, err)
	}
	return nil
}
