package utils

import (
	"io"
	"log"
	"os"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	// 指定日志输出格式
	LogFormat string `yaml:"logFormat"`
	// 指定日志输出目录
	LogDir string `yaml:"logDir"`
	// 日志保留天数
	LogKeepDays int `yaml:"logKeepDays"`
}

var file *os.File
var logtime string

// 日志处理器
var Log = GetLogConfig()

func (log *Logger) Init() {
	// logrus.SetReportCaller(true)

	if log.LogFormat == "text" {
		logrus.SetFormatter(&nested.Formatter{
			HideKeys:        true,
			TimestampFormat: "15:04:05",
			FieldsOrder:     []string{"component", "level", "msg"},
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	// 同步输出到文件
	if log.LogDir != "" {
		log.LogDir = "./log"
		logtime = time.Now().Format("2006-01-02")
		// 输出日志到文件
		file, err := os.OpenFile(Log.LogDir+"/"+logtime+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Panic("日志文件打开失败，系统日志将无法正常生成：", err)
		}
		opt := io.MultiWriter(os.Stdout, file)
		logrus.SetOutput(opt)
	}
	go log.CleanLog()
}

// 定时清理日志
func (log *Logger) CleanLog() {
	go func() {
		for {
			time.Sleep(time.Hour * 24)
			log.Clean()
		}
	}()
}

// 清理日志
func (log *Logger) Clean() {
	now := time.Now()
	lastday := now.AddDate(0, 0, -(log.LogKeepDays))
	lastdayDate := lastday.Format("2006-01-02")
	err := os.Remove(Log.LogDir + "/" + lastdayDate + ".log")
	if err != nil {
		log.Panic("日志文件删除失败：", err)
	}
}

// 日志处理器
func LogFileCheck() {
	// 检查时间
	if time.Now().Format("2006-01-02") != logtime {
		logtime = time.Now().Format("2006-01-02")
		file.Close()
		file, err := os.OpenFile(Log.LogDir+"/"+logtime+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Panicln("日志文件切换失败，系统日志将无法正常生成，原因：", err)
		}
		opt := io.MultiWriter(os.Stdout, file)
		logrus.SetOutput(opt)
	}
}

// 警告
func (log *Logger) Warning(args ...interface{}) {
	LogFileCheck()
	logrus.Warning(args...)
}

// 错误
func (log *Logger) Error(args ...interface{}) {
	LogFileCheck()
	logrus.Error(args...)
}

// 信息
func (log *Logger) Info(args ...interface{}) {
	LogFileCheck()
	logrus.Info(args...)
}

// 调试
func (log *Logger) Debug(args ...interface{}) {
	LogFileCheck()
	logrus.Debug(args...)
}

// 跟踪
func (log *Logger) Trace(args ...interface{}) {
	LogFileCheck()
	logrus.Trace(args...)
}

// 致命
func (log *Logger) Fatal(args ...interface{}) {
	LogFileCheck()
	logrus.Fatal(args...)
}

// 紧急
func (log *Logger) Panic(args ...interface{}) {
	LogFileCheck()
	logrus.Panic(args...)
}
