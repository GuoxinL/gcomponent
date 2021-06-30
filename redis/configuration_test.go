// Package gredis
// Created by guoxin in 2020/4/13 1:34 下午
package gredis

import (
    "fmt"
    "github.com/GuoxinL/gcomponent/core"
    "testing"
)

func TestGetInstance(t *testing.T) {
    if err := core.SetWorkDirectory(); err != nil {
        t.Error(err)
    }
    new(Configuration).Initialize()
    instance := GetInstance("root")
    instance.Get("xxx")
}

func TestXxx(t *testing.T) {
    var initLock bool

    fmt.Println(initLock)
}
