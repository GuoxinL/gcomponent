// Package environment
// Create by guoxin 2020.12.14
package environment

import (
    "github.com/GuoxinL/gcomponent/core"
    "testing"
)

// test Configuration.Initialize  It's only initialized once
func TestConfiguration_Initialize(t *testing.T) {
    if err := core.SetWorkDirectory(); err != nil {
        t.Error(err)
    }
    initialize0 := New()
    initialize1 := New()
    if initialize0 != initialize1 {
        t.Error("The two instances have different addresses")
    }
}

func TestGetName(t *testing.T) {
    if err := core.SetWorkDirectory(); err != nil {
        t.Error(err)
    }
    _ = New()
    name := GetName();
    if name != "application_name" {
        t.Error("Expected: xxx\nActual:", name)
    }
    t.Log(name)
}

// Test for repeated loading
func TestGetProfile(t *testing.T) {
    if err := core.SetWorkDirectory(); err != nil {
        t.Error(err)
    }
    _ = New()
    if profile := GetProfile(); profile != "dev" {
        t.Error("Expected: xxx\nActual:", profile)
    }
}
