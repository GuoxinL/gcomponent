// Package ggorm
// Created by guoxin in 2020/4/10 11:18 上午
package ggorm

import (
	"errors"
	"github.com/GuoxinL/gcomponent/core"
	"github.com/GuoxinL/gcomponent/core/tools"
	"github.com/GuoxinL/gcomponent/environment"
	gzap "github.com/GuoxinL/gcomponent/zap"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"time"
)

var (
	initializeLock   = core.NewInitLock()
	instances        *dbInstanceMap
	DatabaseNotFound = errors.New(`database instance not found`)
)

func New(params ...interface{}) {
	c := &Configuration{
		InitializeLock: initializeLock,
	}
	c.Initialize(params...)
}

type Configuration struct {
	// Make sure you only initialize it once
	core.InitializeLock

	DataSources []mysqlConfigEntity `json:"dataSources" mapstructure:"dataSources"`
}

func (c *Configuration) Initialize(params ...interface{}) interface{} {
	if c.IsInit() {
		return &instances
	}
	gzap.New()
	logger := gzap.GetInstance()
	err := environment.GetProperty("components.gorm", &c)
	if err != nil {
		logger.Fatal("GComponent [gorm] read config exception Exit!", zap.Error(err))
	}
	instanceMap := newDbInstanceMap()
	for _, dataSource := range c.DataSources {
		logger0 := zapgorm2.New(logger)
		logger0.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
		db, err := gorm.Open(mysql.Open(dataSource.url()), &gorm.Config{
			Logger:                                   logger0,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			logger.Fatal("GComponent [gorm] database connection failed.", zap.Error(err), zap.String("url", dataSource.Url))
			return nil
		}
		db = db.Debug()

		sqldb, err := db.DB()
		if err != nil {
			logger.Fatal("GComponent [gorm] connection exception.", zap.Error(err))
		}

		sqldb.SetMaxIdleConns(dataSource.MaxIdleConns)
		sqldb.SetMaxOpenConns(dataSource.MaxOpenConns)
		sqldb.SetConnMaxLifetime(time.Duration(dataSource.ConnMaxLifetime) * time.Second)
		instanceMap.Put(dataSource.Name, db)
		logger.Info("GComponent [gorm] properties",
			zap.String("name", dataSource.Name),
			zap.String("username", dataSource.Username),
			zap.String("password", dataSource.Password),
			zap.String("url", dataSource.Url+"/"+dataSource.Database),
			zap.Int("MaxIdleConns", dataSource.MaxIdleConns),
			zap.Int("MaxOpenConns", dataSource.MaxOpenConns),
			zap.Int("ConnMaxLifetime", dataSource.ConnMaxLifetime),
		)
		logger.Info("GComponent [gorm] connection '" + dataSource.Name + "' init success")
	}
	instances = instanceMap
	return nil
}

// GetInstance0
// 通过该方法获得*gorm.DB对象
func GetInstance0(name string) (*gorm.DB, error) {
	db := instances.Get(name)
	if db == nil {
		zap.L().Error("'" + name + "' collection not found")
		return nil, DatabaseNotFound
	}
	return db.(*gorm.DB), nil
}

// GetInstance
// 通过该方法获得*gorm.DB对象
func GetInstance(name string) *gorm.DB {
	instance0, err := GetInstance0(name)
	if err == DatabaseNotFound {
		zap.L().Error("'" + name + "' collection not found")
		return nil
	}

	return instance0
}

type dbInstanceMap struct {
	*tools.ConcurrentMap
}

func newDbInstanceMap() *dbInstanceMap {
	return &dbInstanceMap{tools.NewConcurrentMap()}
}

type mysqlConfigEntity struct {
	Name            string
	Url             string
	Database        string
	Username        string
	Password        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
}

func (c mysqlConfigEntity) url() string {
	return c.Username + core.C + c.Password + core.AT + "(" + c.Url + ")" + core.B + c.Database + core.Q + "charset=utf8&parseTime=True&loc=Local"
}
