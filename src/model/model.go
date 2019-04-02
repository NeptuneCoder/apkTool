package model

import (
	"fmt"
	"strconv"
	. "strings"
)

type Environment struct {
	ApkToolPath  string `json:apkToolPath`
	FrameworkRes string `json:frameworkRes`
}

/**
{
  "games": [
    {
      "name": "聚合Demo",
      "apktoolVersion": "2-1-0",
      "debug": "1",
      "folder": "foyoentdemo",
      "foyoentid": "1",
      "gameMainName": "",
      "id": 0,
      "orientation": "landscape"
    }
  ]
}
 */
type Game struct {
	Name           string `json:name`
	ApktoolVersion string `json:apktoolVersion`
	Debug          bool   `json:debug`
	Folder         string `json:folder`
	Foyoentid      string `json:foyoentid`
	GameMainName   string `json:gameMainName`
	Id             int    `json:id`
	Orientation    string `json:orientation`
}

func (game *Game) GetApktoolVersion() string {
	return game.ApktoolVersion
}

type Games struct {
	Games []Game `json:games`
}

func (games *Games) PrintlnAll() {
	for _, data := range games.Games {
		fmt.Println("GameName:", data.Name, "   gameId:", data.Id)
	}
}

func (games *Games) GetGameConfig(gameId string) (*Game) {
	for _, data := range games.Games {
		if strconv.Itoa(data.Id) == gameId {
			return &data
		}
	}
	return nil

}

/**
{
  "channels": [
    {
      "channel": {
        "name": "飓风",
        "icon": "rb",
        "id": "jufeng",
        "splash": "true",
        "param": [
          {
            "name": "FoyoentID",
            "value": "5008"
          },
          {
            "name": "FoyoentCID",
            "value": "1042"
          },
          {
            "name": "PAY_ACTION",
            "value": "tulongzhi.jufeng.MainActivity"
          }
        ],
        "meta-data": [
          {
            "name": "CID",
            "value": "5qwan"
          },
          {
            "name": "PID",
            "value": "5qwan"
          },
          {
            "name": "APPID",
            "value": "tulongzhi"
          },
          {
            "name": "APP_CONFIG",
            "value": "data/www/config"
          }
        ]
      }
    }
  ]
}
 */
type MetaData struct {
	Name  string `json:name`
	Value string `json:value`
}
type Param struct {
	Name  string `json:name`
	Value string `json:value`
}
type ChannelArray struct {
	Channels []*GameChannel `json:channels`
}

func (channelArray *ChannelArray) GetChannel(channelId string) []*GameChannel {
	if "all" == channelId {
		return channelArray.Channels
	} else {
		splits := Split(channelId, ",")
		byIds := make([]*GameChannel, len(splits))
		for i := 0; i < len(splits); i++ {
			byId := channelArray.GetChannelById(splits[i])
			byIds[i] = byId
		}

		return byIds
	}
}
func (channelArray *ChannelArray) GetChannelById(id string) *GameChannel {
	for _, v := range channelArray.Channels {
		if v.Id == id {
			return v
		}
	}
	return nil
}
func (channelArray *ChannelArray) PrintlnAll() {
	for _, v := range channelArray.Channels {
		fmt.Println("渠道名:", v.Name, "    渠道Id:", v.Id)
	}
}

type GameChannel struct {
	Name     string     `json:name`
	Icon     string     `json:icon`
	Id       string     `json:id`
	Splash   string     `json:splash`
	Param    []Param    `json:param`
	MetaData []MetaData `json:metaData`
}

/**
{
  "keystore": {
    "name": "sina.keystore",
    "password": "android",
    "alias": {
      "name": "androiddebugkey",
      "password": "android"
    }
  }
}
 */
type Alias struct {
	Name     string `json:name`
	Password string `json:password`
}
type StoreConfig struct {
	Name     string `json:name`
	Password string `json:password`
	Alias    Alias  `json:alias`
}
type KeyStoreConfig struct {
	Keystore StoreConfig `json:keystore`
}

/**
{
  "config": {
    "operations": [
      {
        "step": "1",
        "type": "merge",
        "from": "SDKManifest.xml",
        "to": "AndroidManifest.xml"
      },
      {
        "step": "2",
        "type": "copylib",
        "from": "libs",
        "to": "lib"
      },
      {
        "step": "3",
        "type": "copy",
        "from": "res",
        "to": "res"
      }
    ],
    "plugins": [
      {
        "name": "com.foyoent.jifeng.plugin.FoyoentJifengUser",
        "type": "1"
      },
      {
        "name": "com.foyoent.jifeng.plugin.FoyoentJifengPay",
        "type": "2"
      }
    ]
  }
}
 */
type SdkRootConfig struct {
	Config SdkConfig `json:config`
}
type SdkConfig struct {
	Application string      `json:application`
	Operations  []operation `json:operations`
	Plugins     []plugin    `json:plugins`
}

type plugin struct {
	Name string `json:name`
	Mold string `json:type`
}
type operation struct {
	Step string `json:step`
	Mold string `json:type`
	From string `json:from`
	To   string `json:to`
}
