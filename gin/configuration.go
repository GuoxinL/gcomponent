// Package ggin
// Created by guoxin in 2020/10/25 5:58 下午
package ggin

import (
	"github.com/GuoxinL/gcomponent/core"
	"github.com/GuoxinL/gcomponent/environment"
	"github.com/GuoxinL/gcomponent/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

var (
	initializeLock = core.NewInitLock()
	instance       *gin.Engine
)

func New(params ...interface{}) {
	c := &Configuration{
		InitializeLock: initializeLock,
	}
	c.Initialize(params)
}

type Configuration struct {
	// Make sure you only initialize it once
	core.InitializeLock
	Port   string
	router *gin.Engine
}

// Initialize
// https://www.cnblogs.com/kainhuck/p/13333765.html
func (c *Configuration) Initialize(params ...interface{}) interface{} {
	if c.IsInit() {
		return &instance
	}
	gzap.New()
	zap.L().Info("GComponent [gin] Initialize")
	err := environment.GetProperty("components.web.gin", &c)
	if err != nil {
		zap.L().Fatal("GComponent [gin] read config exception exit.", zap.Error(err))
	}
	c.router = gin.New()
	// method not found
	c.router.NoMethod(HandlerNotFound)
	// route not found
	c.router.NoRoute(HandlerNotFound)
	// request_id
	isRequestId := environment.IsRequestId()
	if isRequestId {
		c.router.Use(requestId())
	}
	// Recovery error log, enable stack.
	c.router.Use(RecoveryWithZap(gzap.GetInstance(), true))
	// gin zap logger
	c.router.Use(Ginzap(gzap.GetInstance(), time.RFC3339, true))

	if webConfigurationInterface := params[0]; webConfigurationInterface != nil {
		webConfiguration, ok := webConfigurationInterface.(ControllerConfiguration)
		if !ok {
			zap.L().Warn("GComponent [gin] please implement ggin.ControllerConfiguration")
		}
		webConfiguration.Configure(Controllers{Engine: c.router})
	}
	instance = c.router
	zap.L().Info("GComponent [gin] server init success", zap.String("port", c.Port))
	zap.L().Error("router.init error", zap.Error(c.router.Run(c.Port)))
	return nil
}
