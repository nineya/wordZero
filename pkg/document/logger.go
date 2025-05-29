// Package document 日志系统
package document

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// LogLevel 日志级别
type LogLevel int

const (
	// LogLevelDebug 调试级别
	LogLevelDebug LogLevel = iota
	// LogLevelInfo 信息级别
	LogLevelInfo
	// LogLevelWarn 警告级别
	LogLevelWarn
	// LogLevelError 错误级别
	LogLevelError
	// LogLevelSilent 静默级别
	LogLevelSilent
)

// String 返回日志级别的字符串表示
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelSilent:
		return "SILENT"
	default:
		return "UNKNOWN"
	}
}

// Logger 日志记录器
type Logger struct {
	level  LogLevel    // 日志级别
	output io.Writer   // 输出目标
	logger *log.Logger // 内部日志记录器
}

// defaultLogger 默认全局日志记录器
var defaultLogger = NewLogger(LogLevelInfo, os.Stdout)

// NewLogger 创建新的日志记录器
func NewLogger(level LogLevel, output io.Writer) *Logger {
	return &Logger{
		level:  level,
		output: output,
		logger: log.New(output, "", 0),
	}
}

// SetLevel 设置日志级别
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// SetOutput 设置输出目标
func (l *Logger) SetOutput(output io.Writer) {
	l.output = output
	l.logger.SetOutput(output)
}

// logf 格式化日志输出
func (l *Logger) logf(level LogLevel, format string, args ...interface{}) {
	if l.level > level {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	l.logger.Printf("[%s] %s - %s", timestamp, level.String(), message)
}

// Debugf 输出调试日志
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logf(LogLevelDebug, format, args...)
}

// Infof 输出信息日志
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logf(LogLevelInfo, format, args...)
}

// Warnf 输出警告日志
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logf(LogLevelWarn, format, args...)
}

// Errorf 输出错误日志
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logf(LogLevelError, format, args...)
}

// Debug 输出调试日志
func (l *Logger) Debug(msg string) {
	l.Debugf("%s", msg)
}

// Info 输出信息日志
func (l *Logger) Info(msg string) {
	l.Infof("%s", msg)
}

// Warn 输出警告日志
func (l *Logger) Warn(msg string) {
	l.Warnf("%s", msg)
}

// Error 输出错误日志
func (l *Logger) Error(msg string) {
	l.Errorf("%s", msg)
}

// 全局日志函数，使用默认日志记录器

// SetGlobalLevel 设置全局日志级别
func SetGlobalLevel(level LogLevel) {
	defaultLogger.SetLevel(level)
}

// SetGlobalOutput 设置全局日志输出
func SetGlobalOutput(output io.Writer) {
	defaultLogger.SetOutput(output)
}

// Debugf 全局调试日志
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Infof 全局信息日志
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warnf 全局警告日志
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Errorf 全局错误日志
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Debug 全局调试日志
func Debug(msg string) {
	defaultLogger.Debug(msg)
}

// Info 全局信息日志
func Info(msg string) {
	defaultLogger.Info(msg)
}

// Warn 全局警告日志
func Warn(msg string) {
	defaultLogger.Warn(msg)
}

// Error 全局错误日志
func Error(msg string) {
	defaultLogger.Error(msg)
}
