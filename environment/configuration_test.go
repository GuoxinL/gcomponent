/**
  Create by guoxin 2020.12.14
*/
package environment

import (
    "github.com/GuoxinL/gcomponent/core"
    "testing"
)

// test Configuration.Initialize  It's only initialized once
func TestConfiguration_Initialize(t *testing.T) {
    _ = core.SetWorkDirectory()
    initialize0 := new(Configuration).Initialize()
    initialize1 := new(Configuration).Initialize()
    if initialize0 != initialize1 {
        t.Error("The two instances have different addresses")
    }
}

func TestGetName(t *testing.T) {
    _ = core.SetWorkDirectory()
    new(Configuration).Initialize()
    if name := GetName(); name != "xxx" {
        t.Error("Expected: xxx\nActual:", name)
    }
}

func TestGetProfile(t *testing.T) {
    _ = core.SetWorkDirectory()
    new(Configuration).Initialize()
    if profile := GetProfile(); profile != "dev" {
        t.Error("Expected: xxx\nActual:", profile)
    }
}
