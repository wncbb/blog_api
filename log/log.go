package log

import (
	"io"
	"os"
	"runtime"
	"time"

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

func DefaultLogError(id string, msg string, err error) {
	_, file, line, _ := runtime.Caller(1)
	LogMap[defaultKey].Errorf(
		"ts:`%s`, id:`%s`, location:`%s@%d`, msg:`%s`, err:`%v`",
		time.Now().Format("2006-01-02 15:04:05"),
		id,
		file,
		line,
		msg,
		err,
	)
}

func DefaultLogDebug(id string, msg string) {
	_, file, line, _ := runtime.Caller(1)
	LogMap[defaultKey].Debugf(
		"ts: `%s`, id:`%s`, location:`%s@%d`, msg:`%s`",
		time.Now().Format("2006-01-02 15:04:05"),
		file,
		line,
		msg,
	)
}
