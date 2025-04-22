package fetch

import (
	"context"
	"encoding/json"
	"net/http"

	"GrpcClientForTenderService/internal/gateway/types/transport"
)

type log interface {
	Error(msg string, args ...any)
}

type useCasesTenderFetch interface {
	FetchList(ctx context.Context, serviceType string) ([]transport.Tender, error)
	FetchListByUser(ctx context.Context, username string) ([]transport.Tender, error)
	FetchStatus(ctx context.Context, username string, tenderId string) (string, error)
}

type Handler struct {
	log    log
	tender useCasesTenderFetch
}

func NewHandler(l log, t useCasesTenderFetch) Handler {
	return Handler{
		l, t,
	}
}

func (h *Handler) Register(router *http.ServeMux) {
	router.HandleFunc(http.MethodGet+" /api/tenders", h.FetchList)
	router.HandleFunc(http.MethodGet+" /api/tenders/my", h.FetchListByUser)
	router.HandleFunc(http.MethodGet+" /api/tenders/status", h.FetchStatus)
}

func (h *Handler) FetchList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rq := r.URL.Query()

	param := "servicetype"

	if len(rq) >= 1 && rq.Get(param) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tendersTransport, err := h.tender.FetchList(r.Context(), rq.Get(param))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(tendersTransport)
	if err != nil {
		h.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(tendersTransport) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		h.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) FetchListByUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rq := r.URL.Query()
	param := "username"
	if len(rq) >= 1 && rq.Get(param) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tenders, err := h.tender.FetchListByUser(r.Context(), rq.Get(param))
	if err != nil {
		h.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(tenders)
	if err != nil {
		h.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(tenders) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		h.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) FetchStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rq := r.URL.Query()
	user := "username"
	tenderId := "tenderId"

	if len(rq) > 2 || rq.Get(user) == "" || rq.Get(tenderId) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	status, err := h.tender.FetchStatus(r.Context(), rq.Get(user), rq.Get(tenderId))
	if err != nil {
		h.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(status)
	if err != nil {
		h.log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		h.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
