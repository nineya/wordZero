package markdown

import (
	"errors"
	"fmt"
)

var (
	// ErrUnsupportedMarkdown 不支持的Markdown语法
	ErrUnsupportedMarkdown = errors.New("unsupported markdown syntax")

	// ErrInvalidImagePath 无效的图片路径
	ErrInvalidImagePath = errors.New("invalid image path")

	// ErrFileNotFound 文件未找到
	ErrFileNotFound = errors.New("file not found")

	// ErrInvalidMarkdown 无效的Markdown内容
	ErrInvalidMarkdown = errors.New("invalid markdown content")

	// ErrConversionFailed 转换失败
	ErrConversionFailed = errors.New("conversion failed")

	// ErrUnsupportedWordElement 不支持的Word元素
	ErrUnsupportedWordElement = errors.New("unsupported word element")

	// ErrExportFailed 导出失败
	ErrExportFailed = errors.New("export failed")

	// ErrInvalidDocument 无效的Word文档
	ErrInvalidDocument = errors.New("invalid word document")
)

// ConversionError 转换错误，包含详细信息
type ConversionError struct {
	Type    string // 错误类型
	Message string // 错误消息
	Line    int    // 错误行号（如果适用）
	Column  int    // 错误列号（如果适用）
	Cause   error  // 原始错误
}

// Error 实现error接口
func (e *ConversionError) Error() string {
	if e.Line > 0 {
		return fmt.Sprintf("%s at line %d, column %d: %s", e.Type, e.Line, e.Column, e.Message)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// Unwrap 返回原始错误，支持errors.Unwrap
func (e *ConversionError) Unwrap() error {
	return e.Cause
}

// NewConversionError 创建新的转换错误
func NewConversionError(errorType, message string, line, column int, cause error) *ConversionError {
	return &ConversionError{
		Type:    errorType,
		Message: message,
		Line:    line,
		Column:  column,
		Cause:   cause,
	}
}

// ExportError 导出错误，包含详细信息
type ExportError struct {
	Type    string // 错误类型
	Message string // 错误消息
	Cause   error  // 原始错误
}

// Error 实现error接口
func (e *ExportError) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// Unwrap 返回原始错误，支持errors.Unwrap
func (e *ExportError) Unwrap() error {
	return e.Cause
}

// NewExportError 创建新的导出错误
func NewExportError(errorType, message string, cause error) *ExportError {
	return &ExportError{
		Type:    errorType,
		Message: message,
		Cause:   cause,
	}
}
