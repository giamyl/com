package blog

import (
	"fmt"
	"log"
	"os"
)

const (
	logLevelOff = 1 << iota
	logLevelFatal
	logLevelError
	logLevelWarn
	logLevelInfo
	logLevelDebug
)

const (
	levelOff   = logLevelOff
	levelFatal = levelOff | logLevelFatal
	levelError = levelFatal | logLevelError
	levelWarn  = levelError | logLevelWarn
	levelInfo  = levelWarn | logLevelInfo
	levelDebug = levelInfo | logLevelDebug
)

var text = map[string]int{
	"off":   levelOff,
	"fatal": levelFatal,
	"error": levelError,
	"warn":  levelWarn,
	"info":  levelInfo,
	"debug": levelDebug,
}

var level = newLevel()
var defautWriter = os.Stderr

func newLevel() int {

	def := text["debug"]

	if lv := os.Getenv("LOG_LEVEL"); lv != "" {
		if l, ok := text[lv]; ok {
			def = l
		}
	} else {
		if mode := os.Getenv("DEBUG"); mode == "release" {
			def = text["info"]
		}
	}

	return def

}

var logger = log.New(defautWriter, "[debug]", log.LstdFlags|log.Llongfile)
var loggerErr = log.New(defautWriter, "[error]", log.LstdFlags|log.Llongfile)
var loggerWarn = log.New(defautWriter, "[warn]", log.LstdFlags|log.Lshortfile)
var loggerInfo = log.New(defautWriter, "[info]", log.LstdFlags)

// Debug 模式的日志
func Debug(values ...interface{}) {
	if level&logLevelDebug > 0 {
		logger.Output(2, fmt.Sprintln(values...))
	}
}

// Debugf 指定格式的日志
func Debugf(format string, values ...interface{}) {
	if level&logLevelDebug > 0 {
		logger.Output(2, fmt.Sprintf(format, values...))
	}
}

// DebugOutput 自定义日志深度和格式
func DebugOutput(calldepth int, format string, values ...interface{}) {
	if level&logLevelDebug > 0 {
		logger.Output(calldepth+2, fmt.Sprintf(format, values...))
	}
}

// Error 模式的日志
func Error(values ...interface{}) {
	if level&logLevelError > 0 {
		loggerErr.Output(2, fmt.Sprintln(values...))
	}
}

// Errorf 指定格式的错误日志
func Errorf(format string, values ...interface{}) {
	if level&logLevelError > 0 {
		loggerErr.Output(2, fmt.Sprintf(format, values...))
	}
}

// ErrorOutput 自定义日志深度和格式
func ErrorOutput(calldepth int, format string, values ...interface{}) {
	if level&logLevelError > 0 {
		loggerErr.Output(calldepth+2, fmt.Sprintf(format, values...))
	}
}

// Info 模式的日志
func Info(values ...interface{}) {
	if level&logLevelInfo > 0 {
		loggerInfo.Output(2, fmt.Sprintln(values...))
	}
}

// Infof 指定格式的错误日志
func Infof(format string, values ...interface{}) {
	if level&logLevelInfo > 0 {
		loggerInfo.Output(2, fmt.Sprintf(format, values...))
	}
}

// InfoOutput 自定义日志深度和格式
func InfoOutput(calldepth int, format string, values ...interface{}) {
	if level&logLevelInfo > 0 {
		loggerInfo.Output(calldepth+2, fmt.Sprintf(format, values...))
	}
}

// Warn 模式的日志
func Warn(values ...interface{}) {
	if level&logLevelWarn > 0 {
		loggerWarn.Output(2, fmt.Sprintln(values...))
	}
}

// Warnf 指定格式的错误日志
func Warnf(format string, values ...interface{}) {
	if level&logLevelWarn > 0 {
		loggerWarn.Output(2, fmt.Sprintf(format, values...))
	}
}

// WarnOutput 自定义日志深度和格式
func WarnOutput(calldepth int, format string, values ...interface{}) {
	if level&logLevelWarn > 0 {
		loggerWarn.Output(calldepth+2, fmt.Sprintf(format, values...))
	}
}
