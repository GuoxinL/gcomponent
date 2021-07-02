// Package ggin Create by guoxin 2021.07.02
package ggin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, ResultMessage{http.StatusNotFound, "not found"})
	return
}
