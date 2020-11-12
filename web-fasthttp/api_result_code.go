package web_fasthttp

import "net/http"

/*
响应码对象
*/
type ResultStatus struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

/**
http.StatusText(int)
*/
func HttpResultStatus(code int) ResultStatus {
	return ResultStatus{code, http.StatusText(code)}
}

/*
因为和方法名称有冲突所以改用全大写加下划线，也表示静态变量了
*/
var (
	OK                    = ResultStatus{http.StatusOK, http.StatusText(http.StatusOK)}
	BAD_REQUEST           = ResultStatus{400, "bad request"}
	UNAUTHORIZED          = ResultStatus{401, "unauthorized"}
	FORBIDDEN             = ResultStatus{403, "forbidden"}
	NOT_FOUND             = ResultStatus{404, "not found"}
	INTERNAL_SERVER_ERROR = ResultStatus{500, "internal server error"}
)
