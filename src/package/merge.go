package pack

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"strings"
)

func RenamePackage(gchannel *model.GameChannel, tempPath string) {
	xmlPath := tempPath + "/AndroidManifest.xml"
	content, err := ioutil.ReadFile(xmlPath)
	newPageName := ""
	if err != nil {
		log.Fatal(err)
	}

	name := GetOldPackageName(tempPath)
	fmt.Println("content :", name)
	if gchannel.PackageName == "" || gchannel.Suffix == "" {
		newPageName = name
		return
	}
	if gchannel.PackageName != "" || gchannel.Suffix == "" {
		newPageName = gchannel.PackageName

	}
	if gchannel.PackageName == "" || gchannel.Suffix != "" {
		newPageName = name + gchannel.Suffix
	}
	if gchannel.PackageName != "" || gchannel.Suffix != "" {
		newPageName = gchannel.PackageName + gchannel.Suffix
	}
	all := strings.ReplaceAll(string(content), name, newPageName)
	ioutil.WriteFile(xmlPath, []byte(all), 0777)
}

func GetOldPackageName(tempPath string) string {
	xmlPath := tempPath + "/AndroidManifest.xml"
	content, err := ioutil.ReadFile(xmlPath)
	if err != nil {
		log.Fatal(err)
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(content))
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				if "package" == attrName {
					return attr.Value
				}

			}
		// 处理元素结束（标签）
		case xml.EndElement:
			fmt.Printf("Token of '%s' end\n", token.Name.Local)
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
			// ...
		}

	}
	return ""
}
