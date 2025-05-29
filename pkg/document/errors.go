// Package document 错误处理
package document

import (
	"errors"
	"fmt"
)

// 预定义错误类型
var (
	// ErrInvalidDocument 无效文档
	ErrInvalidDocument = errors.New("invalid document")

	// ErrDocumentNotFound 文档未找到
	ErrDocumentNotFound = errors.New("document not found")

	// ErrInvalidFormat 无效格式
	ErrInvalidFormat = errors.New("invalid format")

	// ErrCorruptedFile 文件损坏
	ErrCorruptedFile = errors.New("corrupted file")

	// ErrUnsupportedOperation 不支持的操作
	ErrUnsupportedOperation = errors.New("unsupported operation")
)

// DocumentError 文档操作错误
type DocumentError struct {
	Operation string // 操作名称
	Cause     error  // 原因
	Context   string // 上下文信息
}

// Error 实现error接口
func (e *DocumentError) Error() string {
	if e.Context != "" {
		return fmt.Sprintf("document operation failed: %s (%s): %v", e.Operation, e.Context, e.Cause)
	}
	return fmt.Sprintf("document operation failed: %s: %v", e.Operation, e.Cause)
}

// Unwrap 解包错误，支持errors.Is和errors.As
func (e *DocumentError) Unwrap() error {
	return e.Cause
}

// NewDocumentError 创建新的文档错误
func NewDocumentError(operation string, cause error, context string) *DocumentError {
	return &DocumentError{
		Operation: operation,
		Cause:     cause,
		Context:   context,
	}
}

// WrapError 包装错误，添加操作上下文
func WrapError(operation string, err error) error {
	if err == nil {
		return nil
	}
	return NewDocumentError(operation, err, "")
}

// WrapErrorWithContext 包装错误，添加操作和上下文信息
func WrapErrorWithContext(operation string, err error, context string) error {
	if err == nil {
		return nil
	}
	return NewDocumentError(operation, err, context)
}

// ValidationError 验证错误
type ValidationError struct {
	Field   string // 字段名
	Value   string // 错误值
	Message string // 错误消息
}

// Error 实现error接口
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s' with value '%s': %s", e.Field, e.Value, e.Message)
}

// NewValidationError 创建新的验证错误
func NewValidationError(field, value, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
	}
}
