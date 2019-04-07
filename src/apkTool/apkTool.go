package main

import (
	"analysis"
	"cmad"
	"env"
	"fmt"
	"gameInfo"
	"github.com/yanghai23/GoLib/atfile"
	"model"
	"utils"
)

func init() {
	env.ReadEnvConfig()
	gameInfo.ReadGameConfig()
}
func main() {

	apkToolsPath := utils.GetDirRel() + "apktools/"
	fmt.Println("apkToolsPath:", apkToolsPath)
	//选中渠道名
	channelId := utils.KeyBordIn("请选择渠道(多个用逗号隔开，all为所有)：")
	channel := gameInfo.Channels.GetChannel(channelId)
	if i := channel[0]; i == nil {
		fmt.Println("没有找到您输入的渠道")
		return
	}
	workPath := utils.CreateNewFolder(utils.GetDirRel() + "work")
	fmt.Println("workPath:", workPath)
	//instanllFramework-res.apk
	installFrameworkRes(env.Env)
	//获取母包的路径
	gameApkPath, boo := utils.GetGameApkPath(gameInfo.Game.Folder)
	if !boo {
		fmt.Println("当前文件夹下面没有添加apk文件；请添加游戏母包apk!")
		return
	}
	fmt.Println("解压apk包")
	unZipApk(gameInfo.Game, gameApkPath)
	analysis.ExplainChannels(apkToolsPath, workPath, gameInfo.Game, channel)
}

func unZipApk(game *model.Game, apkPath string) {
	bakPath := atfile.GetCurrentDirectory() + "/" + "bak"
	cmad.Exec("java", []string{"-jar", "-Xms512m", "-Xmx512m", "./apktools/" + "apktool" + game.GetApktoolVersion() + ".jar", "d", "-o", bakPath, apkPath})

}
func installFrameworkRes(env *model.Environment) {
	cmad.Exec("java", []string{"-jar", "-Xms512m", "-Xmx512m", env.ApkToolPath, "if", env.FrameworkRes})
}
