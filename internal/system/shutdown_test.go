package system

import (
	"errors"
	"os/exec"
	"testing"
)

func TestShutdown_Windows(t *testing.T) {
	called := false
	ShutdownFunc = func(goos string, command func(string, ...string) *exec.Cmd) error {
		if goos != "windows" {
			t.Errorf("expected windows, got %s", goos)
		}
		cmd := command("shutdown", "/s", "/t", "0")
		want := []string{"shutdown", "/s", "/t", "0"}
		if len(cmd.Args) != len(want) {
			t.Errorf("unexpected command args length: got %d, want %d", len(cmd.Args), len(want))
		}
		for i, v := range want {
			if cmd.Args[i] != v {
				t.Errorf("unexpected command arg at %d: got %v, want %v", i, cmd.Args[i], v)
			}
		}
		called = true
		return nil
	}
	defer func() { ShutdownFunc = shutdownImpl }()
	if err := Shutdown(); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if !called {
		t.Error("ShutdownFunc was not called")
	}
}

func TestShutdown_NonWindows(t *testing.T) {
	ShutdownFunc = shutdownImpl
	if err := ShutdownFunc("linux", exec.Command); err == nil {
		t.Error("expected error for non-windows, got nil")
	} else if !errors.Is(err, errors.New("shutdown only supported on Windows")) && err.Error() != "shutdown only supported on Windows" {
		t.Errorf("unexpected error: %v", err)
	}
}
