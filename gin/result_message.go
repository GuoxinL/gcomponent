package ggin

import "net/http"

var (
    Ok                  = ResultMessage{http.StatusOK, http.StatusText(http.StatusOK)}
    BadRequest          = ResultMessage{400, "bad request"}
    Unauthorized        = ResultMessage{401, "unauthorized"}
    Forbidden           = ResultMessage{403, "forbidden"}
    NotFound            = ResultMessage{404, "not found"}
    InternalServerError = ResultMessage{500, "internal server error"}
)

/*
响应码对象
*/
type ResultMessage struct {
    Code    int    `json:"code"`
    Message string `json:"msg"`
}

/**
http.StatusText(int)
*/
func HttpResultStatus(code int) ResultMessage {
    return ResultMessage{code, http.StatusText(code)}
}
