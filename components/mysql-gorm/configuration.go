/*
   Created by guoxin in 2020/4/10 11:18 上午
*/
package mysql

import (
	"github.com/GuoxinL/gcomponent/components/environment"
	"github.com/GuoxinL/gcomponent/components/logging"
	"github.com/GuoxinL/gcomponent/components/tools"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func init() {
	new(Configuration).Initialize()
}

var dbInstances *dbInstanceMap

type Configuration struct {
	DataSources []mysqlConfigEntity
}

func (this *Configuration) Initialize(params ...interface{}) interface{} {
	err := environment.GetConfig("components.mysql", &this)
	if err != nil {
		logging.Exitf("组件[Mysql]读取配置异常, 退出程序！！！\n异常信息: %v", err.Error())
	}
	instanceMap := newDbInstanceMap()
	for _, dataSource := range this.DataSources {
		url := dataSource.Username + ":" + dataSource.Password + "@(" + dataSource.Url + ")/" + dataSource.Database + "?" + "charset=utf8&parseTime=True&loc=Local"
		db, err := gorm.Open("mysql", url)

		if err != nil {
			logging.Exitf("组件[Mysql]数据库连接失败, 退出程序！！！\n异常信息: %v, dsn:%v, errorDetail:%v", err.Error(), dataSource.Url, err.Error())
			return nil
		}
		db.SingularTable(true)
		db.SetLogger(Logger{})
		db.DB().SetMaxIdleConns(dataSource.MaxIdleConns)
		db.DB().SetMaxOpenConns(dataSource.MaxOpenConns)
		db.DB().SetConnMaxLifetime(time.Duration(dataSource.ConnMaxLifetime) * time.Second)
		instanceMap.Put(dataSource.Name, db)
		logging.Info("组件[Mysql]Name：%v", dataSource.Name)
		logging.Info("组件[Mysql]Username: %v, Password: %v", dataSource.Username, dataSource.Password)
		logging.Info("组件[Mysql]URL：%v", dataSource.Url+"/"+dataSource.Database)
		logging.Info("组件[Mysql]MaxIdleConns：%v, MaxOpenConns：%v, ConnMaxLifetime：%v",
			dataSource.MaxIdleConns,
			dataSource.MaxOpenConns,
			dataSource.ConnMaxLifetime)
		logging.Info("组件[Mysql]Connection '%v' init success", dataSource.Name)
	}
	dbInstances = instanceMap
	return nil
}

/**
废弃 请使用GetInstance(name)
*/
//Deprecated
func DB(name string) *gorm.DB {
	db := dbInstances.Get(name)
	if db == nil {
		logging.Error0("未找到`" + name + "`对应的数据库连接，请核对配置文件")
		return nil
	}
	return db.(*gorm.DB)
}

/**
通过该方法获得*gorm.DB对象
*/
func GetInstance(name string) *gorm.DB {
	return DB(name)
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

// gorm日志整合
var levelMap = map[string]logging.Level{
	"finest":   logging.FINEST,
	"fine":     logging.FINE,
	"debug":    logging.DEBUG,
	"trace":    logging.TRACE,
	"info":     logging.INFO,
	"warning":  logging.WARNING,
	"error":    logging.ERROR,
	"critical": logging.CRITICAL,
}

// Logger default logger
type Logger struct {
}

func getLevel(gormLevel string) logging.Level {
	return levelMap[gormLevel]
}

// Print format & print log
func (logger Logger) Print(values ...interface{}) {
	if !(len(values) >= 2) {
		logging.Info(values)
		return
	}
	gormLevel, ok := values[0].(string)
	if !ok {
		logging.Info(values)
	}
	massage, ok := values[1].(string)
	if !ok {
		logging.Info(values)
	}
	logging.Logf(getLevel(gormLevel), massage)
}
