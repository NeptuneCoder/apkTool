package replace

import (
	"fmt"
	"io/ioutil"
	"model"
	"strings"
)

func ReplacePkgManifest(tempPath, newPkgName string, channel *model.GameChannel) {
	params := channel.Param
	if params == nil {
		params = []model.Param{}
	}
	params[0] = model.Param{Name: "PACKAGENAME", Value: newPkgName}
	androidManifestName := tempPath + "/AndroidManifest.xml"
	content, _ := ioutil.ReadFile(androidManifestName)
	str := string(content)
	for _, v := range channel.Param {

		fmt.Println("name:", v.Name, "  value:", v.Value)
		str = strings.ReplaceAll(str, "{"+v.Name+"}", v.Value)
	}
	_ = ioutil.WriteFile(androidManifestName, []byte(str), 0777)
}
