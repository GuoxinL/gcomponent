// Package environment
// Create by guoxin 2020.12.14
package environment

import (
    "github.com/GuoxinL/gcomponent/core"
    "testing"
)

// test Configuration.Initialize  It's only initialized once
func TestConfiguration_Initialize(t *testing.T) {
    _ = core.SetWorkDirectory()
    initialize0 := New()
    initialize1 := New()
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

// Test for repeated loading
func TestGetProfile(t *testing.T) {
    _ = core.SetWorkDirectory()
    new(Configuration).Initialize()
    if profile := GetProfile(); profile != "dev" {
        t.Error("Expected: xxx\nActual:", profile)
    }
}
