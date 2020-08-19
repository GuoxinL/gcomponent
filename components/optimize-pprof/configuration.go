/*
   Created by guoxin in 2020/6/5 4:35 下午
*/
package optimize

import (
	"github.com/GuoxinL/gcomponent/components/environment"
	"github.com/GuoxinL/gcomponent/components/logging"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func init() {
	new(Configuration).Initialize()
}

type Configuration struct {
	Port   string `yaml:"port"`
	Enable bool   `yaml:"enable"`
}

func (this *Configuration) Initialize(params ...interface{}) interface{} {
	logging.Info("组件[optimize-pprof]初始化接口")
	err := environment.GetConfig("components.optimize.pprof", &this)
	if err != nil {
		logging.Error0("组件[optimize-pprof]读取配置异常, 退出程序！！！")
	}

	if this.Enable {
		_, err = strconv.Atoi(this.Port)
		if err != nil {
			logging.Error0("组件[optimize-pprof]端口号格式异常。 Port: %v", this.Port)
		}
		go func() {
			_ = http.ListenAndServe("0.0.0.0:"+this.Port, nil)
		}()
		logging.Info("组件[optimize-pprof]启动成功0.0.0.0:%v", this.Port)
	} else {
		_ = logging.Warn("组件[optimize-pprof]pprof未开启，如需开启请配置'components.optimize.pprof.enable=true'")
	}
	return nil
}
