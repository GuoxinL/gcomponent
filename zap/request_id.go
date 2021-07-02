// Package core Create by guoxin 2021.07.02
package gzap

import (
	"github.com/go-eden/routine"
	"go.uber.org/zap"
)

var requestIdLocalStorage = routine.NewLocalStorage()

const RequestIdKey = "request_id"

func GetRequestIdFromLocal() string {
	requestId0 := requestIdLocalStorage.Get()
	return requestId0.(string)
}

func SetRequestIdFromLocal(requestId string) {
	requestIdLocalStorage.Set(requestId)
}

func getRequestIdField() zap.Field {
	var requestIdField zap.Field
	requestId := GetRequestIdFromLocal()
	if len(requestId) != 0 {
		requestIdField = zap.String(RequestIdKey, requestId)
	}
	return requestIdField
}
