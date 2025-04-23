package kafka

import (
	"context"
	"fmt"
	"sync"
)

type Manager struct {
	log      Log
	consumer *Consumer
	router   *RouterKafka
	wg       sync.WaitGroup
	msgChan  chan []byte
}

func NewManager(
	log Log,
	consumer *Consumer,
	router *RouterKafka,
	msgChan chan []byte,
) *Manager {
	return &Manager{
		log:      log,
		consumer: consumer,
		router:   router,
		msgChan:  msgChan,
	}
}

func (m *Manager) Run(ctx context.Context, topics []string) {
	m.wg.Add(2)

	go m.processMessages(ctx)

	go m.runConsumer(ctx, topics)
}

func (m *Manager) processMessages(ctx context.Context) {
	const op = "kafka.manager.processMessages"
	defer m.wg.Done()

	for {
		select {
		case msg := <-m.msgChan:
			if err := m.router.RouteMessage(ctx, msg); err != nil {
				m.log.Error(fmt.Errorf("%s: %w", op, err).Error())
			}
		case <-ctx.Done():
			m.log.Info("stopping message processor")
			return
		}
	}
}

func (m *Manager) runConsumer(ctx context.Context, topics []string) {
	defer m.wg.Done()
	m.consumer.Start(ctx, topics)
}

func (m *Manager) Close() {
	close(m.msgChan)
	m.wg.Wait()
	func() { _ = m.consumer.Close() }()
}
