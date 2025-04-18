package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

type Consumer struct {
	client      sarama.ConsumerGroup
	ready       chan bool
	messageChan chan<- string
}

func NewConsumer(brokers []string, groupID string, messageChan chan<- string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRange(),
	}

	client, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, fmt.Errorf("error creating consumer group client: %w", err)
	}

	return &Consumer{
		client:      client,
		ready:       make(chan bool),
		messageChan: messageChan,
	}, nil
}

func (c *Consumer) Start(ctx context.Context, topics []string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := c.client.Consume(ctx, topics, c); err != nil {
				log.Printf("Error from consumer: %v", err)
				if ctx.Err() != nil {
					return
				}
			}
			if ctx.Err() != nil {
				return
			}
			c.ready = make(chan bool)
		}
	}()

	<-c.ready
	log.Println("Sarama consumer up and running...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Println("Terminating: context cancelled")
	case <-sigterm:
		log.Println("Terminating: via signal")
	}

	wg.Wait()
	if err := c.client.Close(); err != nil {
		log.Printf("Error closing client: %v", err)
	}
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s",
			string(message.Value), message.Timestamp, message.Topic)

		c.messageChan <- string(message.Value)

		session.MarkMessage(message, "")
	}
	return nil
}
