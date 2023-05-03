package slogpipeline

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/exp/slog"
)

type ChannelLogger struct {
	channel string
}

func (l *ChannelLogger) Log(level slog.Level, msg string, keyValues ...interface{}) error {
	timestamp := time.Now().Format(time.RFC3339)
	message := fmt.Sprintf("[%s] %s", l.channel, msg)

	var keyvalPairs []string
	for i := 0; i < len(keyValues); i += 2 {
		key := fmt.Sprintf("%v", keyValues[i])
		val := fmt.Sprintf("%v", keyValues[i+1])
		keyvalPairs = append(keyvalPairs, fmt.Sprintf("%s=%s", key, val))
	}

	keyvalString := ""
	if len(keyvalPairs) > 0 {
		keyvalString = " " + strings.Join(keyvalPairs, ", ")
	}

	fmt.Printf("%s [%s] %s%s\n", timestamp, level.String(), message, keyvalString)

	return nil
}

func NewPipelineChannel(channel string) *ChannelLogger {
	return &ChannelLogger{
		channel: channel,
	}
}
