package log

import (
	"fmt"
	"log"
	"os"
)

var logpath = "log/service.log"
var defaultLoger *log.Logger

func init() {
	logFile, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open log file failed")
	}
	defaultLoger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
}

func Info(v ...interface{}) {
	defaultLoger.SetPrefix("[Info]")
	defaultLoger.Output(2, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	defaultLoger.SetPrefix("[Info]")
	defaultLoger.Output(2, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	defaultLoger.SetPrefix("[Error]")
	defaultLoger.Output(2, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	defaultLoger.SetPrefix("[Error]")
	defaultLoger.Output(2, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	defaultLoger.SetPrefix("[Fatal]")
	defaultLoger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	defaultLoger.SetPrefix("[Fatal]")
	defaultLoger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
