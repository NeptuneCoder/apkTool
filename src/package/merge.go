package pack

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"strings"
	"utils"
)

func MergeIcon(sdkPath, tempPath string, itemChannel *model.GameChannel) {

}

func AddSplashImg(sdkPath, tempPath string, channel *model.GameChannel, game *model.Game) {
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

func ExecuteOperation(sdkPath, tempPath string, operations []model.Operation) {
	fmt.Println("sdkPath:", sdkPath)
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
	fmt.Println("content :", oldPackName)
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
	all := strings.ReplaceAll(string(content), oldPackName, newPageName)
	ioutil.WriteFile(xmlPath, []byte(all), 0777)

	smaliPath := tempPath + "/smali"
	oldPath := strings.ReplaceAll(oldPackName, ".", "/")
	newPath := strings.ReplaceAll(newPageName, ".", "/")

	if oldPath != newPath {
		oldSmaliPath := smaliPath + "/" + oldPath
		newSmaliPath := smaliPath + "/" + newPath
		fmt.Println("oldSmaliPath", oldSmaliPath)

		fmt.Println("newSmaliPath", newSmaliPath)
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
	decoder := xml.NewDecoder(bytes.NewBuffer(content))
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				if "package" == attrName {
					return attr.Value
				}

			}
		// 处理元素结束（标签）
		case xml.EndElement:
			fmt.Printf("Token of '%s' end\n", token.Name.Local)
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
			// ...
		}

	}
	return ""
}
