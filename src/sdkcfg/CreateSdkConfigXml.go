package sdkcfg

import (
	"dom4g"
	"io/ioutil"
	"model"
	"strings"
	"utils"
)

func AddParamElement(rootEle *dom4g.Element, name, value string) {
	el2 := dom4g.NewElement("param", "")
	el2.AddAttr("name", name)
	el2.AddAttr("value", value)
	_ = rootEle.AddNode(el2)

}
func CreateSdkConfigXml(channel *model.GameChannel, sdkCfg *model.SdkRootConfig, game *model.Game, mainActivityName, tempPath string) {

	el := dom4g.NewElement("config", "")
	for _, v := range channel.Param {
		_ = el.AddNode(dom4g.NewElement(v.Name, v.Value))
	}
	if game.FoyoentId != "" {
		AddParamElement(el, "FoyoentCID", game.FoyoentId)
	}
	if game.Url != "" {
		AddParamElement(el, "BASE_HOST_URL", game.Url)
	}

	AddParamElement(el, "GAMEACTIVITYNAME", mainActivityName)

	if game.Debug {
		AddParamElement(el, "DEBUG", "true")
	}
	if game.Orientation != "" && strings.Contains("PORTRAIT", strings.ToLower(game.Orientation)) {

		AddParamElement(el, "ORIENTATION", "PORTRAIT")
	}

	if "" != sdkCfg.Config.Application {
		appsElement := dom4g.NewElement("applications", "")
		app := dom4g.NewElement("application", "")
		app.AddAttr("name", sdkCfg.Config.Application)
		_ = appsElement.AddNode(app)
		_ = el.AddNode(appsElement)
	}

	for _, v := range sdkCfg.Config.Plugins {
		appsElement := dom4g.NewElement("plugins", "")
		app := dom4g.NewElement("plugin", "")
		app.AddAttr("name", v.Name)
		app.AddAttr("type", v.Mold)
		_ = appsElement.AddNode(app)
		_ = el.AddNode(appsElement)
	}

	// 另一种输出方式，记得要调用flush()方法,否则输出的文件中显示空白
	assetsPath := tempPath + "/assets"
	utils.CreateNewFolder(assetsPath)
	configNomalXmlPath := assetsPath + "/foyoentcfg_nomal.xml"
	xml := el.SyncToXml()
	_ = ioutil.WriteFile(configNomalXmlPath, []byte(xml), 777)
}
