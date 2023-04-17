package os

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type E日志类 struct {
	logger *zap.SugaredLogger
}

func New日志类(filepath string, loglevel string) *E日志类 {
	var log E日志类
	v := log.E初始化(filepath, loglevel)
	return v
}

func (this *E日志类) E初始化(filepath string, loglevel string) *E日志类 {
	hook := lumberjack.Logger{
		Filename:   filepath, // ⽇志⽂件路径
		MaxSize:    1024,     // megabytes
		MaxBackups: 3,        // 最多保留3个备份
		MaxAge:     365,      //days
		Compress:   true,     // 是否压缩 disabled by default
	}
	fileWriter := zapcore.AddSync(&hook)
	var highPriority zapcore.Level
	switch loglevel {
	case "debug":
		highPriority = zap.DebugLevel
	case "info":
		highPriority = zap.InfoLevel
	case "error":
		highPriority = zap.ErrorLevel
	default:
		highPriority = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//core := zapcore.NewCore(
	//	zapcore.NewConsoleEncoder(encoderConfig),
	//	fileWriter,
	//	level,
	//)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
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
	//logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	this.logger = zap.New(core).Sugar()
	return this
}

func (this *E日志类) Log(msg string, keysAndValues ...interface{}) {
	//this.logger.Info(msg, fields...)
	this.logger.Infow(msg, keysAndValues...)
}

func (this *E日志类) E错误日志(msg string, keysAndValues ...interface{}) {
	//this.logger.Info(msg, fields...)
	this.logger.Errorw(msg, keysAndValues...)
}
func (this *E日志类) E警告日志(msg string, keysAndValues ...interface{}) {
	//this.logger.Info(msg, fields...)
	this.logger.Warnw(msg, keysAndValues...)
}
func (this *E日志类) E信息日志(msg string, keysAndValues ...interface{}) {
	//this.logger.Info(msg, fields...)
	this.logger.Infow(msg, keysAndValues...)
}
