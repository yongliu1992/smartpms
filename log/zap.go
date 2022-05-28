package log

import (
	"github.com/natefinch/lumberjack/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
import "time"

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

var log *zap.SugaredLogger

func GetLog() *zap.SugaredLogger {
	return log
}
func InitZap(appName string) {
	core := zapcore.NewCore(getEncoder(), getLogW(appName), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	log = logger.Sugar()
}

func getLogW(appName string) zapcore.WriteSyncer {
	l, _ := lumberjack.NewRoller(
		"/tmp/"+appName+".log",
		100*3, // 500 megabytes
		&lumberjack.Options{
			MaxBackups: 3,
			MaxAge:     28 * time.Hour * 24, // 28 days
			Compress:   true,
		})

	return zapcore.AddSync(l)
}
