package rjar

import (
	"cmad"
	"env"
	"fmt"
	"io/ioutil"
	"model"
	"os"
	"os/exec"
	"strings"
	"utils"
)

func GetAapt() string {
	return env.Env.AaptVersion
}
func GetAndroidJarName() string {
	return env.Env.AndroidJarName
}
func ComplieR(apkToolsPath, tempPath, workPath, newPackageVal string, sdkConfig *model.SdkConfig) {
	aaptPath := apkToolsPath + GetAapt()
	resPath := tempPath + "/res"
	androidJarPath := apkToolsPath + GetAndroidJarName()
	manifestPath := tempPath + "/AndroidManifest.xml"
	rClazzPath := workPath + "/r"
	utils.CreateNewFolder(rClazzPath)
	GenerateR(aaptPath, rClazzPath, resPath, androidJarPath, manifestPath)
	//todo 编译R.java文件
	packagePath := rClazzPath + "/" + strings.ReplaceAll(newPackageVal, ".", "/")
	RFilePath := packagePath + "/R.java"
	fmt.Println("RFilePath:", RFilePath)
	CopyR2NewPkg(newPackageVal, RFilePath, rClazzPath, sdkConfig.RJars)
	CompileR(RFilePath)
	GenerateRClazz2Jar(tempPath, rClazzPath)
	os.RemoveAll(rClazzPath)
}

func CopyR2NewPkg(newPackageVal, sourceR, rPath string, jars []model.RJar) {

	if len(jars) == 0 {
		return
	}

	for _, itemVal := range jars {
		newPkgPath := strings.ReplaceAll(itemVal.Name, ".", "/")
		resPath := rPath + "/" + newPkgPath
		os.MkdirAll(resPath, 0777)
		mRName := resPath + "/" + "R.java"
		utils.CopyFile(sourceR, mRName) //copy完成后

		//TODO   替换老的包名
		content, err := ioutil.ReadFile(mRName)
		if err != nil {
			fmt.Println("ReadFile Err:", err)
			return
		}
		all := strings.ReplaceAll(string(content), newPackageVal, itemVal.Name)
		ioutil.WriteFile(mRName, []byte(all), 0777)
		CompileR(mRName)
	}
}

func GenerateRClazz2Jar(tempPath, rClazzPath string) {
	libPath := tempPath + "/lib"
	jarPath := libPath + "/R.jar"
	cmd := exec.Command("jar", "-cvf", jarPath, "./")
	cmd.Dir = rClazzPath + "/"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmad.Output: ", err)
		return
	}
}

func CompileR(javaPath string) {
	cmad.Exec("javac", []string{"-source", "1.7", "-target", "1.7", "-encoding", "UTF-8", javaPath})
	os.RemoveAll(javaPath)
}

func GenerateR(aaptPath, rPath, resPath, androidJarPath, manifestPath string) {
	cmad.Exec(aaptPath, []string{"p", "-f", "-m", "-J", rPath, "-S", resPath, "-I", androidJarPath, "-M", manifestPath})
}
