/*
   Created by guoxin in 2020/4/13 1:34 下午
*/
package log_zap

import (
	"fmt"
	"github.com/GuoxinL/gcomponent/environment"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	new(Configuration).Initialize()
}

type Enable struct {
	Enable bool `yaml:"enable"`
}

type Console struct {
	Enable
	Level string `yaml:"level"`
}

type File struct {
	Enable
	Level string `yaml:"level"`
}

type Configuration struct {
	Console Console `yaml:"console"`
	Files   []File  `yaml:"console"`
}

func (this *Configuration) Initialize(params ...interface{}) interface{} {
	err := environment.GetProperty("components.log.zap", &this)
	if err != nil {
		panic(fmt.Sprintf("GComponent [ZAP]读取配置异常, 退出程序！！！\n异常信息: %v", err.Error()))
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
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.DebugLevel
		}),
	)
}
