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
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})
	logFile, _ := os.OpenFile(utils.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	Logger.SetOutput(logFile)
}

func SetLogLevel(level string) {
	switch level {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
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
