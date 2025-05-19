package system

import (
	"errors"
	"os/exec"
	"runtime"
)

func Shutdown() error {
	if runtime.GOOS != "windows" {
		return errors.New("shutdown only supported on Windows")
	}

	// /s = 關機, /t 0 = 延遲 0 秒
	cmd := exec.Command("shutdown", "/s", "/t", "0")
	return cmd.Run()
}
