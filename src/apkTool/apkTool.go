package main

import (
	"cmad"
	"fmt"
	"github.com/yanghai23/GoLib/atfile"
	"jar2smali"
	"merge"
	"model"
	"parse"
	"rjar"
	"time"
	"utils"
)

func main() {
	env, _ := parse.ReadEnvConfig()
	gamesConfig, _ := parse.ReadGamesConfig()
	gamesConfig.PrintlnAll()
	gameId := keybordIn("请选择一个游戏(输入gameId)：")
	game := gamesConfig.GetGameConfig(gameId)

	for game == nil {
		gameId := keybordIn("您输入的gameId有误，请从输：")
		game = gamesConfig.GetGameConfig(gameId)
	}
	channels, _ := parse.ReadGameChannel(game.Folder)
	channels.PrintlnAll()
	apkToolsPath := atfile.GetCurrentDirectory() + "/apktools/"
	fmt.Println("apkToolsPath:", apkToolsPath)
	//选中渠道名
	channelId := keybordIn("请选择渠道(多个用逗号隔开，all为所有)：")
	channel := channels.GetChannel(channelId)
	if i := channel[0]; i == nil {
		fmt.Println("没有找到您输入的渠道")
		return
	}
	workPath := utils.CreateNewFolder(atfile.GetCurrentDirectory() + "/" + "work")
	fmt.Println("workPath:", workPath)
	//instanllFramework-res.apk
	installFrameworkRes(env)
	//获取母包的路径
	gameApkPath, boo := utils.GetGameApkPath(game.Folder)
	if !boo {
		fmt.Println("当前文件夹下面没有添加apk文件；请添加游戏母包apk!")
		return
	}
	fmt.Println("解压apk包")
	unZipApk(game, gameApkPath)

	for _, itemChannel := range channel {
		fmt.Println("开始打包，渠道【" + itemChannel.Id + "】")
		//读取sdk的配置信息
		sdkPath := atfile.GetCurrentDirectory() + "/config/sdk/" + itemChannel.Id
		sdkConfig, _ := parse.ReadSdkConfig(sdkPath)
		fmt.Println(sdkConfig)

		fmt.Println("清空temp目录")
		tempPath := utils.CreateNewFolder(atfile.GetCurrentDirectory() + "/work/temp")
		//复制原始的文件
		fmt.Println("复制原始文件到新的目录:", tempPath)
		fmt.Println(tempPath)
		utils.CopyBakToTemp(atfile.GetCurrentDirectory()+"/bak", tempPath, func() string {
			return "../../bak"
		})
		fmt.Println("修改包名")
		newPackageVal := merge.RenamePackage(itemChannel, tempPath)
		fmt.Println("NewPageName:", newPackageVal)
		fmt.Println("合并资源")
		// 将配置的jar，res，等资源进行合并
		if len(sdkConfig.Config.Operations) != 0 {
			merge.MergeSource(sdkPath, tempPath, sdkConfig.Config.Operations)
		}

		//是否添加闪屏页面
		if itemChannel.Splash {
			fmt.Println("添加闪屏图片")
			merge.AddSplashImg(sdkPath, tempPath, itemChannel, game)
		}

		//处理Icon图标
		if itemChannel.IsIcon() {
			merge.MergeIcon(sdkPath, tempPath, itemChannel)
		}
		fmt.Println("生成R文件")
		rjar.ComplieR(apkToolsPath, tempPath, workPath, newPackageVal, &sdkConfig.Config)
		fmt.Println("jar2smali")
		t6 := time.Now().Unix() //秒
		jar2smali.Jar2Smali(apkToolsPath, tempPath)
		fmt.Println("花费的时间:", time.Now().Unix() - t6, "s")

		fmt.Println("合并meta-data")
		merge.MergeMetaData(tempPath, itemChannel)
	}

}

func unZipApk(game *model.Game, apkPath string) {
	bakPath := atfile.GetCurrentDirectory() + "/" + "bak"
	cmad.Exec("java", []string{"-jar", "-Xms512m", "-Xmx512m", "./apktools/" + "apktool" + game.GetApktoolVersion() + ".jar", "d", "-o", bakPath, apkPath})

}
func installFrameworkRes(env *model.Environment) {
	cmad.Exec("java", []string{"-jar", "-Xms512m", "-Xmx512m", env.ApkToolPath, "if", env.FrameworkRes})
}

func keybordIn(tip string) string {
	fmt.Println(tip)
	var gameId string
	fmt.Scanln(&gameId)
	return gameId
}
