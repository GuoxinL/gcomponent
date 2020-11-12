/*
   Created by guoxin in 2020/4/11 7:14 下午
*/
package web_fasthttp

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

type parameter0 struct {
	ctx *fasthttp.RequestCtx
}

func (this *parameter0) Read(i interface{}) (interface{}, error) {
	body := this.ctx.PostBody()
	this.ctx.Method()
	err := json.Unmarshal(body, i)
	if err != nil {
		return nil, errors.New("Json转换异常：" + err.Error())
	}
	return i, nil
}

func (this *parameter0) Ctx(ctx *fasthttp.RequestCtx) *parameter0 {
	this.ctx = ctx
	ctx.PostBody()
	return this
}

func Parameter() *parameter0 {
	return new(parameter0)
}
