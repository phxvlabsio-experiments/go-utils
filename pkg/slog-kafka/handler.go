package slogkafka

/*
import (
	"context"
	"encoding/json"
	"github.com/twmb/franz-go/pkg/kgo"

	// "github.com/redpanda-data/franz-go/pkg/kgo"
	"golang.org/x/exp/slog"
)

type Option struct {
	Level     slog.Leveler `json:"level,omitempty"`
	Client    *kgo.Client  `json:"client,omitempty"`
	Converter Converter    `json:"converter,omitempty"`
}

func (o Option) NewKafkaHandler() slog.Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	if o.Client == nil {
		panic("missing Kafka client")
	}

	return &KafkaHandler{
		option: o,
		attrs:  []slog.Attr{},
		groups: []string{},
	}
}

type KafkaHandler struct {
	option Option
	attrs  []slog.Attr
	groups []string
}

func (k KafkaHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= k.option.Level.Level()
}

func (k KafkaHandler) Handle(ctx context.Context, record slog.Record) error {
	if k.Enabled(ctx, record.Level) {
		return nil
	}

	payload := k.option.Converter(k.attrs, &record)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	msg := kgo.Record{
		Value: payloadBytes,
	}

	ctx := context.Background()

	return nil
}

func (k KafkaHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	//TODO implement me
	panic("implement me")
}

func (k KafkaHandler) WithGroup(name string) slog.Handler {
	//TODO implement me
	panic("implement me")
}
*/
