package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger(mode string) {
	var config zap.Config

	if mode == "prod" {
		// 生产环境：输出到 JSON，适合日志采集系统
		config = zap.NewProductionConfig()
	} else {
		// 开发环境：更易读的控制台输出
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 日志级别带颜色
	}

	config.OutputPaths = []string{"stdout"} // 输出到标准输出（Docker 会捕获）

	baseLogger, _ := config.Build()
	Log = baseLogger.Sugar() // Sugar 提供更方便的接口，如 Log.Infof(...)
}
