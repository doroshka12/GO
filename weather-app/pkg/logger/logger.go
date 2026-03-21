package logger

import (
    "fmt"
    "time"
)

// Уровни логирования
var (
    INFO  = "INFO"
    DEBUG = "DEBUG"
    ERROR = "ERROR"
)

// Logger структура логгера
type Logger struct{}

// New создает новый экземпляр логгера
func New() *Logger {
    return &Logger{}
}

// Info выводит информационное сообщение
func (l *Logger) Info(msg string) {
    fmt.Println(l.msg(INFO, msg))
}

// Debug выводит отладочное сообщение
func (l *Logger) Debug(msg string) {
    fmt.Println(l.msg(DEBUG, msg))
}

// Error выводит сообщение об ошибке
func (l *Logger) Error(msg string, err error) {
    fmt.Println(l.msg(ERROR, msg+" err - "+err.Error()))
}

// msg форматирует сообщение с уровнем и временем
func (l *Logger) msg(level string, msg string) string {
    timeStr := time.Now().Format(time.RFC3339)
    return fmt.Sprintf(
        "[%s] %s, message - %s",
        level,
        timeStr,
        msg,
    )
}