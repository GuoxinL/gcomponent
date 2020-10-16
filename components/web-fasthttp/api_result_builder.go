package web_fasthttp

import (
	"fmt"
	"github.com/GuoxinL/gcomponent/components/tools"
	"github.com/valyala/fasthttp"
)

type result0 struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	ctx     *fasthttp.RequestCtx
}

func (this result0) response() {
	json, _ := tools.ToJson(this)
	_, _ = fmt.Fprintf(this.ctx, json)
}

func Result() *result0 {
	return new(result0)
}
func (this result0) Ctx(ctx *fasthttp.RequestCtx) result0 {
	this.ctx = ctx
	return this
}

/*
build code and message
*/
func (this result0) BuildCodeMessage(code int, msg string) {
	this.Code = code
	this.Message = msg
	this.response()
}

/*
build ResultStatus
*/
func (this result0) BuildResultStatus(resultCode ResultStatus) {
	this.Code = resultCode.Code
	this.Message = resultCode.Message
	this.response()
}

/*
build code, message and data
*/
func (this result0) Build0(code int, message string, data interface{}) {
	this.Code = code
	this.Message = message
	this.Data = data
	this.response()
}

/*
build ok no params
*/
func (this result0) Ok() {
	this.Code = OK.Code
	this.Message = OK.Message
	this.response()
}

/*
build ok no params
*/
func (this result0) Ok0(data interface{}) {
	this.Code = OK.Code
	this.Message = OK.Message
	this.Data = data
	this.response()
}

/*
build bad request no params
*/
func (this result0) BadRequest() {
	this.Code = BAD_REQUEST.Code
	this.Message = BAD_REQUEST.Message
	this.response()
}

/*
build bad request msg
*/
func (this result0) BadRequest0(message string) {
	this.Code = BAD_REQUEST.Code
	this.Message = message
	this.response()
}

/*
build internal server error
*/
func (this result0) Error() {
	this.Code = INTERNAL_SERVER_ERROR.Code
	this.Message = INTERNAL_SERVER_ERROR.Message
	this.response()
}

/*
build internal server error msg
*/
func (this result0) Error0(message string) {
	this.Code = INTERNAL_SERVER_ERROR.Code
	this.Message = message
	this.response()
}
