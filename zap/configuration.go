// Package gzap
// Created by guoxin in 2020/4/13 1:34 下午
package gzap

import (
    "fmt"
    "github.com/GuoxinL/gcomponent/core"
    "github.com/GuoxinL/gcomponent/environment"
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
    "time"
)

var (
    initializeLock = core.NewInitLock()
    instance       *zap.Logger
)

// New Make sure you only initialize it once
func New(params ...interface{}) {
    c := &Configuration{
        InitializeLock: initializeLock,
    }
    c.Initialize(params...)
}

type Console struct {
    core.InitializeLock
    core.BEnable `mapstructure:",squash"`
    Level        string        `mapstructure:"level"`
    Encoder      EncoderConfig `mapstructure:"encoder"`
}

// Logger See lumberjack.Logger
type Logger struct {
    Filename   string `json:"filename" mapstructure:"filename"`
    MaxSize    int    `json:"maxsize" mapstructure:"maxsize"`
    MaxAge     int    `json:"maxage" mapstructure:"maxage"`
    MaxBackups int    `json:"maxbackups" mapstructure:"maxbackups"`
    LocalTime  bool   `json:"localtime" mapstructure:"localtime"`
    Compress   bool   `json:"compress" mapstructure:"compress"`
}

func (l Logger) convert(name string) *lumberjack.Logger {
    logger := defaultLoggerFile(name)
    if len(l.Filename) != 0 {
        logger.Filename = l.Filename
    }
    if l.MaxSize != 0 {
        logger.MaxSize = l.MaxSize
    }
    if l.MaxAge != 0 {
        logger.MaxAge = l.MaxAge
    }
    if l.MaxBackups != 0 {
        logger.MaxBackups = l.MaxBackups
    }
    logger.LocalTime = l.LocalTime
    logger.Compress = l.Compress
    return logger
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
    TimeFormat string `json:"timeFormat" mapstructure:"timeFormat"`
}

func (ec EncoderConfig) convert() zapcore.EncoderConfig {
    encoder := defaultEncoderConfig()
    if len(ec.MessageKey) != 0 {
        encoder.MessageKey = ec.MessageKey
    }
    if len(ec.LevelKey) != 0 {
        encoder.LevelKey = ec.LevelKey
    }
    if len(ec.TimeKey) != 0 {
        encoder.TimeKey = ec.TimeKey
    }
    if len(ec.NameKey) != 0 {
        encoder.NameKey = ec.NameKey
    }
    if len(ec.CallerKey) != 0 {
        encoder.CallerKey = ec.CallerKey
    }
    if len(ec.StacktraceKey) != 0 {
        encoder.StacktraceKey = ec.StacktraceKey
    }
    if len(ec.StacktraceKey) != 0 {
        encoder.StacktraceKey = ec.StacktraceKey
    }
    if ec.LineEnding == "\\n" {
        encoder.LineEnding = zapcore.DefaultLineEnding
    } else {
        encoder.LineEnding = ec.LineEnding
    }
    if len(ec.TimeFormat) != 0 {
        encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(t.Format(ec.TimeFormat))
        }
    }
    return encoder
}
type File struct {
    core.BEnable `mapstructure:",squash"`
    Level        string        `mapstructure:"level"`
    Encoder      EncoderConfig `mapstructure:"encoder"`
    Logger       Logger        `mapstructure:"logger"`
}

type Files struct {
    core.BEnable `mapstructure:",squash"`
    Files        []File `mapstructure:"console"`
}

type Configuration struct {
    core.InitializeLock
    Console       Console `mapstructure:"console"`
    Files         []File  `mapstructure:"files"`
    AddCallerSkip int     `mapstructure:"addCallerSkip"`
}

