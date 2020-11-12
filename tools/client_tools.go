/*
   Created by guoxin in 2020/5/12 11:47 上午
*/
package tools

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"time"
)

/**
Post 请求
url: 请求路径
request: 请求对象会被放到request body中
response: 响应对象会从响应body中通过json反序列化到response中
args: query参数将会放到路径"？"后面，例如：?operation=1
*/
func Post(url string, request, response interface{}, timeout time.Duration) error {
	r, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req := &fasthttp.Request{}
	req.SetRequestURI(url)
	req.SetBody(r)
	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	res := &fasthttp.Response{}
	client := &fasthttp.Client{}
	if err := client.DoTimeout(req, res, timeout); err != nil {
		return errors.New("请求失败:" + err.Error())
	}
	body := res.Body()
	err = json.Unmarshal(body, response)
	// TODO json.Unmarshal 内存占用持续累计不释放
	req.ConnectionClose()
	res.ConnectionClose()
	if err != nil {
		return errors.New("Json转换异常" + err.Error() + ", Json: " + string(res.Body()))
	}
	return nil
}
