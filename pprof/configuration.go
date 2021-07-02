// Package gpprof
// Created by guoxin in 2020/6/5 4:35 下午
package gpprof

import (
	"github.com/GuoxinL/gcomponent/core"
	"github.com/GuoxinL/gcomponent/environment"
	gzap "github.com/GuoxinL/gcomponent/zap"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

var (
	initializeLock = core.NewInitLock()
)

func New(params ...interface{}) {
	c := &Configuration{
		InitializeLock: initializeLock,
	}
	c.Initialize(params...)
}

type Configuration struct {
	core.InitializeLock
	Port string `mapstructure:"port"`
	core.BEnable
}

func (c *Configuration) Initialize(params ...interface{}) interface{} {
	if c.IsInit() {
		//return &instances
		return nil
	}
	gzap.New()
	logger := gzap.GetInstance()
	logger.Info("GComponent [pprof] Initialize")
	err := environment.GetProperty("components.pprof", &c)
	if err != nil {
		logger.Fatal("GComponent [pprof] read config exception Exit!")
	}

	if c.Enable {
		_, err = strconv.Atoi(c.Port)
		if err != nil {
			logger.Error("GComponent [pprof] port number format exception.", zap.String("port", c.Port))
		}
		go func() {
			_ = http.ListenAndServe("0.0.0.0:"+c.Port, nil)
		}()
		logger.Info("GComponent [pprof] server init success", zap.String("port", c.Port))
	} else {
		logger.Warn("GComponent [pprof] pprof did not open, if you want to open please configure 'components.pprof.enable=true'")
	}
	return nil
}
