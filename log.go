package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type TDLog struct {
	log *logrus.Logger
}

var logMap map[string]*TDLog
var getLogMutex sync.Mutex

// GetLogger gets logrus.Logger object according to module name
// each module can have its own logger
func GetLogger(module string) *TDLog {
	getLogMutex.Lock()
	defer getLogMutex.Unlock()
	if logMap == nil {
		logMap = make(map[string]*TDLog)
	}
	curLog, ok := logMap[module]
	if ok {
		return curLog
	}

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05.000",
		FullTimestamp:    true,
		DisableTimestamp: false,
		ForceColors:      true,
	})

	if LogConfiguration.PrintLog {
		log.Out = os.Stdout
	} else {
		logDir := LogConfiguration.DataDir
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("failed to create log dir: %s", err.Error()))
		}
		logFileName := fmt.Sprintf("%s%s", "%Y%m%d%H", LogConfiguration.LogExtension)

		writer, err := rotatelogs.New(
			filepath.Join(logDir, logFileName),
			rotatelogs.WithClock(rotatelogs.Local),
			rotatelogs.WithMaxAge(LogConfiguration.MaxAge),
			rotatelogs.WithRotationTime(LogConfiguration.RotationTime),
		)

		if err != nil {
			panic(fmt.Sprintf("failed to create log file: %s", err))
		}

		log.Out = writer
	}

	if LogConfiguration.IsDebug {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	log.AddHook(&CallerHook{module: module}) // add caller hook to print caller's file and line number
	curLog = &TDLog{
		log: log,
	}
	logMap[module] = curLog
	return curLog
}
