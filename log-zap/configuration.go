/*
   Created by guoxin in 2020/4/13 1:34 下午
*/
package log_zap

import (
	"encoding/json"
	"fmt"
	"github.com/GuoxinL/gcomponent/core"
	"github.com/GuoxinL/gcomponent/environment"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	new(Configuration).Initialize()
}

type Console struct {
	core.Enable
	Level string `mapstructure:"level"`
}

type Logger struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" mapstructure:"filename"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" mapstructure:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" mapstructure:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" mapstructure:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" mapstructure:"localtime"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" mapstructure:"compress"`
}

type EncoderConfig struct {
	// Set the keys used for each log entry. If any key is empty, that portion
	// of the entry is omitted.
	MessageKey    string `json:"messageKey" mapstructure:"messageKey"`
	LevelKey      string `json:"levelKey" mapstructure:"levelKey"`
	TimeKey       string `json:"timeKey" mapstructure:"timeKey"`
	NameKey       string `json:"nameKey" mapstructure:"nameKey"`
	CallerKey     string `json:"callerKey" mapstructure:"callerKey"`
	StacktraceKey string `json:"stacktraceKey" mapstructure:"stacktraceKey"`
	LineEnding    string `json:"lineEnding" mapstructure:"lineEnding"`
	// Configure the primitive representations of common complex types. For
	// example, some users may want all time.Times serialized as floating-point
	// seconds since epoch, while others may prefer ISO8601 strings.
	EncodeLevel    zapcore.LevelEncoder    `json:"levelEncoder" mapstructure:"levelEncoder"`
	EncodeTime     zapcore.TimeEncoder     `json:"timeEncoder" mapstructure:"timeEncoder"`
	EncodeDuration zapcore.DurationEncoder `json:"durationEncoder" mapstructure:"durationEncoder"`
	EncodeCaller   zapcore.CallerEncoder   `json:"callerEncoder" mapstructure:"callerEncoder"`
	// Unlike the other primitive type encoders, EncodeName is optional. The
	// zero value falls back to FullNameEncoder.
	EncodeName zapcore.NameEncoder `json:"nameEncoder" mapstructure:"nameEncoder"`
}

type File struct {
	core.Enable
	Level string `mapstructure:"level"`
}
type Files struct {
	core.Enable
	Files []File `mapstructure:"console"`
}
type Configuration struct {
	Console Console `mapstructure:"console"`
	Files   []File  `mapstructure:"console"`
}

func (this *Configuration) Initialize(params ...interface{}) interface{} {
	err := environment.GetProperty("components.log.zap", &this)
	if err != nil {
		panic(fmt.Sprintf("GComponent [ZAP]read config exception Exit！！！\nException message: %v", err.Error()))
	}
	name := "appName"
	cores := []zapcore.Core{
		consoleCore(),
		jsonCore(),
	}
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	field := zap.Fields(zap.String("appName", name))

	// 构造日志
	logger := zap.New(zapcore.NewTee(cores...)).WithOptions(caller, development, field)
	zap.L().Info("zap.L().Info", zap.Int("balabala", 1))
	zap.ReplaceGlobals(logger)

	logger.Info("logger.Info 初始化成功")
	zap.S().Infof("zap.S().Infof %v", "fdsa")
	return nil
}

func consoleCore() zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.DebugLevel
		}),
	)
}

func jsonCore() zapcore.Core {
	logPath := "xxx.log"
	hook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    128,     // 每个日志文件保存的大小 单位:M
		MaxAge:     7,       // 文件最多保存多少天
		MaxBackups: 30,      // 日志文件最多保存多少个备份
		Compress:   false,   // 是否压缩
	}
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "file",
		StacktraceKey: "stacktrace",
		// 一行结束标识符
		LineEnding: zapcore.DefaultLineEnding,
		// A LevelEncoder serializes a Level to a primitive type.
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		// 日期格式
		EncodeTime: zapcore.ISO8601TimeEncoder,

		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	marshal, _ := json.Marshal(encoderConfig)
	fmt.Println(string(marshal))
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.DebugLevel
		}),
	)
}
