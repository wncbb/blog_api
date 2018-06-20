package log

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
}

func Test_Init(t *testing.T) {
	cfgMap := make(map[string]*LogConfig)
	cfgMap["a"] = &LogConfig{
		Writer: os.Stdout,
		Level:  logrus.DebugLevel,
	}
	Init(cfgMap)
	GetLog("a").Info("LINE30")

}

func Test_A(t *testing.T) {
	logger := DefaultLog()
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	logger.Fatal("fatal")

}
