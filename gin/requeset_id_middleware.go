// Package ggin
// Created by guoxin in 2020/10/31 10:01 下午
package ggin

import (
	"github.com/GuoxinL/gcomponent/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const headerXRequestID = "X-Request-ID"

func requestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader(headerXRequestID)
		if len(requestId) == 0 {
			requestId = uuid.New().String()
		}
		gzap.SetRequestIdFromLocal(requestId)
	}
}
