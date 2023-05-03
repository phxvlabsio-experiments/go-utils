package main

import (
	"context"
	"github.com/Shopify/sarama"
	slogsarama "github.com/phxvlabs.dev/meshkit-utils/pkg/slog-sarama"
	"golang.org/x/exp/slog"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var topic string
var broker string

const defaultBroker = "redpanda-0.redpanda.redpanda.svc.cluster.local.:9093"
const defaultTopic = "demo"
const consumerGroup = "test-demo"

func init() {
	broker = os.Getenv("REDPANDA_BROKER")
	if broker == "" {
		broker = defaultBroker
		log.Println("using default value for redpanda broker", defaultBroker)
	}

	topic = os.Getenv("TOPIC")
	if topic == "" {
		topic = defaultTopic
		log.Println("using default value for redpanda topic", defaultTopic)
	}
}

func main() {
	keepRunning := true
	log.Println("starting consumer")

	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// logger := slogsarama.Option{Level: slog.LevelDebug}

	sarama.Logger = slog.New(slogsarama.Option{Level: slog.LevelDebug})

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup([]string{broker}, consumerGroup, config)
	if err != nil {
		log.Panicf("error creating consumer group client: %v", err)
	}

	consumer := SimpleConsumerHandler{
		ready: make(chan bool),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			err := client.Consume(ctx, []string{topic}, &consumer)
			if err != nil {
				log.Panicf("error joining consumer group: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	log.Println("consumer ready")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		}
	}
	cancel()
	wg.Wait()

	err = client.Close()
	if err != nil {
		log.Panicf("error closing kafka client: %v", err)
	}
}

type SimpleConsumerHandler struct {
	ready chan bool
}

func (s SimpleConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	close(s.ready)
	return nil
}

func (s SimpleConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (s SimpleConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			log.Printf("message received: value = %s, topic = %v, partition = %v, topic = %v", string(message.Value), message.Topic, message.Partition, message.Offset)

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
