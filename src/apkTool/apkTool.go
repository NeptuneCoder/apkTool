package main

import (
	"fmt"
	"github.com/yanghai23/GoLib/atfile"
	"model"
	"os"
	"os/exec"
	"package"
	"parse"
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
	//选中渠道名
	channelId := keybordIn("请选择渠道(多个用逗号隔开，all为所有)：")
	channel := channels.GetChannel(channelId)
	if i := channel[0]; i == nil {
		fmt.Println("没有找到您输入的渠道")
		return
	}
	_ = utils.CreateNewFolder(atfile.GetCurrentDirectory() + "/" + "work")
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

		rootConfig, _ := parse.ReadSdkConfig(itemChannel.Id)
		fmt.Println(rootConfig)

		fmt.Println("清空temp目录")
		tempPath := utils.CreateNewFolder(atfile.GetCurrentDirectory() + "/" + "work/temp")
		//复制原始的文件
		fmt.Println("复制原始文件到新的目录:", tempPath)
		fmt.Println(tempPath)
		utils.CopyBakToTemp(atfile.GetCurrentDirectory()+"/"+"bak", tempPath)
		fmt.Println("修改包名")
		pack.RenamePackage(itemChannel, tempPath)

	}

}

func unZipApk(game *model.Game, apkPath string) {
	bakPath := atfile.GetCurrentDirectory() + "/" + "bak"
	Exec("java", []string{"-jar", "-Xms512m", "-Xmx512m", "./apktools/" + "apktool" + game.GetApktoolVersion() + ".jar", "d", "-o", bakPath, apkPath})

}
func installFrameworkRes(env *model.Environment) {
	Exec("java", []string{"-jar", "-Xms512m", "-Xmx512m", env.ApkToolPath, "if", env.FrameworkRes})
}

func Exec(cmdStr string, args []string) {
	cmd := exec.Command(cmdStr, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}
}

func keybordIn(tip string) string {
	fmt.Println(tip)
	var gameId string
	fmt.Scanln(&gameId)
	return gameId
}
