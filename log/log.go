package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

const logFormat = `date=%s, method=%s, url=%s,  response_time=%s`

// Config of log
type Config struct {
	LogLevel     string
	ErrorLogPath string
	AppName      string
}

var (
	// DefaultConfig of log
	DefaultConfig = Config{
		LogLevel:     "info",
		ErrorLogPath: "/var/log",
		AppName:      "unknown",
	}
)

// SetConfig for log object
func SetConfig(config Config) {
	// set default config if settings not there
	if config.LogLevel == "" {
		config.LogLevel = DefaultConfig.LogLevel
	}
	if config.ErrorLogPath == "" {
		config.ErrorLogPath = DefaultConfig.ErrorLogPath + fmt.Sprintf("/%s/%s.error.log", config.AppName, config.AppName)
	}
	SetLevel(config.LogLevel)
	if f := reopen(1, config.ErrorLogPath); f != nil {
		logrus.SetOutput(f)
	}
}

func reopen(fd int, filename string) *os.File {
	if filename == "" {
		return nil
	}

	logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// logrus.Error("Error in opening ", filename, err)
		return nil
	}
	return logFile
}

// Fields for logrus fields
type Fields logrus.Fields

// SetLevel of logs
func SetLevel(level string) {
	switch level {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

// GetLevel get log level
func GetLevel() string {
	return strings.ToUpper(logrus.GetLevel().String())
}

// Info log
func Info(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Info(args...)
}

// Infoln log
func Infoln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infoln(args...)
}

// Infof log
func Infof(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infof(format, args...)
}

// Print log
func Print(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Info(args...)
}

// Println log
func Println(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infoln(args...)
}

// Printf log
func Printf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infof(format, args...)
}

// Debug log
func Debug(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Debug(args...)
}

// Debugln log
func Debugln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Debugln(args...)
}

// Debugf log
func Debugf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Debugf(format, args...)
}

// Warn log
func Warn(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Warn(args...)
}

func Warnln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Warnln(args...)
}

func Warnf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Warnf(format, args...)
}

func Error(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Error(args...)
}

func Errorln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Fatal(args...)
}

func Fatalln(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Fatalf(format, args...)
}

func WithFields(fields Fields) *logrus.Entry {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}

	fields["source"] = fmt.Sprintf("%s:%d", file, line)

	logrusFields := logrus.Fields{}

	for key, value := range fields {
		logrusFields[key] = value
	}

	return logrus.WithFields(logrusFields)
}
