package main

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	topic := "demo"
	seeds := []string{"localhost:9093"}

	client, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	admin := kadm.NewClient(client)
	defer admin.Close()

	ctx := context.Background()
	resp, _ := admin.CreateTopics(ctx, 1, 1, nil, topic)
	for _, ctr := range resp {
		if ctx.Err != nil {
			fmt.Printf("Unable to create topic '%s': %s", ctr.Topic, ctr.Err)
		} else {
			fmt.Printf("Created topic '%s'", ctr.Topic)
		}
	}
}
