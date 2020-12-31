package pkg

import (
	"path/filepath"
	"testing"
)

func TestSomething(t *testing.T) {
	Something(filepath.Join("..", "test", "script"))
}
