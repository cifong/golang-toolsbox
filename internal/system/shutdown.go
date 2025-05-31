package system

import (
	"errors"
	"os/exec"
	"runtime"
)

// ShutdownFunc allows injection for testing
var ShutdownFunc = shutdownImpl

func shutdownImpl(goos string, command func(name string, arg ...string) *exec.Cmd) error {
	if goos != "windows" {
		return errors.New("shutdown only supported on Windows")
	}
	cmd := command("shutdown", "/s", "/t", "0")
	return cmd.Run()
}

// Shutdown calls the injected ShutdownFunc with real dependencies
func Shutdown() error {
	return ShutdownFunc(runtime.GOOS, exec.Command)
}
