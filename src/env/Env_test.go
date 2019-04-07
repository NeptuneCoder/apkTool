package env

import (
	"fmt"
	"testing"
)

func TestReadEnvConfig(t *testing.T) {

	ReadEnvConfig()
	fmt.Println("evn:", Env.AndroidJarName)
}
