package mylogger

import (
	"fmt"
	"time"
)

// 往终端写

// Logger日志类
type Logger struct {
	LoggerLevel Level
}

// 日志构造函数
func NewLog(LoggerLevel string) Logger {
	Level, err := parseLoggerLevel(LoggerLevel)
	if err != nil {
		panic(err)
	}
	return Logger{Level}
}

func log(lv Level, msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, logstring(lv), fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(s string) {
	if l.LoggerLevel <= DEBUG {
		log(DEBUG, s)
	}

}

func (l Logger) Info(s string) {
	if l.LoggerLevel <= INFO {
		log(INFO, s)
	}
}

func (l Logger) Warning(s string) {
	if l.LoggerLevel <= WARNING {
		log(WARNING, s)
	}
}

func (l Logger) Error(s string) {
	if l.LoggerLevel <= ERROR {
		log(ERROR, s)
	}
}
