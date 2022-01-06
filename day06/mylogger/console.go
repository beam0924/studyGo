package mylogger

import (
	"fmt"
	"time"
)

// 往终端写

// Logger日志类
type consoleLogger struct {
	LoggerLevel Level
}

// 日志构造函数
func NewConsoleLog(LoggerLevel string) consoleLogger {
	Level, err := parseLoggerLevel(LoggerLevel)
	if err != nil {
		panic(err)
	}
	return consoleLogger{Level}
}

func log(lv Level, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, logstring(lv), fileName, funcName, lineNo, msg)
}

func (c consoleLogger) Debug(format string, a ...interface{}) {
	if c.LoggerLevel <= DEBUG {
		log(DEBUG, format, a...)
	}

}

func (c consoleLogger) Info(format string, a ...interface{}) {
	if c.LoggerLevel <= INFO {
		log(INFO, format, a...)
	}
}

func (c consoleLogger) Warning(format string, a ...interface{}) {
	if c.LoggerLevel <= WARNING {
		log(WARNING, format, a...)
	}
}

func (c consoleLogger) Error(format string, a ...interface{}) {
	if c.LoggerLevel <= ERROR {
		log(ERROR, format, a...)
	}
}
