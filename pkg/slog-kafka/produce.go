package slogkafka

import (
	"context"
	"github.com/twmb/franz-go/pkg/kgo"

	// "github.com/redpanda-data/franz-go/pkg/kgo"
	"golang.org/x/exp/slog"
)

type KafkaProducer struct {
	client *kgo.Client
	topic  string
}

func (k *KafkaProducer) Produce(ctx context.Context, r slog.Record) {
	//
	return
}
