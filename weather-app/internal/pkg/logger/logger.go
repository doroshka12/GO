package logger

import (
    "fmt"
    "io"
    "log"
    "os"
)

// Level уровень логирования
type Level int

const (
    DEBUG Level = iota
    INFO
    ERROR
)

// Logger интерфейс логгера
type Logger interface {
    Debug(format string, v ...interface{})
    Info(format string, v ...interface{})
    Error(format string, v ...interface{})
    SetLevel(level Level)
}

// SimpleLogger реализация логгера
type SimpleLogger struct {
    level  Level
    debug  *log.Logger
    info   *log.Logger
    errLog *log.Logger
}

// NewLogger создает новый логгер
func NewLogger(level Level, output io.Writer) *SimpleLogger {
    return &SimpleLogger{
        level:  level,
        debug:  log.New(output, "DEBUG: ", log.Ldate|log.Ltime),
        info:   log.New(output, "INFO: ", log.Ldate|log.Ltime),
        errLog: log.New(output, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
    }
}

// NewDefaultLogger создает логгер с выводом в stdout
func NewDefaultLogger(level Level) *SimpleLogger {
    return NewLogger(level, os.Stdout)
}

func (l *SimpleLogger) Debug(format string, v ...interface{}) {
    if l.level <= DEBUG {
        l.debug.Output(2, fmt.Sprintf(format, v...))
    }
}

func (l *SimpleLogger) Info(format string, v ...interface{}) {
    if l.level <= INFO {
        l.info.Output(2, fmt.Sprintf(format, v...))
    }
}

func (l *SimpleLogger) Error(format string, v ...interface{}) {
    if l.level <= ERROR {
        l.errLog.Output(2, fmt.Sprintf(format, v...))
    }
}

func (l *SimpleLogger) SetLevel(level Level) {
    l.level = level
}