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
	if _, err := os.Stat(cfg.BASE_LOG_DIR); os.IsNotExist(err) {
		return os.MkdirAll(cfg.BASE_LOG_DIR, 0755)
	}
	return nil
}

func GetNextLogFileName(cfg *LoggerConfig) (string, error) {
	date_str := time.Now().Format(cfg.FILE_PATTERN)

	base_name := fmt.Sprintf("%s_%s.log", "log", date_str)

	if err := EnsureBaseLogDirectoryExist(cfg); err != nil {
		return "", err
	}
	var index int
	var log_file string

	for {
		log_file = fmt.Sprintf("%s_%d.log", base_name, index+1)
		log_path := filepath.Join(cfg.BASE_LOG_DIR, log_file)
		if _, err := os.Stat(log_file); os.IsNotExist(err) {
			return log_file, nil
		}

		file_info, err := os.Stat(log_file)
		if err != nil {
			return "", err
		}

		if file_info.Size() < int64(cfg.MAX_SIZE*1024*1024) {
			return log_path, nil
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
	log_filename, err := GetNextLogFileName(cfg)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(log_filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	writer := zapcore.AddSync(file)

	log_level := ParseLogLevel(cfg.LOG_LEVEL)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		writer,
		log_level,
	)

	return zap.New(core, zap.AddCaller()), nil
}

var LogInstance *zap.Logger

func InitializeLogger() {
	log_cfg, err := LoadLoggerConfig()
	if err != nil {
		//
	}
	LogInstance, err = InitializeNewLogInstance(log_cfg)
	if err != nil {
		//
	}
}

func CleanupQueuedLogs() {
	if LogInstance != nil {
		_ = LogInstance.Sync()
	}
}
