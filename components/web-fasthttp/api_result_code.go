package web_fasthttp

/*
响应码对象
*/
type resultCode struct {
	Code    int32  `json:"code"`
	Message string `json:"msg"`
}

/*
因为和方法名称有冲突所以改用全大写加下划线，也表示静态变量了
*/
var (
	OK                    = resultCode{0, "OK"}
	BAD_REQUEST           = resultCode{400, "bad request"}
	UNAUTHORIZED          = resultCode{401, "unauthorized"}
	FORBIDDEN             = resultCode{403, "forbidden"}
	NOT_FOUND             = resultCode{404, "not found"}
	INTERNAL_SERVER_ERROR = resultCode{500, "internal server error"}
)
