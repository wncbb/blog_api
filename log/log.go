package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	defaultKey string = "default"
)

type LogConfig struct {
	Writer io.Writer
	Level  logrus.Level
}

var LogMap map[string]*logrus.Logger

func Init(configMap map[string]*LogConfig) {
	for k, v := range configMap {
		lpLog := logrus.New()
		lpLog.Out = configMap[k].Writer
		lpLog.SetLevel(v.Level)
		LogMap[k] = lpLog
	}
}

func init() {
	LogMap = make(map[string]*logrus.Logger)

	log := logrus.New()
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)

	LogMap[defaultKey] = log
}

func GetLog(key string) *logrus.Logger {
	retLog, ok := LogMap[key]
	if !ok {
		return LogMap[defaultKey]
	}
	return retLog
}

func DefaultLog() *logrus.Logger {
	return LogMap[defaultKey]
}
