package main

import (
	"fmt"
	"github.com/yanghai23/GoLib/atfile"
	"model"
	"os"
	"os/exec"
	"parse"
	"utils"
)

func main() {
	env, _ := parse.ReadEnvConfig()
	gamesConfig, _ := parse.ReadGamsConfig()
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
	createNewFolder("work")
	//instanllFramework-res.apk
	installFrameworkRes(env)
	//获取母包的路径
	gameApkPath, boo := utils.GetGameApkPath(game.Folder)
	if !boo {
		fmt.Println("当前文件夹下面没有添加apk文件；请添加游戏母包apk!")
		return
	}
	unZipApk(game, gameApkPath)

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

/**
	创建一个工作目录
 */
func createNewFolder(workPath string) {
	path := atfile.GetCurrentDirectory() + "/" + workPath
	boo, _ := PathExists(path)
	if boo {
		_ = os.RemoveAll(path)
	}
	_ = os.Mkdir(path, os.ModePerm)
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

func keybordIn(tip string) string {
	fmt.Println(tip)
	var gameId string
	fmt.Scanln(&gameId)
	return gameId
}
