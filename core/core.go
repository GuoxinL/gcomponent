/*
   	The core package contains some core interfaces and common methods
    Responsibility
    1. 初始化
    2. core interfaces
    3.
	Created by guoxin in 2020/4/10 11:25 上午
*/
package core

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

type Enable struct {
	Enable bool `mapstructure:"enable"`
}

const (
	/*
		分隔符(separator)
	*/
	S string = "."
	/*
		斜杠(Backslash)
	*/
	B string = "/"
	/*
		冒号(colon)
	*/
	C string = ":"
	/*
		艾特
	*/
	AT string = "@"
)
