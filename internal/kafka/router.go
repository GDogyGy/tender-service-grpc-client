package kafka

import (
	"context"
	"encoding/json"
	"fmt"
)

type EventHandler interface {
	HandleEvent(ctx context.Context, payload []byte) error
}

type RouterKafka struct {
	handlers map[string]EventHandler
	logger   Log
}

func NewRouterKafka(logger Log) *RouterKafka {
	return &RouterKafka{
		handlers: make(map[string]EventHandler),
		logger:   logger,
	}
}

func (r *RouterKafka) RegisterHandler(eventType string, handler EventHandler) {
	r.handlers[eventType] = handler
}

func (r *RouterKafka) RouteMessage(ctx context.Context, message []byte) error {
	var baseEvent struct {
		EventType string `json:"event_type"`
	}

	if err := json.Unmarshal(message, &baseEvent); err != nil {
		return fmt.Errorf("failed to unmarshal base event: %w", err)
	}

	handler, exists := r.handlers[baseEvent.EventType]
	if !exists {
		return fmt.Errorf("no handler found for event type: %s", baseEvent.EventType)
	}

	return handler.HandleEvent(ctx, message)
}
