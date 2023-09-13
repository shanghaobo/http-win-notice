package log

import (
	"github.com/sirupsen/logrus"
	"http-win-notice/utils"
	"os"
	"time"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})
	logFile, _ := os.OpenFile(utils.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	Logger.SetOutput(logFile)
}

func Println(args ...interface{}) {
	Logger.Println(args)
}

func Info(args ...interface{}) {
	Logger.Info(args)
}

func Debug(args ...interface{}) {
	Logger.Debug(args)
}

func Fatalln(args ...interface{}) {
	Logger.Fatalln(args)
}
