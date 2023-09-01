package util

import "github.com/sirupsen/logrus"

func setFormatter() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
func loggerLogrus() {
	logrus.Debugf("This is a debug message %v", getExampleData())
	logrus.Infof("This is an info message %v", getExampleData())
	logrus.Warnf("This is a warning message %v", getExampleData())
	logrus.Errorf("This is an error message %v", getExampleData())
}

func loggerLogrusWithoutStruct() {
	logrus.Debug("This is a debug message")
	logrus.Info("This is an info message")
	logrus.Warn("This is a warning message")
	logrus.Error("This is an error message")
}
