package fetch

import (
	"context"
	"fmt"

	"GrpcClientForTenderService/internal/domain"
	"GrpcClientForTenderService/internal/gateway/types/transport"
	"GrpcClientForTenderService/internal/protos/gen/tender/fetch"
)

type log interface {
	Error(msg string, args ...any)
}

type Gateway struct {
	log log
	t   fetch.TenderServiceFetchClient
}

func NewGateway(l log, t fetch.TenderServiceFetchClient) *Gateway {
	return &Gateway{
		l, t,
	}
}

func (h *Gateway) FetchList(ctx context.Context, serviceType string) ([]transport.Tender, error) {
	const op = "gateway.tender.fetch.FetchList"
	var tenders []transport.Tender

	resp, err := h.t.FetchList(ctx, &fetch.RequestFetchListV1{ServiceType: serviceType})
	if err != nil {
		h.log.Error(fmt.Errorf("%s: %w", op, err).Error())
		return tenders, err
	}

	tenders = make([]transport.Tender, 0, len(resp.Tenders))
	for _, t := range resp.GetTenders() {
		tenders = append(tenders, transport.Tender{
			Id:          t.GetId(),
			Name:        t.GetName(),
			Description: t.GetDescription(),
			ServiceType: t.GetServiceType(),
			Status:      t.GetStatus(),
			Version:     int(t.GetVersion()),
			Responsible: t.GetResponsible(),
		})
	}

	return tenders, nil
}

func (h *Gateway) FetchListByUser(ctx context.Context, username string) ([]transport.Tender, error) {
	const op = "gateway.tender.fetch.FetchListByUser"
	var tenders []transport.Tender

	if username == "" {
		h.log.Error(fmt.Errorf("%s: %w", op, domain.ArgumentsInvalid(op, fmt.Errorf(""))).Error())
		return nil, domain.ArgumentsInvalid(op, fmt.Errorf(""))
	}

	resp, err := h.t.FetchListByUser(ctx, &fetch.RequestFetchListByUserV1{Username: username})
	if err != nil {
		h.log.Error(fmt.Errorf("%s: %w", op, err).Error())
		return nil, err
	}

	tenders = make([]transport.Tender, 0, len(resp.Tenders))
	for _, t := range resp.GetTenders() {
		tenders = append(tenders, transport.Tender{
			Id:          t.GetId(),
			Name:        t.GetName(),
			Description: t.GetDescription(),
			ServiceType: t.GetServiceType(),
			Status:      t.GetStatus(),
			Version:     int(t.GetVersion()),
			Responsible: t.GetResponsible(),
		})
	}

	return tenders, nil
}

func (h *Gateway) FetchStatus(ctx context.Context, username string, tenderId string) (string, error) {
	const op = "gateway.tender.fetch.FetchStatus"
	var result string

	if username == "" || tenderId == "" {
		err := domain.ErrInvalidArgument(op, "username and tenderId are required")
		h.log.Error(fmt.Errorf("%s: %w", op, err).Error())
		return result, err
	}

	resp, err := h.t.FetchStatus(ctx, &fetch.RequestFetchStatusV1{Username: username, TenderId: tenderId})
	if err != nil {
		err = domain.ErrInternal(op, "failed to fetch tender status", err)
		h.log.Error("%s: %v", op, err)
		return result, err
	}

	if len(resp.GetStatus()) == 0 {
		err = domain.ErrNotFound(op, fmt.Errorf("tender not found"))
		h.log.Error("%s: %v", op, err)
		return result, err
	}

	return resp.GetStatus(), nil
}
