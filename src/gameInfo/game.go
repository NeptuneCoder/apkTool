package gameInfo

import (
	"model"
	"parse"
	"utils"
)

var Channels *model.ChannelArray
var Game *model.Game

func ReadGameConfig() {
	gamesConfig, _ := parse.ReadGamesConfig()
	gamesConfig.PrintlnAll()
	gameId := utils.KeyBordIn("请选择一个游戏(输入gameId)：")
	Game = gamesConfig.GetGameConfig(gameId)

	for Game == nil {
		gameId := utils.KeyBordIn("您输入的gameId有误，请从输：")
		Game = gamesConfig.GetGameConfig(gameId)
	}
	Channels, _ = parse.ReadGameChannel(Game.Folder)
	Channels.PrintlnAll()
}
