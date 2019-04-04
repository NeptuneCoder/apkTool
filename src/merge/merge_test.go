package merge

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestRenamePackage(t *testing.T) {
	xmlPath := "/AndroidManifest.xml"
	content, err := ioutil.ReadFile(xmlPath)
	if err != nil {
		log.Fatal(err)
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(content))
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", token.Name)
			fmt.Printf("Token name: %s\n", name)
			fmt.Printf("Token len: %d\n", len(token.Attr))
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				if "merge" == attrName {
				}

				fmt.Printf("An attribute is: %s %s\n", attrName, attr.Value)
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
}
