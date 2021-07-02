// Package ggin
// Created by guoxin in 2020/10/31 10:01 下午
package ggin

// ResultMessage 响应码对象
type ResultMessage struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
