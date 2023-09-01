package util

import "go.uber.org/zap"
import "log"

func getZapSugarLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	return logger
}
func loggerZapSugar(logger *zap.Logger) {
	sugar := logger.Sugar()

	sugar.Debugf("this is a debug message %v", getExampleData())
	sugar.Infof("this is an info message %v", getExampleData())
	sugar.Warnf("this is a warn message %v", getExampleData())
	sugar.Errorf("this is an error message %v", getExampleData())
}
func loggerZapSugarWithoutStruct(logger *zap.Logger) {
	sugar := logger.Sugar()

	sugar.Debug("this is a debug message")
	sugar.Info("this is an info message")
	sugar.Warn("this is a warn message")
	sugar.Error("this is an error message")
}
func loggerZapwWithoutStruct(logger *zap.Logger) {
	logger.Debug("this is a debug message")
	logger.Info("this is an info message")
	logger.Warn("this is a warn message")
	logger.Error("this is an error message")
}
