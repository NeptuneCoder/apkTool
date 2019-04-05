package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateOutDir() {
	os.MkdirAll("./", 0777)
}
func CopyJar(fromPath, toPath string) {
	_ = filepath.Walk(fromPath, func(path string, info os.FileInfo, err error) error {
		//检测目录正确性
		srcInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		if srcInfo.IsDir() {
			res, _ := filepath.Rel(fromPath, path)
			newLibPath := toPath + "/" + GetCurPath(res)
			CreateNewFolder(newLibPath)
		} else {
			res, _ := filepath.Rel(fromPath, path)
			//TODO copy
			newLibPath := toPath + "/" + res
			CopyFile(path, newLibPath)
		}
		return nil
	})
}
func GetCurPath(path string) string {
	if path == "" {
		return ""
	} else {
		return path + "/"
	}

}
func CopyBakToTemp(fromPath, toPath string, ReplaceStr func() string) {
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
			newPath := strings.ReplaceAll(res, ReplaceStr(), "")
			CreateNewFolder(toPath + newPath)
		} else {
			//TODO copy
			res, _ := filepath.Rel(toPath, path)
			newPath := strings.ReplaceAll(res, ReplaceStr(), "")
			CopyFile(path, toPath+newPath)
		}
		return nil
	})
}
func CopySmali(oldPackName, newPackName, fromPath, toPath string) {
	_ = filepath.Walk(fromPath, func(path string, info os.FileInfo, err error) error {
		//检测目录正确性
		srcInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		if !srcInfo.IsDir() {
			//TODO copy
			res, _ := filepath.Rel(fromPath, path)
			newFileName := toPath + "/" + res
			//TODO 这里分为两步，第一步是copy老的smali文件到新的目录中
			CopyFile(path, newFileName)
			content, err := ioutil.ReadFile(newFileName)
			if err != nil {
				fmt.Println("ReadFile Err:", err)
				return err
			}

			all := strings.ReplaceAll(string(content), oldPackName, newPackName)
			ioutil.WriteFile(newFileName, []byte(all), 0777)

		}
		return nil
	})
	os.RemoveAll(fromPath)
}

func CopyFile(srcName, targetName string) {
	file, err := os.Open(srcName)
	if err != nil {
		fmt.Println("Open File:", err)
	}
	defer file.Close()

	targetFile, err := os.Create(targetName)
	if err != nil {
		fmt.Println("CreateFile Err:", err)
		return
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
	fmt.Println("path:", path)
	boo, _ := PathExists(path)
	if boo {
		_ = os.RemoveAll(path)
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("create New Path err status:", err)
	}
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
