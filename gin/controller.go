/*
   Created by guoxin in 2020/11/1 4:08 下午
*/
package ggin

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

func (c *Controllers) Register(method string, path string, controller Controller) {
    c.Handle(method, path, func(context *gin.Context) {
        p := &Parameter{context}
        result := controller(p)
        p.SecureJSON(result.code, result.data)
    })
}
func (c *Controllers) GET(path string, controller Controller) {
    c.Register(http.MethodGet, path, controller)
}

func (c *Controllers) HEAD(path string, controller Controller) {
    c.Register(http.MethodGet, path, controller)
}

func (c *Controllers) POST(path string, controller Controller) {
    c.Register(http.MethodPost, path, controller)
}

func (c *Controllers) PUT(path string, controller Controller) {
    c.Register(http.MethodPut, path, controller)
}

func (c *Controllers) PATCH(path string, controller Controller) {
    c.Register(http.MethodPatch, path, controller)
}

func (c *Controllers) DELETE(path string, controller Controller) {
    c.Register(http.MethodDelete, path, controller)
}

func (c *Controllers) CONNECT(path string, controller Controller) {
    c.Register(http.MethodConnect, path, controller)
}

func (c *Controllers) OPTIONS(path string, controller Controller) {
    c.Register(http.MethodOptions, path, controller)
}

func (c *Controllers) TRACE(path string, controller Controller) {
    c.Register(http.MethodTrace, path, controller)
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
