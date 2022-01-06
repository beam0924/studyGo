package mylogger

// 往文件里面写日志相关的代码
import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 往终端写

// Logger日志类
type fileLogger struct {
	LoggerLevel Level
	filePath    string
	maxFileSize int64
}

// 日志构造函数
func NewFileLog(LoggerLevel string, filePath string, maxFileSize int64) *fileLogger {
	Level, err := parseLoggerLevel(LoggerLevel)
	if err != nil {
		panic(err)
	}
	return &fileLogger{
		LoggerLevel: Level,
		filePath:    filePath,
		maxFileSize: maxFileSize,
	}
}

func (f *fileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
	}
	if fileInfo.Size() > f.maxFileSize {
		return true
	} else {
		return false
	}
}

func (f *fileLogger) fileLog(lv Level, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now().Format("2006-01-02 15:04:05")
	var prefix string
	if strings.HasSuffix(f.filePath, "/") {
		prefix = f.filePath
	} else {
		prefix = f.filePath + "/"
	}
	funcName, fileName, lineNo := getInfo(3)
	os.MkdirAll(prefix, 0766)
	logName := time.Now().Format("2006-01-02")
	logName = prefix + logName + ".log"
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open logfile err, err:%v\n", err)
		return
	}
	if f.checkSize(fileObj) {
		//关闭当前日志文件
		fileObj.Close()
		//备份原来的日志文件
		nowstr := time.Now().Format("150405")
		newLogName := fmt.Sprintf("%s.bak%s", logName, nowstr)
		os.Rename(logName, newLogName)
		//打开一个新的日志文件
		//将打开的新日志文件赋值给 fileObj
		fileObj, err = os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("open new logfile err, err:%v\n", err)
			return
		}

	}
	defer fileObj.Close()
	logLine := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", now, logstring(lv), fileName, funcName, lineNo, msg)
	fileObj.WriteString(logLine)
}

func (f *fileLogger) Debug(format string, a ...interface{}) {
	if f.LoggerLevel <= DEBUG {
		f.fileLog(DEBUG, format, a...)
	}
}

func (f *fileLogger) Info(format string, a ...interface{}) {
	if f.LoggerLevel <= INFO {
		f.fileLog(INFO, format, a...)
	}
}

func (f *fileLogger) Warning(format string, a ...interface{}) {
	if f.LoggerLevel <= WARNING {
		f.fileLog(WARNING, format, a...)
	}
}

func (f *fileLogger) Error(format string, a ...interface{}) {
	if f.LoggerLevel <= ERROR {
		f.fileLog(ERROR, format, a...)
	}
}
