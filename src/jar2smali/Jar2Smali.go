package jar2smali

import (
	"os"
	"path/filepath"
	"strings"
)

func Jar2Smali(apktoolPath, tempPath string) {
	_ = filepath.Walk(tempPath+"/lib", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".jar") {

		}
		return nil
	})

}
