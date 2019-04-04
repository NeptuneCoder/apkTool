package jar2smali

import (
	"cmad"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup //定义一个同步等待的组
func Jar2Smali(apktoolPath, tempPath string) {
	_ = filepath.Walk(tempPath+"/lib", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".jar") {
			wg.Add(1)
			go func() {
				fmt.Println("name:", info.Name())
				newStr := strings.ReplaceAll(info.Name(), ".jar", ".dex")
				fmt.Println(newStr)
				newDexName := tempPath + "/lib/" + newStr
				outputPath := tempPath + "/lib"
				fmt.Println("outputPath:", outputPath)
				cmad.ExecDir("java", []string{"-jar", "-Xms512m", "-Xmx512m", "dx.jar", "--dex", "--output=" + newDexName, tempPath + "/lib/"}, func() string {
					return apktoolPath
				})
				os.RemoveAll(path)
				wg.Done()
			}()
		}
		return nil
	})
	wg.Wait()

}
