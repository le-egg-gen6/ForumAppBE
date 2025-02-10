package logger

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

var LogInstance *zap.Logger

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

func InitializeNewLogInstance(cfg *LoggerConfig) (*zap.Logger, error) {

}
