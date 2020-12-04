/*
   Created by guoxin in 2020/4/10 11:18 上午
*/
package environment

import (
	"flag"
	"fmt"
	"github.com/GuoxinL/gcomponent/core"
	"github.com/gobike/envflag"
)

const (
	YAML            = "yaml"
	APPLICATION     = "application"
	ConfigDirectory = "conf"

	DefaultProfile        = "dev"
	DefaultConfigFileName = APPLICATION + core.S + YAML
)

// TODO 这里如果初始化会因为调用了 envflag.Parse 出现和 testing.Init 争抢初始化，会报异常进而影响到项目编写测试
//func init() {
//	new(Configuration).Initialize()
//}

var instance Configuration

// The properties that Application has, If the future properties out more application-specific properties put in here
type Application struct {
	Name    string `mapstructure:"name"`
	Profile string `mapstructure:"profile"`
}

type Configuration struct {
	// configuration file
	configurationFile *ApplicationFile
	// Application properties
	application Application
}

func (c *Configuration) Initialize(params ...interface{}) interface{} {
	// Environment variables
	profile := *flag.String("profile", "", "Allowing us to map our beans to different profiles")
	directoryName := *flag.String("dir", ConfigDirectory, "The directory where the configuration files are located")
	envflag.Parse()

	// Application init
	c.configurationFile = newApplicationFile(profile, directoryName)
	application := Application{}
	if err := c.configurationFile.UnmarshalKey("components.application", &application); err != nil {
		fmt.Println("Parse environment.Application exception")
		application.Name = "wuming"
	}
	c.application = application

	instance = *c
	return nil
}

func IsProfile(profile string) bool {
	return instance.application.Profile == profile
}

// Get the configuration profile
func GetProfile() string {
	return instance.application.Profile
}

// Get the configuration profile
func GetName() string {
	return instance.application.Name
}

// Get the configuration in application.yaml under the current environment directory
func GetProperty(prefix string, config interface{}) error {
	if len(prefix) == 0 {
		err := instance.configurationFile.Unmarshal(&config)
		if err != nil {
			return err
		}
		return nil
	}
	err := instance.configurationFile.UnmarshalKey(prefix, &config)
	if err != nil {
		return err
	}
	return nil
}

// Get the configuration directory application.yaml viper.Viper
func GetApplicationFile() *ApplicationFile {
	return instance.configurationFile
}
