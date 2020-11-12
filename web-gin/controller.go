/*
   Created by guoxin in 2020/11/1 4:08 下午
*/
package web_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register the Controller with the gin.RouterGroup route
type ControllerConfiguration interface {
	Configure(Controllers)
}

// Implement Controller coding the interface
type Controller func(p *Parameter) *Result

func test(p *Parameter) *Result {
	return nil
}

// Here can use gin.Engine and gin.RouterGroup.
// gin.RouterGroup.Group() No good ideas yet
// Controllers are c *gin.Engine wrappers
type Controllers struct {
	*gin.Engine
}

func (this *Controllers) Register(method string, path string, controller Controller) {
	this.Handle(method, path, func(context *gin.Context) {
		p := &Parameter{context}
		result := controller(p)
		p.SecureJSON(result.code, result.data)
	})
}
func (this *Controllers) GET(path string, controller Controller) {
	this.Register(http.MethodGet, path, controller)
}

func (this *Controllers) HEAD(path string, controller Controller) {
	this.Register(http.MethodGet, path, controller)
}

func (this *Controllers) POST(path string, controller Controller) {
	this.Register(http.MethodPost, path, controller)
}

func (this *Controllers) PUT(path string, controller Controller) {
	this.Register(http.MethodPut, path, controller)
}

func (this *Controllers) PATCH(path string, controller Controller) {
	this.Register(http.MethodPatch, path, controller)
}

func (this *Controllers) DELETE(path string, controller Controller) {
	this.Register(http.MethodDelete, path, controller)
}

func (this *Controllers) CONNECT(path string, controller Controller) {
	this.Register(http.MethodConnect, path, controller)
}

func (this *Controllers) OPTIONS(path string, controller Controller) {
	this.Register(http.MethodOptions, path, controller)
}

func (this *Controllers) TRACE(path string, controller Controller) {
	this.Register(http.MethodTrace, path, controller)
}

/**
Http result
*/
type Result struct {
	// Http status code
	code int
	// Http response body
	data map[string]interface{}
}
