/*
   Created by guoxin in 2020/4/10 11:18 上午
*/
package environment

import (
	"flag"
	"github.com/GuoxinL/gcomponent/components"
	"github.com/gobike/envflag"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	new(Configuration).Initialize()
}

var instance Configuration

type Configuration struct {
	applicationFile      viper.Viper
	environmentDirectory string
	profile              string
}

func (this *Configuration) Initialize(params ...interface{}) interface{} {
	// 环境变量
	this.initProfile()
	// 配置文件
	this.initApplicationFile()
	return nil
}

/**
获得当前环境目录
*/
func GetEnvironmentDirectory() string {
	return instance.environmentDirectory
}

/**
获得当前环境目录下application.yaml中的配置
*/
func GetConfig(prefix string, config interface{}) error {
	err := instance.applicationFile.UnmarshalKey(prefix, &config)
	//设置配置文件类型
	if err != nil {
		return err
	}
	return nil
}

/**
获得当前环境目录下application.yaml的Viper对象
*/
func GetApplicationFile() viper.Viper {
	return instance.applicationFile
}

func (this *Configuration) initProfile() {
	flag.StringVar(&this.profile, "profile", "dev", "this is #")
	envflag.Parse()
	//The following profiles are active: dev
}

func (this *Configuration) initApplicationFile() {
	this.environmentDirectory = components.ConfigDirectory + components.B + this.profile + components.B
	location := this.environmentDirectory + components.ApplicationFile

	configType := location[strings.LastIndex(location, ".")+1:]
	configName := location[strings.LastIndex(location, "/")+1 : strings.LastIndex(location, ".")]
	configPath := location[:strings.LastIndex(location, "/")+1]
	newViper := viper.New()
	// 配置文件的名字
	newViper.SetConfigName(configName)
	// 配置文件的类型
	newViper.SetConfigType(configType)
	// 配置文件的位置
	//workDir, _ := os.Getwd()
	newViper.AddConfigPath(configPath)
	// 观察
	newViper.WatchConfig()
	_ = newViper.ReadInConfig()
	this.applicationFile = *newViper
	instance = *this
}
