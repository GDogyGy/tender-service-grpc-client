package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"sync"
)

type Log interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type Consumer struct {
	log     Log
	client  sarama.ConsumerGroup
	ready   chan bool
	msgChan chan<- []byte
}

func NewConsumer(log Log, brokers []string, groupID string, msgChan chan<- []byte) (*Consumer, error) {
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
		log:     log,
		client:  client,
		ready:   make(chan bool),
		msgChan: msgChan,
	}, nil
}

func (c *Consumer) Start(ctx context.Context, topics []string) {
	const op = "kafka.consumer.Start"
	wg := &sync.WaitGroup{}
	wg.Add(1)

	defer func() {
		if r := recover(); r != nil {
			c.log.Error(fmt.Sprintf("%s: panic recovered: %v", op, r))
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := c.client.Consume(ctx, topics, c); err != nil {
					c.log.Error(fmt.Sprintf("%s: %v", op, err))
				}
			}
		}
	}()
	<-c.ready
	c.log.Info("Kafka consumer started")

	<-ctx.Done()
	wg.Wait()
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

		c.msgChan <- message.Value

		session.MarkMessage(message, "")
	}
	return nil
}

func (c *Consumer) Close() error {
	close(c.msgChan)
	return c.client.Close()
}
