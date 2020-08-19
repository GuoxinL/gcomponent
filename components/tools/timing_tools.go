/*
   Created by guoxin in 2020/5/19 10:49 上午
*/
package tools

import (
	"go.uber.org/atomic"
	"time"
)

/**
定时执行工具
*/
type timing struct {
	d time.Duration
	b *atomic.Bool
}

func NewTiming(d time.Duration) timing {
	t := timing{}
	t.b = atomic.NewBool(false)
	t.d = d
	return t
}

func (this timing) Execute(fn func()) timing {
	go func(b *atomic.Bool) {
		for {
			if b.Load() {
				return
			}
			fn()
			time.Sleep(this.d)
		}
	}(this.b)
	return this
}

func (this timing) Stop() {
	this.b.Toggle()
}
