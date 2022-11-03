package ctool

import (
	"bytes"
	"os/exec"
	"runtime"
)

//IsWindows 判断当前系统是否为windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

//IsLinux 判断当前系统是否为linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

//IsMac 判断当前系统是否为mac
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// ExecCommand 执行/bin/sh -c "command"
func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindows() {
		cmd = exec.Command("cmd")
	}
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err = cmd.Run()

	if err != nil {
		stderr = string(errOut.Bytes())
	}
	stdout = string(out.Bytes())

	return
}

//GetOsBits 获取当前系统的位数
func GetOsBits() int {
	return 32 << (^uint(0) >> 63)
}
