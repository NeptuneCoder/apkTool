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
func ExecDir(cmdStr string, args []string, dir func() string) {
	cmd := exec.Command(cmdStr, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir()
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmad.Output: ", err)
		return
	}
}
