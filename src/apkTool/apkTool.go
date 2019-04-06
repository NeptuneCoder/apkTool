package main

import (
	"cmad"
	"env"
	"fmt"
	"github.com/yanghai23/GoLib/atfile"
	"model"
	"parse"
	"utils"
)

func init() {
	env.ReadEnvConfig()
}
func main() {

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
	installFrameworkRes(env.Env)
	//获取母包的路径
	gameApkPath, boo := utils.GetGameApkPath(game.Folder)
	if !boo {
		fmt.Println("当前文件夹下面没有添加apk文件；请添加游戏母包apk!")
		return
	}
	fmt.Println("解压apk包")
	unZipApk(game, gameApkPath)
	utils.ExplainChannels(apkToolsPath, workPath, game, channel)
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
