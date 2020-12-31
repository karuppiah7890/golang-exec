package pkg

import (
	"fmt"
	"os/exec"
	"strings"
)

// Something does something
func Something(scriptPath string) {
	cmd := exec.Command("bash", "-c", strings.Join([]string{scriptPath, "wow", "cool!"}, " "))
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
