package logger

import (
	"log"
	"time"
)

type logLevel int

const (
	levelInfo = iota
	levelError
)

type logMsg struct {
	Time    time.Time
	Level   logLevel
	Message string
}

type Logger struct {
	logger  *log.Logger
	msgChan chan logMsg
}

// FAQ: buffered channel for logging?
func NewLogger(bufferSize int) *Logger {
	logger := &Logger{
		logger:  log.Default(),
		msgChan: make(chan logMsg, bufferSize),
	}

	go logger.start()

	return logger
}

func (l *Logger) start() {
	for msg := range l.msgChan {
		// TODO: Extension: abstract logger to change logger provider
		switch msg.Level {
		case levelInfo:
			l.logger.Println("[INFO]", msg.Time.Format(time.RFC3339), msg.Message)
		case levelError:
			l.logger.Println("[ERROR]", msg.Time.Format(time.RFC3339), msg.Message)
		default:
			l.logger.Println("[UNKNOWN]", msg.Time.Format(time.RFC3339), msg.Message)
		}
	}
}

func (l *Logger) LogInfo(message string) {
	l.msgChan <- logMsg{
		Time:    time.Now(),
		Level:   levelInfo,
		Message: message,
	}
}

func (l *Logger) LogError(message string) {
	l.msgChan <- logMsg{
		Time:    time.Now(),
		Level:   levelError,
		Message: message,
	}
}

func (l *Logger) Close() {
	close(l.msgChan)
}