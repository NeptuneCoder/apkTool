package cmad

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec(cmdStr string, args []string) {
	cmd := exec.Command(cmdStr, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmad.Output: ", err)
		return
	}
}
