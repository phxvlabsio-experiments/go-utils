package slogsarama

import (
	"fmt"
	"golang.org/x/exp/slog"
)

type SaramaLogger struct {
	level slog.Level
}

func (sl SaramaLogger) Print(v ...interface{}) {
	if sl.level >= slog.LevelInfo {
		slog.Info(fmt.Sprint(v...))
	}
}

func (sl *SaramaLogger) Printf(format string, v ...interface{}) {
	if sl.level >= slog.LevelInfo {
		slog.Info(fmt.Sprintf(format, v...))
	}
}

func (sl *SaramaLogger) Println(v ...interface{}) {
	if sl.level >= slog.LevelInfo {
		slog.Info(fmt.Sprintln(v...))
	}
}
