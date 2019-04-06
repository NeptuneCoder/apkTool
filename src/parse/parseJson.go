package parse

import (
	"encoding/json"
	"fmt"
	"github.com/yanghai23/GoLib/atfile"
	"model"
)


/**
	加载所有游戏列表
 */
func ReadGamesConfig() (*model.Games, error) {
	s := atfile.GetCurrentDirectory() + "/games/"
	data, err := atfile.ReadConfig(s, "games.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	games := new(model.Games)
	json.Unmarshal(data, &games)
	return games, nil
}

/**
	加载选中的游戏的渠道列表
 */
func ReadGameChannel(gamePath string) (*model.ChannelArray, error) {
	s := atfile.GetCurrentDirectory() + "/games/" + gamePath

	data, err := atfile.ReadConfig(s, "channel.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	channels := new(model.ChannelArray)
	json.Unmarshal(data, &channels)
	return channels, nil
}

/**
	加载选中的游戏的渠道列表
 */
func ReadSdkConfig(sdkPathName string) (*model.SdkRootConfig, error) {

	data, err := atfile.ReadConfig(sdkPathName, "config.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rootConfig := new(model.SdkRootConfig)
	json.Unmarshal(data, &rootConfig)
	return rootConfig, nil
}
