package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Environment struct {
	ApkToolPath    string `json:apkToolPath`
	FrameworkRes   string `json:frameworkRes`
	AaptVersion    string `json:aaptVersion`
	AndroidJarName string `json:androidJarName`
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
            "value": "data/www/customConfig"
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
		splits := strings.Split(channelId, ",")
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
	PackageName   string     `json:packageName`
	Keystore      string     `json:keystore`
	Suffix        string     `json:suffix` //包名后缀
	Name          string     `json:name`
	Icon          string     `json:icon`
	Id            string     `json:id` //渠道号
	Splash        bool       `json:splash`
	SplashImgName string     `json:splashImgName`
	Param         []Param    `json:param`    //渠道参数，最后生成到aksdk_config.xml中的map元素
	MetaData      []MetaData `json:metaData` //渠道参数，最后生成到AndroidManifest.xml中的meta-data元素
}

func (gameChannel *GameChannel) IsIcon() bool {
	return gameChannel.Icon != ""
}

func (gameChannel *GameChannel) GetSplashImgName(game *Game) string {

	if strings.ToLower(game.Orientation) == "" {
		return "landscape"
	} else {
		return strings.ToLower(game.Orientation)
	}
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
  "customConfig": {
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
	Operations  []Operation `json:operations`
	Plugins     []Plugin    `json:plugins`
	RJars       []RJar      `json:rJars`
}

type Plugin struct {
	Name string `json:name`
	Mold string `json:mold`
}

type RJar struct {
	Name string `json:name`
}
type Operation struct {
	Step string `json:step`
	Mold string `json:mold`
	From string `json:from`
	To   string `json:to`
}
