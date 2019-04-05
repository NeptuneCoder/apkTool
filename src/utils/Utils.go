package utils

import (
	"fmt"
	"github.com/yanghai23/GoLib/atfile"
	"jar2smali"
	"merge"
	"model"
	"os"
	"parse"
	"path/filepath"
	"rjar"
	"strings"
)

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

func ExplainChannels(apkToolsPath, workPath string, game *model.Game, channels []*model.GameChannel) {
	for _, itemChannel := range channels {
		fmt.Println("开始打包，渠道【" + itemChannel.Id + "】")
		//读取sdk的配置信息
		sdkPath := atfile.GetCurrentDirectory() + "/config/sdk/" + itemChannel.Id
		sdkConfig, _ := parse.ReadSdkConfig(sdkPath)
		fmt.Println(sdkConfig)

		fmt.Println("清空temp目录")
		tempPath := CreateNewFolder(atfile.GetCurrentDirectory() + "/work/temp")
		//复制原始的文件
		fmt.Println("复制原始文件到新的目录:", tempPath)
		fmt.Println(tempPath)
		CopyBakToTemp(atfile.GetCurrentDirectory()+"/bak", tempPath, func() string {
			return "../../bak"
		})
		fmt.Println("修改包名")
		newPackageVal := merge.RenamePackage(itemChannel, tempPath)
		fmt.Println("NewPageName:", newPackageVal)
		fmt.Println("合并资源")
		// 将配置的jar，res，等资源进行合并
		merge.MergeSource(sdkPath, tempPath, sdkConfig.Config.Operations)

		//是否添加闪屏页面

		fmt.Println("添加闪屏图片")
		merge.AddSplashImg(sdkPath, tempPath, itemChannel, game)

		//处理Icon图标

		merge.MergeIcon(sdkPath, tempPath, itemChannel)
		fmt.Println("生成R文件")
		rjar.ComplieR(apkToolsPath, tempPath, workPath, newPackageVal, &sdkConfig.Config)
		fmt.Println("jar2smali")
		jar2smali.Jar2Smali(apkToolsPath, tempPath)
		fmt.Println("合并meta-data")
		merge.MergeMetaData(tempPath, itemChannel)

		merge.AddSplashActivity(tempPath, itemChannel)
	}
}
