/*
   Created by guoxin in 2020/11/8 6:38 下午
*/
package environment

import (
	"fmt"
	"github.com/GuoxinL/gcomponent/core"
	"github.com/spf13/viper"
	"strings"
)

type ApplicationConfig struct {
	viper.Viper
	directory     string
	directoryName string
}

// Get the current configuration directory
func (a *ApplicationConfig) Directory() string {
	return a.directory
}

func newApplicationConfig(profile string, directoryName string) *ApplicationConfig {
	a := &ApplicationConfig{Viper: *viper.New()}
	a.directoryName = directoryName
	if len(profile) == 0 {
		a.directory = a.directoryName + core.B
	} else {
		a.directory = a.directoryName + core.B + profile + core.B
	}
	return a
}

func (a *ApplicationConfig) init(fullPath string) {
	configType := fullPath[strings.LastIndex(fullPath, ".")+1:]
	configName := fullPath[strings.LastIndex(fullPath, "/")+1 : strings.LastIndex(fullPath, ".")]
	configPath := fullPath[:strings.LastIndex(fullPath, "/")+1]

	// 配置文件的名字
	a.SetConfigName(configName)
	// 配置文件的类型
	a.SetConfigType(configType)
	// 配置文件的位置
	//workDir, _ := os.Getwd()
	a.AddConfigPath(configPath)
	// 观察
	a.WatchConfig()
	err := a.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Read in config error: %v ", err.Error()))
	}
}
