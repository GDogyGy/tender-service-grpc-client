package fetch

import (
	"context"

	"GrpcClientForTenderService/internal/gateway/types/transport"
)

type Service struct {
	bids gateway
}

type gateway interface {
	FetchListByTender(ctx context.Context, username string, tenderId string) ([]transport.Bids, error)
	FetchListByUser(ctx context.Context, username string) ([]transport.Bids, error)
	FetchStatus(ctx context.Context, username string, bidsID string) (string, error)
	FetchReviews(ctx context.Context, username string, tenderID string, authorUsername string, organizationID string) ([]transport.BidFeedback, error)
}

func NewService(g gateway) *Service {
	return &Service{bids: g}
}

func (s *Service) FetchListByTender(ctx context.Context, username string, tenderId string) ([]transport.Bids, error) {
	return s.bids.FetchListByTender(ctx, username, tenderId)
}

func (s *Service) FetchListByUser(ctx context.Context, username string) ([]transport.Bids, error) {
	return s.bids.FetchListByUser(ctx, username)
}

func (s *Service) FetchStatus(ctx context.Context, username string, bidsID string) (string, error) {
	return s.bids.FetchStatus(ctx, username, bidsID)
}

func (s *Service) FetchReviews(ctx context.Context, username string, tenderID string, authorUsername string, organizationID string) ([]transport.BidFeedback, error) {
	return s.bids.FetchReviews(ctx, username, tenderID, authorUsername, organizationID)
}
