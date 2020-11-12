/*
   Created by guoxin in 2020/4/10 11:18 上午
*/
package environment

import (
	"flag"
	"github.com/GuoxinL/gcomponent/core"
	"github.com/gobike/envflag"
)

const DefaultProfile = "dev"

// TODO 这里如果初始化会因为调用了 envflag.Parse 出现和 testing.Init 争抢初始化，会报异常进而影响到项目编写测试
func init() {
	//new(Configuration).Initialize()
}

var instance Configuration

type Configuration struct {
	application *ApplicationConfig
	profile     string
}

func (c *Configuration) Initialize(params ...interface{}) interface{} {
	// Environment variables
	c.profile = *flag.String("profile", "", "Allowing us to map our beans to different profiles")
	directoryName := *flag.String("dir", core.ConfigDirectory, "The directory where the configuration files are located")
	config := *flag.String("config", core.ApplicationFile, "The name of the configuration fileThe name of the configuration file")
	envflag.Parse()

	c.application = newApplicationConfig(c.profile, directoryName)
	c.application.init(c.application.directory + config)

	instance = *c
	return nil
}

// Get the configuration profile
func GetProfile() string {
	return instance.profile
}

// Get the configuration in application.yaml under the current environment directory
func GetProperty(prefix string, config interface{}) error {
	if len(prefix) == 0 {
		err := instance.application.Unmarshal(&config)
		//设置配置文件类型
		if err != nil {
			return err
		}
		return nil
	}
	err := instance.application.UnmarshalKey(prefix, &config)
	//设置配置文件类型
	if err != nil {
		return err
	}
	return nil
}

// Get the configuration directory application.yaml viper.Viper
func GetApplication() *ApplicationConfig {
	return instance.application
}