func (c *Configuration) Initialize(params ...interface{}) interface{} {
    if c.IsInit() {
        return &instance
    }
    environment.New()
    err := environment.GetProperty("components.zap", &c)
    if err != nil {
        panic(fmt.Sprintf("GComponent [zap] read config exception Exit！！！\nException message: %v", err.Error()))
    }

    var cores []zapcore.Core
    // console
    if c.Console.Enable {
        cores = append(cores, writerConsole(
            c.Console.Encoder.convert(),
            func(lvl zapcore.Level) bool {
                return levelFormat(lvl, c.Console.Level)
            },
        ))
    }
    // Files
    if len(c.Files) > 0 {
        for _, file := range c.Files {
            if file.Enable {
                leve := "debug"
                if len(file.Level) != 0 {
                    leve = file.Level
                }
                cores = append(cores, writerJson(
                    file.Encoder.convert(),
                    []zapcore.WriteSyncer{zapcore.AddSync(file.Logger.convert(environment.GetName()))},
                    func(lvl zapcore.Level) bool {
                        return levelFormat(lvl, leve)
                    },
                ))
            }
        }
    }
    // Start the default rules if neither the file nor console output is configured
    if !(len(c.Files) > 0) && !c.Console.Enable {
        cores = append(cores, writerConsole(defaultEncoderConfig(), defaultLevel))
        cores = append(cores, writerJson(defaultEncoderConfig(), []zapcore.WriteSyncer{zapcore.AddSync(defaultLoggerFile("default"))}, defaultLevel))
    }

    // 开启文件及行号
    caller := zap.AddCaller()
    // 开启开发模式，堆栈跟踪
    development := zap.Development()
    zap.NewProductionEncoderConfig()
    // 设置初始化字段
    field := zap.Fields(zap.String("app", environment.GetName()))
    skip := zap.AddCallerSkip(c.AddCallerSkip)

    instance = zap.New(zapcore.NewTee(cores...)).WithOptions(caller, development, field, skip)

    zap.ReplaceGlobals(instance)

    instance.Info("logger.Info 初始化成功")
    zap.L().Info("zap.L().Info", zap.Int("balabala", 1))
    zap.S().Infof("zap.S().Infof %v", "balabala")
    zap.S().With()
    return nil
}

// Default logger encoder
func defaultEncoderConfig() zapcore.EncoderConfig {
    return zapcore.EncoderConfig{
        // Keys can be anything except the empty string.
        MessageKey:     "msg",
        LevelKey:       "level",
        TimeKey:        "time",
        NameKey:        "logger",
        CallerKey:      "file",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalLevelEncoder,
        EncodeTime:     ISOTimeEncoder,
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }
}

// Default logger config
func defaultLoggerFile(name string) *lumberjack.Logger {
    return &lumberjack.Logger{
        Filename:   name + ".log", // 日志文件路径
        MaxSize:    128,           // 每个日志文件保存的大小 单位:M
        MaxAge:     7,             // 文件最多保存多少天
        MaxBackups: 30,            // 日志文件最多保存多少个备份
        Compress:   false,         // 是否压缩
    }
}

// Default logger level debug
func defaultLevel(lvl zapcore.Level) bool {
    return setLevel(lvl, zapcore.DebugLevel)
}

func levelFormat(lvl zapcore.Level, level string) bool {
    var l zapcore.Level
    switch level {
    case "debug", "DEBUG":
        l = zapcore.DebugLevel
    case "info", "INFO", "": // make the zero value useful
        l = zapcore.InfoLevel
    case "warn", "WARN":
        l = zapcore.WarnLevel
    case "error", "ERROR":
        l = zapcore.ErrorLevel
    case "dpanic", "DPANIC":
        l = zapcore.DPanicLevel
    case "panic", "PANIC":
        l = zapcore.PanicLevel
    case "fatal", "FATAL":
        l = zapcore.FatalLevel
    default:
        return false
    }
    return lvl >= l
}

// set logger level
func setLevel(lvl zapcore.Level, level zapcore.Level) bool {
    return lvl >= level
}

func writerJson(cfg zapcore.EncoderConfig, writer []zapcore.WriteSyncer, levelFunc zap.LevelEnablerFunc) zapcore.Core {
    return zapcore.NewCore(
        zapcore.NewJSONEncoder(cfg),
        // 输出多个文件
        zapcore.NewMultiWriteSyncer(writer...),
        // 输出单个文件
        //zapcore.AddSync(defaultLoggerFile(name)),
        levelFunc,
    )
}

func writerConsole(encoder zapcore.EncoderConfig, lvl func(lvl zapcore.Level) bool) zapcore.Core {
    return zapcore.NewCore(
        zapcore.NewConsoleEncoder(encoder),
        zapcore.Lock(os.Stdout),
        zap.LevelEnablerFunc(lvl),
    )
}

func GetInstance() *zap.Logger {
    return instance
}
