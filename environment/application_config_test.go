/*
   Created by guoxin in 2020/11/10 8:34 上午
*/
package environment

import (
	"github.com/GuoxinL/gcomponent/core"
	"github.com/GuoxinL/gcomponent/tools"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	application := newApplicationConfig("", "")
	getwd, err := os.Getwd()
	if err != nil {
		t.Error("Getwd error")
	}
	application.init(getwd + core.B + core.ApplicationFile)
	ins := struct {
		Foo struct {
			Boo string   `yaml:"boo"`
			Coo int      `yaml:"coo"`
			Doo []string `yaml:"doo"`
		} `yaml:"foo"`
	}{}
	err = application.Unmarshal(&ins)
	if err != nil {
		t.Error("UnmarshalKey error")
	}
	json, err := tools.ToJson(ins)
	if err != nil {
		t.Error("ToJson error")
	}
	t.Log("yaml print: ", json)
}
