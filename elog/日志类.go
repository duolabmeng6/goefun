// Package elog 日志模块
package elog

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
)

type E日志类 struct {
    Logger    *zap.SugaredLogger
    LoggerObj *zap.Logger
}

// New日志类 创建一个日志实例
// 日志文件路径用于输出到文件，日志级别支持 debug、info、error
func New日志类(日志文件路径 string, 日志级别 string) *E日志类 {
    var log E日志类
    v := log.E初始化(日志文件路径, 日志级别)
    return v
}

// E初始化 初始化日志对象
// 配置输出格式、级别以及文件滚动策略
func (this *E日志类) E初始化(日志文件路径 string, 日志级别 string) *E日志类 {
    hook := lumberjack.Logger{
        Filename:   日志文件路径, // ⽇志⽂件路径
        MaxSize:    1024,   // megabytes
        MaxBackups: 3,      // 最多保留3个备份
        MaxAge:     365,    //days
        Compress:   true,   // 是否压缩 disabled by default
    }
    fileWriter := zapcore.AddSync(&hook)
    var highPriority zapcore.Level
    switch 日志级别 {
    case "debug":
        highPriority = zap.DebugLevel
    case "info":
        highPriority = zap.InfoLevel
    case "error":
        highPriority = zap.ErrorLevel
    default:
        highPriority = zap.InfoLevel
    }

    //配置日志的格式
    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    encoderConfig.TimeKey = "time"
    encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
    encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
    encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
    encoderConfig.CallerKey = "caller"
    consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)

    //consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
    consoleDebugging := zapcore.Lock(os.Stdout)
    lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl >= zapcore.DebugLevel
    })
    core := zapcore.NewTee(
        // 打印在控制台
        zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
        // 打印在文件中
        zapcore.NewCore(consoleEncoder, fileWriter, highPriority),
    )

    //代码的位置也可以输出
    //Logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
    this.LoggerObj = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel))
    this.Logger = this.LoggerObj.Sugar()

    // 替换全局log
    zap.ReplaceGlobals(this.LoggerObj)
    return this
}

// Log 记录信息级别日志
// msg 为消息体，keysAndValues 为成对的结构化键值
func (this *E日志类) Log(msg string, keysAndValues ...interface{}) {
    //this.Logger.Info(msg, fields...)
    this.Logger.Infow(msg, keysAndValues...)
}

// E错误日志 记录错误级别日志
func (this *E日志类) E错误日志(msg string, keysAndValues ...interface{}) {
    //this.Logger.Info(msg, fields...)
    this.Logger.Errorw(msg, keysAndValues...)
}
// E警告日志 记录警告级别日志
func (this *E日志类) E警告日志(msg string, keysAndValues ...interface{}) {
    //this.Logger.Info(msg, fields...)
    this.Logger.Warnw(msg, keysAndValues...)
}
// E信息日志 记录信息级别日志（同 Log）
func (this *E日志类) E信息日志(msg string, keysAndValues ...interface{}) {
    //this.Logger.Info(msg, fields...)
    this.Logger.Infow(msg, keysAndValues...)
}
