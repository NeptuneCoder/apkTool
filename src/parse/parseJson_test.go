package parse

import (
	"fmt"
	. "github.com/yanghai23/GoLib/atfile"
	"testing"
)

func TestReadSdkConfig(t *testing.T) {

	directory := GetCurrentDirectory()

	fmt.Println("........ReadSdkConfig", directory)
}

func TestReadGamsConfig(t *testing.T) {
	ReadGamesConfig()
}
