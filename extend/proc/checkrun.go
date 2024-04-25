package proc

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// CheckProcRunning 根据进程名判断进程是否运行
func CheckProcRunning(serverName string) (bool, error) {
	pid, err := RunCommand(serverName)
	if err != nil {
		return false, err
	}
	return pid != "", nil
}

// GetPid 根据进程名称获取进程ID
func GetPid(serverName string) (string, error) {
	pid, err := RunCommand(serverName)
	return pid, err
}

func RunCommand(cmd string) (string, error) {
	switch runtime.GOOS {
	case "windows":
		return runInWindows(cmd)
	case "linux":
		return runInWindows(cmd)
	default:
		return "", fmt.Errorf("unsupported operating system")
	}
}
func runInWindows(appName string) (string, error) {
	if !strings.Contains(appName, ".exe") {
		appName += ".exe"
	}
	cmd := exec.Command("cmd", "/C", "tasklist")
	output, _ := cmd.Output()
	//fmt.Printf("fields: %v\n", output)
	n := strings.Index(string(output), "System")
	if n == -1 {
		fmt.Println("no find")
		return "", fmt.Errorf("no find")
	}
	data := string(output)[n:]
	fields := strings.Fields(data)
	for k, v := range fields {
		if v == appName {
			return fields[k+1], nil
		}
	}

	return "", fmt.Errorf("no find")
}

func runInLinux(serverName string) (string, error) {
	cmd := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	fmt.Println("Running Linux cmd:" + cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}
