/**
  Create by guoxin 2020.12.14
*/
package redis

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

func TestXxx(t *testing.T)  {
    var initLock bool

    fmt.Println(initLock)
}