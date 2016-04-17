package yext

import (
	"log"
	"os"
)

type Logger interface {
	Log(...interface{})
}

type stdLogger struct {
	logger *log.Logger
}

func NewStdLogger() Logger {
	return &stdLogger{
		logger: log.New(os.Stdout, "yext-go", log.LstdFlags),
	}
}

func (s *stdLogger) Log(args ...interface{}) {
	s.logger.Println(args...)
}
