// Package ggin
// Created by guoxin in 2020/10/31 10:01 下午
package ggin

import "net/http"

var (
    Ok                  = ResultMessage{http.StatusOK, http.StatusText(http.StatusOK)}
    BadRequest          = ResultMessage{http.StatusBadRequest, "bad request"}
    Unauthorized        = ResultMessage{http.StatusUnauthorized, "unauthorized"}
    Forbidden           = ResultMessage{http.StatusForbidden, "forbidden"}
    NotFound            = ResultMessage{http.StatusNotFound, "not found"}
    InternalServerError = ResultMessage{http.StatusInternalServerError, "internal server error"}
)

// ResultMessage 响应码对象
type ResultMessage struct {
    Code    int    `json:"code"`
    Message string `json:"msg"`
}

// HttpResultStatus
// http.StatusText(int)
func HttpResultStatus(code int) ResultMessage {
    return ResultMessage{code, http.StatusText(code)}
}
