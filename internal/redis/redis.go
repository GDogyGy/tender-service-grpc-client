package redis

import (
	"context"
	"time"

	"GrpcClientForTenderService/internal/config"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *redis.Client
}

func NewRedisClient(cfg *config.Config) *Client {
	return &Client{redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.RedisPassword,
		DB:       0,
	})}
}

func (r *Client) Get(ctx context.Context, key string) ([]byte, error) {
	bytes, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (r *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) ([]byte, error) {
	bytes, err := r.client.Set(ctx, key, value, expiration).Bytes()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (r *Client) Close() error {
	return r.client.Close()
}
