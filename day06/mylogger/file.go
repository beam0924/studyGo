package mylogger

// 往文件里面写日志相关的代码
import (
	"fmt"
	"os"
	"time"
)

// 往终端写

// Logger日志类
type fileLogger struct {
	LoggerLevel Level
}

// 日志构造函数
func NewFileLog(LoggerLevel string) fileLogger {
	Level, err := parseLoggerLevel(LoggerLevel)
	if err != nil {
		panic(err)
	}
	return fileLogger{Level}
}

func fileLog(lv Level, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now().Format("2006-01-02 15:04:05")
	logName := time.Now().Format("2006-01-02")
	logName = logName + ".log"
	funcName, fileName, lineNo := getInfo(3)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open logfile err, err:%v\n", err)
		return
	}
	defer fileObj.Close()
	logLine := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", now, logstring(lv), fileName, funcName, lineNo, msg)
	fileObj.WriteString(logLine)
}

func (l fileLogger) Debug(format string, a ...interface{}) {
	if l.LoggerLevel <= DEBUG {
		fileLog(DEBUG, format, a...)
	}

}

func (l fileLogger) Info(format string, a ...interface{}) {
	if l.LoggerLevel <= INFO {
		fileLog(INFO, format, a...)
	}
}

func (l fileLogger) Warning(format string, a ...interface{}) {
	if l.LoggerLevel <= WARNING {
		fileLog(WARNING, format, a...)
	}
}

func (l fileLogger) Error(format string, a ...interface{}) {
	if l.LoggerLevel <= ERROR {
		fileLog(ERROR, format, a...)
	}
}
