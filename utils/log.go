package utils

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	// 指定日志输出格式
	LogFormat string `yaml:"logFormat"`
	// 指定日志输出目录
	LogDir string `yaml:"logDir"`
}

var Log = GetLogConfig()

func (log *Logger) Init() {
	logrus.SetLevel(logrus.FatalLevel)
	if log.LogFormat == "text" {
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	// 同步输出到文件
	if log.LogDir != "" {
		// 输出日志到文件
		file, _ := os.OpenFile(log.LogDir+time.Now().String()+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		opt := io.MultiWriter(os.Stdout, file)
		logrus.SetOutput(opt)
	}
}

// 警告
func (log *Logger) Warning(msg string) {
	logrus.Warning(msg)
}

// 错误
func (log *Logger) Error(msg string) {
	logrus.Error(msg)
}

// 信息
func (log *Logger) Info(msg string) {
	logrus.Info(msg)
}

// 调试
func (log *Logger) Debug(msg string) {
	logrus.Debug(msg)
}

// 跟踪
func (log *Logger) Trace(msg string) {
	logrus.Trace(msg)
}

// 致命
func (log *Logger) Fatal(msg string) {
	logrus.Fatal(msg)
}

// 紧急
func (log *Logger) Panic(msg string) {
	logrus.Panic(msg)
}
