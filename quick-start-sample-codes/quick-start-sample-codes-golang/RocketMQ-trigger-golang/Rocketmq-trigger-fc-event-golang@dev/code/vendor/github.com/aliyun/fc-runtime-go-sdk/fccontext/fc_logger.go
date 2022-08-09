package fccontext

import (
	"fmt"
	"log"
	"time"
)

type logWriter struct {
	requestID string
	level string
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " " + writer.requestID + " ["+ writer.level + "] " + string(bytes))
}

func newLogWriter(requestID, level string) *logWriter {
	return &logWriter{
		requestID: requestID,
		level: level,
	}
}

// FcLogger ...
type FcLogger struct {
	debugLogger *log.Logger
	infoLogger *log.Logger
	warnLogger *log.Logger
	errorLogger *log.Logger
}

// NewFcLogger ...
func NewFcLogger(requestID string) *FcLogger {
	return &FcLogger{
		debugLogger:    log.New(newLogWriter(requestID, "DEBUG"), "", log.Lshortfile),
		infoLogger:    log.New(newLogWriter(requestID, "INFO"), "", log.Lshortfile),
		warnLogger:    log.New(newLogWriter(requestID, "WARN"), "", log.Lshortfile),
		errorLogger:    log.New(newLogWriter(requestID, "ERROR"), "", log.Lshortfile),
	}
}

// Debug ...
func (l *FcLogger) Debug(v ...interface{}) {
	l.debugLogger.Output(2, fmt.Sprintln(v...))
}

// Debugf ...
func (l *FcLogger) Debugf(format string, v ...interface{}) {
	l.debugLogger.Output(2, fmt.Sprintf(format, v...))
}

// Info ...
func (l *FcLogger) Info(v ...interface{}) {
	l.infoLogger.Output(2, fmt.Sprintln(v...))
}

// Infof ...
func (l *FcLogger) Infof(format string, v ...interface{}) {
	l.infoLogger.Output(2, fmt.Sprintf(format, v...))
}

// Warn ...
func (l *FcLogger) Warn(v ...interface{}) {
	l.warnLogger.Output(2, fmt.Sprintln(v...))
}

// Warnf ...
func (l *FcLogger) Warnf(format string, v ...interface{}) {
	l.warnLogger.Output(2, fmt.Sprintf(format, v...))
}

// Error ...
func (l *FcLogger) Error(v ...interface{}) {
	l.errorLogger.Output(2, fmt.Sprintln(v...))
}

// Errorf ...
func (l *FcLogger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Output(2, fmt.Sprintf(format, v...))
}

