package pkg

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestSomething(t *testing.T) {
	if runtime.GOOS == "windows" {
		Something(filepath.Join("..", "test", "script.exe"))
		return
	}

	Something(filepath.Join("..", "test", "script"))
}
