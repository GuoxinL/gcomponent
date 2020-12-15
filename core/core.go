/*
   	The core package contains some core interfaces and common methods
    Responsibility
    1. 初始化
    2. core interfaces
    3.
	Created by guoxin in 2020/4/10 11:25 上午
*/
package core

import (
	"os"
	"path"
	"runtime"
)

type Initialize interface {
	/*
		初始化配置
		没有泛型使用interface{}替代
	*/
	Initialize(params ...interface{}) interface{}
}

// Properties.load() -> Properties.prefix() -> Initialize.Initialize()
type Properties interface {
	/*
		前缀
	*/
	prefix() string
	/*
		加载文件配置
	*/
	load()
}

type BEnable struct {
	Enable bool `mapstructure:"enable"`
}

const (
	// separator
	S string = "."
	// Backslash
	B string = "/"
	// colon
	C string = ":"
	// at
	AT string = "@"
)

// set work directory
func SetWorkDirectory() error {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	return os.Chdir(dir)
}
