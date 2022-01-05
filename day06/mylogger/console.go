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

func log(lv Level, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, logstring(lv), fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(format string, a ...interface{}) {
	if l.LoggerLevel <= DEBUG {
		log(DEBUG, format, a...)
	}

}

func (l Logger) Info(format string, a ...interface{}) {
	if l.LoggerLevel <= INFO {
		log(INFO, format, a...)
	}
}

func (l Logger) Warning(format string, a ...interface{}) {
	if l.LoggerLevel <= WARNING {
		log(WARNING, format, a...)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	if l.LoggerLevel <= ERROR {
		log(ERROR, format, a...)
	}
}
