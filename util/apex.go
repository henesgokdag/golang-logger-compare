package util

import (
	apex_logger "github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"os"
)

func setHandler() {
	apex_logger.SetHandler(json.New(os.Stdout))
}
func loggerApex() {
	apex_logger.Debugf("This is a debug message %v", getExampleData())
	apex_logger.Infof("This is an info message %v", getExampleData())
	apex_logger.Warnf("This is a warning message %v", getExampleData())
	apex_logger.Errorf("This is an error message %v", getExampleData())
}
func loggerApexWithoutStruct() {
	apex_logger.Debug("This is a debug message")
	apex_logger.Info("This is an info message")
	apex_logger.Warn("This is a warning message")
	apex_logger.Error("This is an error message")
}
