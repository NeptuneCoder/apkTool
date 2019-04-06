package utils

import (
	"fmt"
	"github.com/yanghai23/GoLib/atfile"
	"os"
	"path/filepath"
	"strings"
)

func KeyBordIn(tip string) string {
	fmt.Println(tip)
	var gameId string
	fmt.Scanln(&gameId)
	return gameId
}
func GetGameApkPath(path string) (resPath string, boo bool) {
	rootPath := atfile.GetCurrentDirectory() + "/games/" + path
	_ = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), "apk") {
			fmt.Println("path = ", path, "   info.Name() :", info.Name())
			resPath = path
			boo = true
			fmt.Println("resPath------true")
			return nil
		}
		return nil
	})
	fmt.Println("resPath:", resPath, "boo  = ", boo)
	return resPath, boo
}
