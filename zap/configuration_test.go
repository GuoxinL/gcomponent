/**
  Create by guoxin 2020.12.15
*/
package zap

import (
    "github.com/GuoxinL/gcomponent/core"
    "go.uber.org/zap"
    "testing"
)

func TestConfiguration_Initialize(t *testing.T) {
    if err := core.SetWorkDirectory(); err != nil {
        t.Error(err)
    }
    new(Configuration).Initialize()

    zap.L().Info("我是一条日志，我是一条日志，我是一条日志")
}