package shared

import (
	"fmt"
	"runtime"
)

type ErrorSeverity string

const (
	SeverityInfo    ErrorSeverity = "info"
	SeverityWarning ErrorSeverity = "warning"
	SeverityError   ErrorSeverity = "error"
	SeverityFatal   ErrorSeverity = "fatal"
)

type AppError struct {
	Code     int           `json:"code"`
	Message  string        `json:"message"`
	Severity ErrorSeverity `json:"severity"`
	Stack    string        `json:"stack,omitempty"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] Error %d: %s", e.Severity, e.Code, e.Message)
}

func New(code int, msg string, severity ErrorSeverity) *AppError {
	return &AppError{
		Code:     code,
		Message:  msg,
		Severity: severity,
		Stack:    getErrorStackTrace(severity),
	}
}

func getErrorStackTrace(severity ErrorSeverity) string {
	var depth = 3
	switch severity {
	case SeverityInfo:
		depth = 3
		break
	case SeverityError:
		depth = 9
		break
	case SeverityFatal:
		depth = 7
	case SeverityWarning:
		depth = 5
		break
	default:
		depth = 3
		break
	}
	return captureErrorStackTrace(depth)
}

func captureErrorStackTrace(depth int) string {
	pc, file, line, ok := runtime.Caller(depth)
	if !ok {
		return ""
	}
	fn := runtime.FuncForPC(pc)
	return fmt.Sprintf("%s:%d (%s)", file, line, fn.Name())
}
