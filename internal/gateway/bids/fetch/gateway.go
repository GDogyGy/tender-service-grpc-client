package fetch

import (
	"context"
	"fmt"

	"GrpcClientForTenderService/internal/domain"
	"GrpcClientForTenderService/internal/gateway/types/transport"
	"GrpcClientForTenderService/internal/protos/gen/bids/fetch"
)

type log interface {
	Error(msg string, args ...any)
}

type Gateway struct {
	log log
	b   fetch.BidsServiceFetchClient
}

func NewGateway(l log, b fetch.BidsServiceFetchClient) *Gateway {
	return &Gateway{
		l, b,
	}
}

func (h *Gateway) FetchListByTender(ctx context.Context, username string, tenderId string) ([]transport.Bids, error) {
	const op = "gateway.bids.fetch.FetchListByTender"
	var bids []transport.Bids

	resp, err := h.b.FetchListByTender(ctx, &fetch.BidsRequestFetchListV1{Username: username, TenderID: tenderId})
	if err != nil {
		h.log.Error(fmt.Errorf("%s: %w", op, err).Error())
		return bids, err
	}

	bids = make([]transport.Bids, 0, len(resp.Bids))
	for _, t := range resp.GetBids() {
		bids = append(bids, transport.Bids{
			Id:          t.GetId(),
			Name:        t.GetName(),
			Description: t.GetDescription(),
			Status:      t.GetStatus(),
			TenderId:    t.GetTenderId(),
			Version:     int(t.GetVersion()),
			Responsible: t.GetResponsible(),
		})
	}

	return bids, nil
}

func (h *Gateway) FetchListByUser(ctx context.Context, username string) ([]transport.Bids, error) {
	const op = "gateway.bids.fetch.FetchListByUser"
	var bids []transport.Bids

	if username == "" {
		h.log.Error(fmt.Errorf("%s: %w", op, domain.ArgumentsInvalid(op, fmt.Errorf(""))).Error())
		return nil, domain.ArgumentsInvalid(op, fmt.Errorf(""))
	}

	resp, err := h.b.FetchListByUser(ctx, &fetch.BidsRequestFetchListByUserV1{Username: username})
	if err != nil {
		h.log.Error(fmt.Errorf("%s: %w", op, err).Error())
		return nil, err
	}

	bids = make([]transport.Bids, 0, len(resp.Bids))
	for _, t := range resp.GetBids() {
		bids = append(bids, transport.Bids{
			Id:          t.GetId(),
			Name:        t.GetName(),
			Description: t.GetDescription(),
			Status:      t.GetStatus(),
			TenderId:    t.GetTenderId(),
			Version:     int(t.GetVersion()),
			Responsible: t.GetResponsible(),
		})
	}

	return bids, nil
}

func (h *Gateway) FetchStatus(ctx context.Context, username string, bidsID string) (string, error) {
	const op = "gateway.bids.fetch.FetchStatus"
	var result string

	if username == "" || bidsID == "" {
		err := domain.ErrInvalidArgument(op, "username and bidsId are required")
		h.log.Error(fmt.Errorf("%s: %w", op, err).Error())
		return result, err
	}

	resp, err := h.b.FetchStatus(ctx, &fetch.BidsRequestFetchStatusV1{Username: username, BidID: bidsID})
	if err != nil {
		err = domain.ErrInternal(op, "failed to fetch bids status", err)
		h.log.Error("%s: %v", op, err)
		return result, err
	}

	if len(resp.GetStatus()) == 0 {
		err = domain.ErrNotFound(op, fmt.Errorf("bids not found"))
		h.log.Error("%s: %v", op, err)
		return result, err
	}

	return resp.GetStatus(), nil
}

func (h *Gateway) FetchReviews(ctx context.Context, username string, tenderID string, authorUsername string, organizationID string) ([]transport.BidFeedback, error) {
	const op = "gateway.bids.fetch.FetchStatus"
	var BidFeedback []transport.BidFeedback

	if username == "" || tenderID == "" || authorUsername == "" || organizationID == "" {
		err := domain.ErrInvalidArgument(op, "username and bidsId are required")
		h.log.Error(fmt.Errorf("%s: %w", op, err).Error())
		return BidFeedback, err
	}

	resp, err := h.b.FetchReviews(ctx, &fetch.BidsRequestFetchReviewsV1{Username: username, TenderId: tenderID, OrganizationId: organizationID, AuthorUsername: authorUsername})
	if err != nil {
		err = domain.ErrInternal(op, "failed to fetch bids status", err)
		h.log.Error("%s: %v", op, err)
		return BidFeedback, err
	}

	BidFeedback = make([]transport.BidFeedback, 0, len(resp.Feedback))
	for _, t := range resp.GetFeedback() {
		BidFeedback = append(BidFeedback, transport.BidFeedback{
			Id:          t.GetId(),
			BidID:       t.GetBidId(),
			Description: t.GetDescription(),
			Responsible: t.GetResponsible(),
			CreatedAt:   t.GetCreatedAt(),
		})
	}

	return BidFeedback, nil
}
