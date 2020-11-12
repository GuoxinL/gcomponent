/*
   Created by guoxin in 2020/4/11 5:13 下午
*/
package web_fasthttp

import (
	"fmt"
	"github.com/GuoxinL/gcomponent/environment"
	"github.com/GuoxinL/gcomponent/logging"
	"github.com/GuoxinL/gcomponent/tools"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

type WebRouteConfiguration interface {
	Configure(router *router.Router)
}

type Configuration struct {
	Port   string
	router *router.Router
}

func (this *Configuration) Initialize(params ...interface{}) interface{} {
	logging.Info("GComponent [web-fasthttp]初始化接口")
	err := environment.GetProperty("components.web", &this)
	if err != nil {
		logging.Exitf("GComponent [web-fasthttp]读取配置异常, 退出程序！！！\n异常信息: %v", err.Error())
	}
	this.router = router.New()
	if webConfigurationInterface := params[0]; webConfigurationInterface != nil {
		webConfiguration, ok := webConfigurationInterface.(WebRouteConfiguration)
		if !ok {
			_ = logging.Warn("GComponent [web-fasthttp]请实现 web_fasthttp.WebRouteConfiguration 接口")
		}

		webConfiguration.Configure(this.router)
	}
	list := this.router.List()
	for method, paths := range list {
		logging.Info("GComponent [web-fasthttp]Method %v\tPath %v", method, strings.Replace(strings.Trim(fmt.Sprint(paths), "[]"), " ", ",", -1))
	}
	logging.Info("GComponent [web-fasthttp]Server init success port: %v", this.Port)
	err = run(":"+this.Port, RequestPanicFilter(RequestInfoFilter(this.router.Handler)))
	if err != nil {
		logging.Exitf("GComponent [web-fasthttp]启动失败: %v 退出程序！！！", err.Error())
	}
	return nil
}

/*
	FastHttp start
*/
func run(addr string, handler fasthttp.RequestHandler) error {
	s := &fasthttp.Server{
		Handler: handler,
		// logging 配合 FastHttp
		Logger: FastHttpLogger{},
	}
	return s.ListenAndServe(addr)
}
func RequestInfoFilter(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		startMillisecond := time.Now().UnixNano() / 1e6
		logging.Error0("-----------------------------------------接口请求开始-----------------------------------------")
		logging.Error0("Method: %v URI: %v", string(ctx.Method()), ctx.Request.URI().String())
		logging.Error0("query: %v", ctx.QueryArgs().String())
		logging.Error0("request body: %v", string(ctx.PostBody()))
		next(ctx)
		endMillisecond := time.Now().UnixNano() / 1e6
		logging.Error0("请求时间（毫秒）: %d", endMillisecond-startMillisecond)
		logging.Error0("response body: %v", string(ctx.Response.Body()))
		logging.Error0("-----------------------------------------接口请求结束-----------------------------------------")
	}
}
func RequestPanicFilter(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tools.TryCatch{}.Try(func() {
			next(ctx)
		}).CatchAll(func(err error) {
			logging.Error0("GComponent [web-fasthttp]请求处理异常 %v", err.Error())
			Result().Ctx(ctx).Error0("请求处理异常，异常信息：" + err.Error())
		}).Finally(func() {})
	}
}
func NotFound(ctx *fasthttp.RequestCtx) {
	Result().Ctx(ctx).BadRequest0("未匹配到接口 URI：" + ctx.URI().String())
}

type FastHttpLogger struct {
}

func (f FastHttpLogger) Printf(format string, args ...interface{}) {
	logging.Info(format, args)
}
