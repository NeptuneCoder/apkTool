package env

import (
	"encoding/json"
	"fmt"
	"github.com/yanghai23/GoLib/atfile"
	"model"
	"utils"
)

var Env *model.Environment

func init() {
	readEnvConfig()
}

func readEnvConfig() error {
	s := utils.GetDirRel()
	fmt.Println(s)
	data, err := atfile.ReadConfig(s, "Environment.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	Env = new(model.Environment)
	json.Unmarshal(data, &Env)
	return nil
}
