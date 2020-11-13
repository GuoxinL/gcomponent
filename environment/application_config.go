/*
   Created by guoxin in 2020/11/8 6:38 下午
*/
package environment

import (
	"fmt"
	"github.com/GuoxinL/gcomponent/core"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type ApplicationFile struct {
	viper.Viper
	absolutePath  string
	directory     string
	directoryName string
}

// Get the current configuration directory
func (a *ApplicationFile) Directory() string {
	return a.absolutePath
}

func newApplicationFile(profile string, directoryName string) *ApplicationFile {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	a := &ApplicationFile{Viper: *viper.New()}
	a.directoryName = directoryName
	if len(profile) == 0 {
		a.absolutePath = pwd + core.B + a.directoryName + core.B
	} else {
		a.absolutePath = pwd + core.B + a.directoryName + core.B + profile + core.B
	}
	a.init(a.absolutePath, DefaultConfigFileName)
	return a
}

func (a *ApplicationFile) init(absolutePath, configurationFileName string) *ApplicationFile {
	name, type0 := getFile(configurationFileName)
	// 配置文件的名字
	a.SetConfigName(name)
	// 配置文件的类型
	a.SetConfigType(type0)
	// 配置文件的位置
	a.AddConfigPath(absolutePath)

	a.WatchConfig()
	a.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	a.AutomaticEnv()
	if err := a.ReadInConfig(); err != nil {
		fmt.Println(fmt.Errorf("Read in config error: %v ", err.Error()))
	}
	return a
}

func getFile(filename string) (name, type0 string) {
	type0 = filename[strings.LastIndex(filename, core.S)+1:]
	name = filename[strings.LastIndex(filename, core.B)+1 : strings.LastIndex(filename, ".")]
	//path = fullPath[:strings.LastIndex(fullPath, core.B)+1]
	return
}
