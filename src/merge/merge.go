package merge

import (
	"dom4g"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"os"
	"path/filepath"
	"strings"
	"utils"
)

func MergeAndroidManifest(sdkPath, tempPath string) {

}

func AddSplashActivity(tempPath string, itemChannel *model.GameChannel) {
	RemoveDefualtSplashActivity(tempPath, itemChannel)
}

/*
	添加闪屏页面，如果已经存在，则需要删除之前的，添加新的
	添加闪屏页，则移除之前的启动页
 */
func RemoveDefualtSplashActivity(tempPath string, itemChannel *model.GameChannel) () {
	if itemChannel.Splash {
		xmlPath := tempPath + "/AndroidManifest.xml"
		buf, _ := ioutil.ReadFile(xmlPath)
		rootElement, _ := dom4g.LoadByXml(string(buf))
		var filter *dom4g.Element
		var main *dom4g.Element
		var launcher *dom4g.Element
		application := rootElement.Node("application")
		activitys := application.Nodes("activity")
		for _, v := range activitys {
			filter = v.Node("intent-filter")
			if filter != nil {
				v1, _ := filter.Node("action").AttrValue("name")
				v2, _ := filter.Node("category").AttrValue("name")
				if "android.intent.action.MAIN" == v1 && "android.intent.category.LAUNCHER" == v2 {
					main = filter.Node("action")
					launcher = filter.Node("category")
					b0 := filter.RemoveNode("action")
					b1 := filter.RemoveNode("category")
					v.RemoveNode("intent-filter")
					fmt.Println("移除启动的intent-filter", "b0 = ", b0, "   b1 = ", b1)
				}

				activityName, _ := v.AttrValue("name")
				fmt.Println("activityName:", activityName)
			}
			v2, _ := v.AttrValue("name")
			fmt.Println("activityName:", v2)
			if "com.foyoent.FoyoentSPlashActivity" == v2 {
				v.AddNode(main)
				v.AddNode(launcher)
			}

		}
		content := rootElement.SyncToXml()
		fmt.Println(content)
		ioutil.WriteFile(xmlPath, []byte(content), 0777)
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
func MergeIcon(sdkPath, tempPath string, gameChannel *model.GameChannel) {
	if gameChannel.IsIcon() {
		androidManifestName := tempPath + "/AndroidManifest.xml"
		content, err := ioutil.ReadFile(androidManifestName)
		if err != nil {
			log.Fatal(err)
		}
		rootElement, err := dom4g.LoadByXml(string(content))
		if err != nil {
			fmt.Println("err------------------------>:", err)
			return
		}
		appElement := rootElement.Node("application")
		iconType := "drawable"
		iconName := "drawable"
		if s, b := appElement.AttrValue("icon"); b {
			if "" != s && strings.Contains(s, "drawable") {
				iconName = strings.ReplaceAll(s, "@drawable/", "")
				iconType = "drawable"
			} else if "" != s && strings.Contains(s, "mipmap") {
				iconName = strings.ReplaceAll(s, "@mipmap/", "")
				iconType = "mipmap"
			}
		}
		fmt.Println("启动图标-类型：", iconType)
		fmt.Println("启动图标名称", iconName)

		//迭代图片目标文件，然后处理全部的icon图标
		//tempPath+"res"
		iteration(iconName, iconType, sdkPath, tempPath, gameChannel)

	}
}

/*
	迭代res目录下所有的存储图片的目录，找到目标图片，对其进行角标合并操作
 */

func iteration(iconName, iconType string, sdkPath, tempPath string, gameChannel *model.GameChannel) {
	resPath := tempPath + "/res"
	iconMarkPath := sdkPath + "/icon/"
	var size int
	_ = filepath.Walk(resPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if strings.Contains(info.Name(), "drawable-ldpi") {
				size = 32
				iconMarkPath += "drawable-ldpi" + "/" + gameChannel.Icon + ".png"
			} else if strings.Contains(info.Name(), "drawable-hdpi") {
				size = 72
				iconMarkPath += "drawable-hdpi" + "/" + gameChannel.Icon + ".png"
			} else if strings.Contains(info.Name(), "drawable-xhdpi") {
				size = 96
				iconMarkPath += "drawable-xhdpi" + "/" + gameChannel.Icon + ".png"
			} else if strings.Contains(info.Name(), "drawable-xxhdpi") {
				size = 144
				iconMarkPath += "drawable-xxhdpi" + "/" + gameChannel.Icon + ".png"
			} else if strings.Contains(info.Name(), "drawable-xxxhdpi") {
				size = 512
				iconMarkPath += "drawable-xxxhdpi" + "/" + gameChannel.Icon + ".png"
			} else if strings.Contains(info.Name(), "mipmap-ldpi") {
				size = 32
				iconMarkPath += "drawable-ldpi" + "/" + gameChannel.Icon + ".png"
			} else if strings.Contains(info.Name(), "mipmap-hdpi") {
				size = 72
				iconMarkPath += "drawable-hdpi" + "/" + gameChannel.Icon + ".png"
			} else if strings.Contains(info.Name(), "mipmap-xhdpi") {
				size = 96
				iconMarkPath += "drawable-xhdpi" + "/" + gameChannel.Icon + ".png"
			} else if strings.Contains(info.Name(), "mipmap-xxhdpi") {
				size = 144
				iconMarkPath += "drawable-xxhdpi" + "/" + gameChannel.Icon + ".png"
			} else if strings.Contains(info.Name(), "mipmap-xxxhdpi") {
				size = 512
				iconMarkPath += "drawable-xxxhdpi" + "/" + gameChannel.Icon + ".png"
			} else { // 默认为drawable-xxhdpi
				size = 144
				iconMarkPath += "drawable-xxhdpi" + "/" + gameChannel.Icon + ".png"
			}
		}
		return nil
	})

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

func RenamePackage(gameChannel *model.GameChannel, tempPath string) string {
	xmlPath := tempPath + "/AndroidManifest.xml"
	content, err := ioutil.ReadFile(xmlPath)
	newPageName := ""
	if err != nil {
		log.Fatal(err)
	}

	oldPackName := GetOldPackageName(tempPath)
	if gameChannel.PackageName == "" || gameChannel.Suffix == "" {
		newPageName = oldPackName
		return oldPackName
	}
	if gameChannel.PackageName != "" || gameChannel.Suffix == "" {
		newPageName = gameChannel.PackageName

	}
	if gameChannel.PackageName == "" || gameChannel.Suffix != "" {
		newPageName = oldPackName + gameChannel.Suffix
	}
	if gameChannel.PackageName != "" || gameChannel.Suffix != "" {
		newPageName = gameChannel.PackageName + gameChannel.Suffix
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
