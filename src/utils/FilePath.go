package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CopyBakToTemp(fromPath, toPath string) {
	_ = filepath.Walk(fromPath, func(path string, info os.FileInfo, err error) error {
		//检测目录正确性
		srcInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		if srcInfo.IsDir() {
			//TODO  创建目录
			res, _ := filepath.Rel(toPath, path)
			newPath := strings.ReplaceAll(res, "../../bak", "")
			CreateNewFolder(toPath + newPath)
		} else {
			//TODO copy
			res, _ := filepath.Rel(toPath, path)
			newPath := strings.ReplaceAll(res, "../../bak", "")

			CopyFile(path, toPath+newPath)
		}

		return nil
	})
}

func CopyFile(srcName, targetName string) {
	file, err := os.Open(srcName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	targetFile, err := os.Create(targetName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buf := make([]byte, 1024) //创建一个初始容量为1024的slice,作为缓冲容器
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		targetFile.Write(buf[:n])
	}

}

/**
	创建一个工作目录
 */
func CreateNewFolder(path string) string {
	boo, _ := PathExists(path)
	if boo {
		_ = os.RemoveAll(path)
	}
	_ = os.Mkdir(path, os.ModePerm)
	return path
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
