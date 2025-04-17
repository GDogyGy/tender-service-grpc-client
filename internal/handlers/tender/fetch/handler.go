package fetch

import (
	"net/http"

	"GrpcClientForTenderService/internal/protos/gen/tender/fetch"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type log interface {
	Error(msg string, args ...any)
}

type Handler struct {
	log log
	b   fetch.TenderServiceFetchClient // TODO: Как тут избавится от этой зависимости и как делают эти клиенты
}

func NewHandler(l log, b fetch.TenderServiceFetchClient) Handler {
	return Handler{
		l, b,
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

	b, _ := h.b.FetchList(r.Context(), &fetch.RequestFetchListV1{ServiceType: rq.Get(param)})
	marshalOptions := protojson.MarshalOptions{
		Multiline:       true,
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	jsonBytes, err := marshalOptions.Marshal(proto.Message(b))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonBytes)
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

	b, _ := h.b.FetchListByUser(r.Context(), &fetch.RequestFetchListByUserV1{Username: rq.Get(param)})
	marshalOptions := protojson.MarshalOptions{
		Multiline:       true,
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	jsonBytes, err := marshalOptions.Marshal(proto.Message(b))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonBytes)
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
	b, _ := h.b.FetchStatus(r.Context(), &fetch.RequestFetchStatusV1{Username: rq.Get(user), TenderId: rq.Get(tenderId)})
	marshalOptions := protojson.MarshalOptions{
		Multiline:       true,
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	jsonBytes, err := marshalOptions.Marshal(proto.Message(b))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonBytes)
	if err != nil {
		h.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
