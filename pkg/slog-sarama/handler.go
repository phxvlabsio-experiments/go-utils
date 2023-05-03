package slogsarama

import (
	"context"
	"github.com/Shopify/sarama"
	"golang.org/x/exp/slog"
)

type Option struct {
	Level  slog.Leveler
	Client sarama.Client
	Topic  string
}

func (o Option) NewSaramaHandler() slog.Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	if o.Client == nil {
		panic("missing kafka connection")
	}

	return &SaramaHandler{
		option: o,
		attrs:  []slog.Attr{},
		groups: []string{},
	}
}

type SaramaHandler struct {
	option Option
	attrs  []slog.Attr
	groups []string
}

func (s SaramaHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= s.option.Level.Level()
}

func (s SaramaHandler) Handle(ctx context.Context, record slog.Record) error {
	if !s.Enabled(ctx, record.Level) {
		return nil
	}

	message := &sarama.ProducerMessage{
		Topic: s.option.Topic,
		Value: sarama.StringEncoder(record.Message),
	}

	messages := []*sarama.ProducerMessage{message}

	producer, err := sarama.NewSyncProducerFromClient(s.option.Client)
	if err != nil {
		return err
	}
	defer producer.Close()

	err = producer.SendMessages(messages)
	return err
}

func (s SaramaHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := make([]slog.Attr, len(s.option.attrs)+len(attrs))
	copy(newAttrs, s.option.attrs)
	copy(newAttrs[len(s.option.attrs):], attrs)

	return &SaramaHandler{
		option: Option{
			Level:  s.option.Level,
			Client: s.option.Client,
			attrs:  newAttrs,
			groups: s.option.groups,
		},
	}
}

func (s SaramaHandler) WithGroup(name string) slog.Handler {
	o := s.option
	o.groups = append(o.groups, name)
	return &SaramaHandler{option: o}
}
