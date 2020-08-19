/*
  Created by guoxin in 2020/4/11 12:54 下午
  由于封装过后在日志打印时总会现在当前文件的目录，故注释
*/
package logging

import (
	"encoding/xml"
	"fmt"
	"github.com/GuoxinL/gcomponent/components"
	"github.com/GuoxinL/gcomponent/components/environment"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func init() {
	new(Configuration).Initialize()
}

type Configuration struct {
	Filename string
}

func (this Configuration) Initialize(params ...interface{}) interface{} {
	err := environment.GetConfig("components.logging", &this)
	if err != nil {
		panic("组件[Logging]读取配置异常" + err.Error())
	}
	filePath := environment.GetEnvironmentDirectory() + this.Filename
	logPath := getLogPath(filePath)
	startsWith := strings.HasPrefix(logPath, components.B)
	var fullLogPath string
	if !startsWith {
		wd, _ := os.Getwd()
		fullLogPath = wd + components.B + logPath
	} else {
		fullLogPath = logPath
	}

	err = os.MkdirAll(fullLogPath, 0711)
	if err != nil {
		Error0(err.Error())
	}
	Init()
	LoadConfiguration(filePath)
	Info("组件[Logging] init success")
	Warn("组件[Logging] init success")
	Error0("组件[Logging] init success")
	Debug("组件[Logging] init success")
	return nil
}
func getLogPath(configPath string) string {
	fmt.Println(GetCurrPath())
	fmt.Println(configPath)
	file, err := os.Open(configPath) // For read access.
	if err != nil {
		panic("组件[Logging]读取文件logging.xml异常" + err.Error())
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic("组件[Logging]读取文件logging.xml异常" + err.Error())
	}
	logging := Logging{}
	err = xml.Unmarshal(data, &logging)
	if err != nil {
		panic("组件[Logging]读取XML配置异常")
	}
	for _, filter := range logging.SFilter {
		for _, property := range filter.Property {
			if property.Name == "filename" {
				dir, _ := path.Split(property.InnerText)
				return dir
			}
		}
	}
	panic("组件[Logging]读取文件logging.xml异常，未配置logging.xml Filter节点")
}
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

type Logging struct {
	XMLName xml.Name  `xml:"logging"`
	SFilter []SFilter `xml:"filter"`
}
type SFilter struct {
	XMLName  xml.Name   `xml:"filter"`
	Property []Property `xml:"property"`
	Enabled  string     `xml:"enabled,attr"`
	Tag      string     `xml:"tag"`
	Type     string     `xml:"type"`
	Level    string     `xml:"level"`
}
type Property struct {
	XMLName   xml.Name `xml:"property"`
	Name      string   `xml:"name,attr"`
	InnerText string   `xml:",innerxml"`
}
