package logger

import(
	"github.com/sirupsen/logrus"
	"os"
)

const  logfile = "logs.log"


func GetLogger() *logrus.Logger{
	var logger = logrus.New()
	logfile, _ := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	logger.Out = logfile
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.ErrorLevel)
	return logger
	
}