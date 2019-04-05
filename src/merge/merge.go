package merge

import (
	"dom4g"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"strings"
	"utils"
)

func MergeAndroidManifest(sdkPath ,tempPath string) {

}
func AddSplashActivity(tempPath string, itemChannel *model.GameChannel) () {
	if itemChannel.Splash {

	}
}
func MergeMetaData(tempPath string, channel *model.GameChannel) {
	xmlPath := tempPath + "/" + "AndroidManifest.xml"
	buf, _ := ioutil.ReadFile(xmlPath)
	fmt.Println("buf:", string(buf))
	rootElement, _ := dom4g.LoadByXml(string(buf))
	appNode := rootElement.Node("application")
	for _, data := range channel.MetaData {
		metaData := dom4g.NewElement("meta-data", "")
		metaData.AddAttr("android:name", data.Name)
		metaData.AddAttr("android:value", data.Value)
		appNode.AddNode(metaData)
	}
	ioutil.WriteFile(xmlPath, []byte(rootElement.SyncToXml()), 0777)

}

//合并icon
func MergeIcon(sdkPath, tempPath string, itemChannel *model.GameChannel) {
	if itemChannel.IsIcon() {

	}
}

func AddSplashImg(sdkPath, tempPath string, channel *model.GameChannel, game *model.Game) {
	if !channel.Splash {
		return
	}
	splashName := channel.GetSplashImgName(game)
	splashImgPath := sdkPath + "/" + "splash" + "/" + splashName + ".png"
	fmt.Println("splashImgPath:", splashImgPath)
	fmt.Println("tempPath:", tempPath)
	boo, _ := utils.PathExists(splashImgPath)
	if !boo {
		fmt.Println("该渠道未配置闪屏图片")
		return
	}
	drawablePath := tempPath + "/res/" + "drawable-hdpi"
	boo, _ = utils.PathExists(splashImgPath)
	if !boo {
		utils.CreateNewFolder(drawablePath)
	}
	fmt.Println("splashImgPath  : ", splashImgPath)
	fmt.Println("drawablePath  : ", drawablePath)
	utils.CopyFile(splashImgPath, drawablePath+"/ic_launcher.png")
}

func MergeSource(sdkPath, tempPath string, operations []model.Operation) {
	for _, operation := range operations {
		if "replaceManifest" == operation.Mold {
			manifestPath := tempPath + "/" + "AndroidManifest.xml"
			utils.CopyFile(operation.From, manifestPath)
		} else if "copylib" == operation.Mold {
			fromPath := sdkPath + "/" + operation.From
			toPath := tempPath + "/" + operation.To
			fmt.Println("fromPath:", fromPath)
			fmt.Println("toPath:", toPath)
			utils.CreateNewFolder(toPath)
			utils.CopyJar(fromPath, toPath)
		} else {
			fromPath := sdkPath + "/" + operation.From
			toPath := tempPath + "/" + operation.To
			utils.CopyJar(fromPath, toPath)
		}
	}
}

func RenamePackage(gchannel *model.GameChannel, tempPath string) string {
	xmlPath := tempPath + "/AndroidManifest.xml"
	content, err := ioutil.ReadFile(xmlPath)
	newPageName := ""
	if err != nil {
		log.Fatal(err)
	}

	oldPackName := GetOldPackageName(tempPath)
	if gchannel.PackageName == "" || gchannel.Suffix == "" {
		newPageName = oldPackName
		return oldPackName
	}
	if gchannel.PackageName != "" || gchannel.Suffix == "" {
		newPageName = gchannel.PackageName

	}
	if gchannel.PackageName == "" || gchannel.Suffix != "" {
		newPageName = oldPackName + gchannel.Suffix
	}
	if gchannel.PackageName != "" || gchannel.Suffix != "" {
		newPageName = gchannel.PackageName + gchannel.Suffix
	}
	fmt.Println("xmlPath:", xmlPath)
	fmt.Println("Manifest.xml:", string(content))
	all := strings.ReplaceAll(string(content), oldPackName, newPageName)
	ioutil.WriteFile(xmlPath, []byte(all), 0777)

	smaliPath := tempPath + "/smali"
	oldPath := strings.ReplaceAll(oldPackName, ".", "/")
	newPath := strings.ReplaceAll(newPageName, ".", "/")

	if oldPath != newPath {
		oldSmaliPath := smaliPath + "/" + oldPath
		newSmaliPath := smaliPath + "/" + newPath
		//创建新的包名的路径
		utils.CreateNewFolder(newSmaliPath)
		//在smali文件中，包路径全部的写法是com/example/demo
		utils.CopySmali(oldPath, newPath, oldSmaliPath, newSmaliPath)
	}
	return newPageName
}

func GetOldPackageName(tempPath string) string {
	xmlPath := tempPath + "/AndroidManifest.xml"
	content, err := ioutil.ReadFile(xmlPath)
	if err != nil {
		log.Fatal(err)
	}
	rootElement, _ := dom4g.LoadByXml(string(content))
	s, _ := rootElement.AttrValue("package")
	return s
}
