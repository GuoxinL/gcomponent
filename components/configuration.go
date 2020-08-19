/*
   Created by guoxin in 2020/4/10 11:25 上午
*/
package components

type Initialize interface {
	/*
		初始化配置
		没有泛型使用interface{}替代
	*/
	Initialize(params ...interface{}) interface{}
}

/*
	load() -> prefix() -> Initialize()
*/
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

	YAML string = "yaml"
	/*
	   配置文件名称
	*/
	ConfigFileName = "application"
	/**
	  配置文件目录
	*/
	ConfigDirectory = "resource"
	/*
	   配置文件名称
	*/
	ConfigFilePath = S + B + ConfigDirectory + B
	/*
		主配置文件：
		存放各种非业务配置，或者轻业务配置
	*/
	ApplicationFile = ConfigFileName + S + YAML
)

const ()
