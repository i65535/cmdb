package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	logger *logrus.Logger
)

func init(){
	if logger == nil{
		LoggerInit()
	}
}

func LoggerInit() {
	mode := "debug"
	logPath := "./cmp.log"
	level := "debug"
	maxDay := 30
	rotationHour := 24

	if err := viper.ReadInConfig(); err == nil {
		mode = viper.GetString("service.mode")
		logPath = viper.GetString("logger.path")
		level = viper.GetString("logger.level")
		maxDay = viper.GetInt("logger.maxDay")
		rotationHour = viper.GetInt("logger.rotationHour")
	}


	log := logrus.New()
	var logDir string
	if mode == gin.ReleaseMode {
		if !FileExists(logPath) {
			logDir, _ = filepath.Split(logPath)
			os.MkdirAll(logDir, os.ModePerm)
		}

		writer, err := rotatelogs.New(
			logPath,
			//rotatelogs.WithLinkName(logFile+time.Now().Format("2006-01-02-03-04-05")), //软连接
			rotatelogs.WithMaxAge(time.Hour*24*time.Duration(maxDay)),                 //文件最大保存时间
			rotatelogs.WithRotationTime(time.Hour*24*time.Duration(rotationHour)),     //文件切割时间
		)
		if err != nil {
			panic(err)
		}
		lfhook := lfshook.NewHook(lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.PanicLevel: writer,
		}, &logrus.JSONFormatter{})

		log.Hooks.Add(lfhook)
	} else {
		log.Out = os.Stdout
	}

	log.Level, _ = logrus.ParseLevel(level)
	logger = log
}

func SetLogLevel(level logrus.Level) {
	logger.Level = level
}
func SetLogFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

type Fields logrus.Fields

// Debug
func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Debug(args...)
	}
}
// Debug
func Debugf(format string, args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Debugf(format, args...)
	}
}

// 带有field的Debug
func DebugWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Debug(l)
	}
}

// Info
func Info(args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Info(args...)
	}
}

// Info
func Infof(format string, args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Infof(format, args...)
	}
}

// 带有field的Info
func InfoWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Info(l)
	}
}

// Warn
func Warn(args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// Warn
func Warnf(format string, args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warnf(format, args...)
	}
}

// 带有Field的Warn
func WarnWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Warn(l)
	}
}

// Error
func Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// Error
func Errorf(format string, args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Errorf(format, args...)
	}
}
// 带有Fields的Error
func ErrorWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
	}
}

// Fatal
func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)
	}
}

// Fatal
func Fatalf(format string, args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatalf(format, args...)
	}
}

// 带有Field的Fatal
func FatalWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(l)
	}
}

// Panic
func Panic(args ...interface{}) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)
	}
}

// Panic
func Panicf(format string, args ...interface{}) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panicf(format, args...)
	}
}

// 带有Field的Panic
func PanicWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Panic(l)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	return fmt.Sprintf("%s:%d", file, line)
}