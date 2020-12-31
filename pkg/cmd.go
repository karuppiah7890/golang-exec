package pkg

import (
	"fmt"
	"os/exec"
)

// Something does something
func Something(scriptPath string) {
	cmd := exec.Command(scriptPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
