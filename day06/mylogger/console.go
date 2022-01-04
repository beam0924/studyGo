package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// 往终端写

type Level uint16

// Logger日志类
type Logger struct {
	LoggerLevel Level
}

const (
	UNKNOWN Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
)

func parseLoggerLevel(s string) (Level, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARNNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

// 日志构造函数
func NewLog(LoggerLevel string) Logger {
	Level, err := parseLoggerLevel(LoggerLevel)
	if err != nil {
		panic(err)
	}
	return Logger{Level}
}

func (l Logger) Debug(s string) {
	if l.LoggerLevel <= DEBUG {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [DEBUG] %s\n", now, s)
	}

}

func (l Logger) Info(s string) {
	if l.LoggerLevel <= INFO {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [INFO] %s\n", now, s)
	}
}

func (l Logger) Warning(s string) {
	if l.LoggerLevel <= WARNING {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [WARNING] %s\n", now, s)
	}
}

func (l Logger) Error(s string) {
	if l.LoggerLevel <= ERROR {
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [ERROR] %s\n", now, s)
	}
}
