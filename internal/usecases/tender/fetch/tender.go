package fetch

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"GrpcClientForTenderService/internal/gateway/types/transport"
)

type log interface {
	Error(msg string, args ...any)
}

type Service struct {
	log    log
	tender gateway
	cache  cacheClient
}

type cacheClient interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) ([]byte, error)
}

type gateway interface {
	FetchList(ctx context.Context, serviceType string) ([]transport.Tender, error)
	FetchListByUser(ctx context.Context, username string) ([]transport.Tender, error)
	FetchStatus(ctx context.Context, username string, tenderId string) (string, error)
}

func NewService(l log, g gateway, c cacheClient) *Service {
	return &Service{log: l, tender: g, cache: c}
}

func (s *Service) FetchList(ctx context.Context, serviceType string) ([]transport.Tender, error) {
	const op = "useCases.tender.fetch.FetchList"
	var cacheKey = "tenderList:" + serviceType

	cachedData, err := s.cache.Get(ctx, cacheKey)
	if err != nil {
		s.log.Error(fmt.Errorf("%s: cache get error for key %s: %v", op, cacheKey, err).Error())
	}

	if err == nil {
		var tenders []transport.Tender
		if err := json.Unmarshal(cachedData, &tenders); err == nil {
			return tenders, nil
		}

		s.log.Error(fmt.Errorf("%s: cache get error for key %s: %v", op, cacheKey, err).Error())
	}

	tenders, err := s.tender.FetchList(ctx, serviceType)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to fetch tenders: %w", op, err)
	}

	go func() {
		tenderData, err := json.Marshal(tenders)
		if err != nil {
			s.log.Error(fmt.Errorf("%s: failed to marshal tenders: %v", op, err).Error())
			return
		}

		if _, err := s.cache.Set(ctx, cacheKey, tenderData, time.Hour); err != nil {
			s.log.Error(fmt.Errorf("%s: failed to update cache: %v", op, err).Error())
		}
	}()

	return tenders, nil
}

func (s *Service) FetchListByUser(ctx context.Context, username string) ([]transport.Tender, error) {
	return s.tender.FetchListByUser(ctx, username)
}

func (s *Service) FetchStatus(ctx context.Context, username string, tenderId string) (string, error) {
	return s.tender.FetchStatus(ctx, username, tenderId)
}
