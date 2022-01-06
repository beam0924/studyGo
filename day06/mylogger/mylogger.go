package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type Level uint16

const (
	UNKNOWN Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
)

type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
}

func parseLoggerLevel(s string) (Level, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getInfo(n int) (funcName, fileName string, lineNum int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Call() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fileName = path.Base(file)
	return funcName, fileName, line
}

func logstring(lv Level) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
