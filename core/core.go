// Package core
// The core package contains some core interfaces and common methods
//    Responsibility
//    1. Initialize
//    2. core interfaces
//    3.
//	Created by guoxin in 2020/4/10 11:25 上午
package core

import (
	"go.uber.org/atomic"
	"os"
	"path"
	"runtime"
)

// InitializeLock Make sure you only initialize it once
type InitializeLock struct {
	*atomic.Bool
}

func (l *InitializeLock) IsInit() bool {
	if l.Load() {
		return true
	}
	l.Store(true)
	return false
}

func NewInitLock() InitializeLock {
	return InitializeLock{atomic.NewBool(false)}
}

type Initialize interface {
	// Initialize
	// 初始化配置
	// 没有泛型使用interface{}替代
	Initialize(params ...interface{}) interface{}
}

// Properties Properties.load() -> Properties.prefix() -> Initialize.Initialize()
type Properties interface {
	/*
		前缀
	*/
	prefix() string
	/*
		加载文件配置
	*/
	load()
}

type BEnable struct {
	Enable bool `mapstructure:"enable"`
}

const (
	// S separator
	S string = "."
	// B Backslash
	B string = "/"
	// C colon
	C string = ":"
	// AT at
	AT string = "@"
	// Q ?
	Q string = "?"
)

// SetWorkDirectory set work directory
func SetWorkDirectory() error {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	return os.Chdir(dir)
}
