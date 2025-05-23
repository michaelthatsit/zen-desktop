//go:build !prod

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

// More on the logging setup in README.md.

func SetupLogger() error {
	logsDir, err := getLogsDir(appName)
	if err != nil {
		return fmt.Errorf("get logs directory: %w", err)
	}

	fileLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logsDir, "application.log"),
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     1,
		Compress:   true,
	}

	log.SetOutput(io.MultiWriter(os.Stdout, fileLogger))

	return nil
}
