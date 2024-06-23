package logger

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	loggerJson slog.Logger
	loggerText slog.Logger
}

func NewLogger() *Logger {
	o := os.Stdout
	loggerJson := slog.New(slog.NewJSONHandler(o, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	e := os.Stderr
	loggerText := slog.New(
		tint.NewHandler(e, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
			NoColor:    false,
		}),
	)
	slog.SetDefault(loggerJson)
	slog.SetDefault(loggerText)

	return &Logger{
		loggerJson: *loggerJson,
		loggerText: *loggerText,
	}
}

func (l *Logger) DebugJson(msg string, info ...any) {
	l.loggerJson.Debug(msg, info...)
}
func (l *Logger) InfoJson(msg string, info ...any) {
	l.loggerJson.Info(msg, info...)
}
func (l *Logger) WarnJson(msg string, info ...any) {
	l.loggerJson.Warn(msg, info...)
}
func (l *Logger) ErrorJson(msg string, info ...any) {
	l.loggerJson.Error(msg, info...)
}

func (l *Logger) DebugText(msg string, info ...any) {
	l.loggerText.Debug(msg, info...)
}

func (l *Logger) InfoText(msg string, info ...any) {
	l.loggerText.Info(msg, info...)
}
func (l *Logger) WarnText(msg string, info ...any) {
	l.loggerText.Warn(msg, info...)
}
func (l *Logger) ErrorText(msg string, info ...any) {
	l.loggerText.Error(msg, info...)
}
