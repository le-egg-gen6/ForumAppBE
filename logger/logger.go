package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func EnsureBaseLogDirectoryExist(cfg *LoggerConfig) error {
	if _, err := os.Stat(cfg.BaseLogDir); os.IsNotExist(err) {
		return os.MkdirAll(cfg.BaseLogDir, 0755)
	}
	return nil
}

func GetNextLogFileName(cfg *LoggerConfig) (string, error) {
	dateStr := time.Now().Format(cfg.FilePattern)

	baseName := fmt.Sprintf("%s_%s", "log", dateStr)

	if err := EnsureBaseLogDirectoryExist(cfg); err != nil {
		return "", err
	}
	var index int
	var logFile string

	for {
		logFile = fmt.Sprintf("%s_%d.log", baseName, index+1)
		logPath := filepath.Join(cfg.BaseLogDir, logFile)
		if _, err := os.Stat(logPath); os.IsNotExist(err) {
			return logPath, nil
		}

		fileInfo, err := os.Stat(logPath)
		if err != nil {
			return "", err
		}

		if fileInfo.Size() < int64(cfg.MaxSize*1024*1024) {
			return logPath, nil
		}
		index++
	}

}

func ParseLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}

func InitializeNewLogInstance(cfg *LoggerConfig) (*zap.Logger, error) {
	logFilename, err := GetNextLogFileName(cfg)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	writer := zapcore.AddSync(file)

	logLevel := ParseLogLevel(cfg.LogLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		writer,
		logLevel,
	)

	return zap.New(core, zap.AddCaller()), nil
}

var Instance *zap.Logger

func InitializeLogger() {
	logCfg, err := LoadLoggerConfig()
	if err != nil {
		panic("Logger configuration file not found")
	}
	Instance, err = InitializeNewLogInstance(logCfg)
	if err != nil {
		panic("Logger not initialized")
	}
}

func GetLogInstance() *zap.Logger {
	return Instance
}

func CleanupQueuedLogs() {
	if Instance != nil {
		_ = Instance.Sync()
	}
}
