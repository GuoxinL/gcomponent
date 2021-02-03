/**
  Create by guoxin 2020.12.16
*/
package ggin

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

const headerXRequestID = "X-Request-ID"
const requestIdKey = "request_id"

func getRequestId(requestId string) string {
    if len(requestId) == 0 {
        requestId = uuid.New().String()
    }
    return requestId
}

func requestId() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestId := getRequestId(c.GetHeader(headerXRequestID))
        c.Set(requestIdKey, requestId)
    }
}