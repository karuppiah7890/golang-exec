package pkg

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestSomething(t *testing.T) {
	if runtime.GOOS != "windows" {
		Something(filepath.Join("..", "test", "script"))
		return
	}

	scriptPath, err := filepath.Abs(filepath.Join("..", "test", "script.bat"))
	if err != nil {
		panic(err)
	}
	Something(scriptPath)
}
