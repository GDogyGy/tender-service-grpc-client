package fetch

import (
	"context"

	"GrpcClientForTenderService/internal/gateway/types/transport"
)

type Service struct {
	tender gateway
}

type gateway interface {
	FetchList(ctx context.Context, serviceType string) ([]transport.Tender, error)
	FetchListByUser(ctx context.Context, username string) ([]transport.Tender, error)
	FetchStatus(ctx context.Context, username string, tenderId string) (string, error)
}

func NewService(g gateway) *Service {
	return &Service{tender: g}
}

func (s *Service) FetchList(ctx context.Context, serviceType string) ([]transport.Tender, error) {
	return s.tender.FetchList(ctx, serviceType)
}

func (s *Service) FetchListByUser(ctx context.Context, username string) ([]transport.Tender, error) {
	return s.tender.FetchListByUser(ctx, username)
}

func (s *Service) FetchStatus(ctx context.Context, username string, tenderId string) (string, error) {
	return s.tender.FetchStatus(ctx, username, tenderId)
}
