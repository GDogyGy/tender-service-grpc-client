package fetch

import (
	"net/http"
	"regexp"

	"GrpcClientForTenderService/internal/protos/gen/bids/fetch"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type log interface {
	Error(msg string, args ...any)
}

type Handler struct {
	log log
	b   fetch.BidsServiceFetchClient // TODO: Как тут избавится от этой зависимости и как делают эти клиенты
}

func NewHandler(l log, b fetch.BidsServiceFetchClient) Handler {
	return Handler{
		l, b,
	}
}

var tenderIdReviewRegexp = regexp.MustCompile(`/api/bids/(.*)/reviews\?`)
var tenderIdRegexp = regexp.MustCompile(`/api/bids/(.*)/list`)
var bidIdStatusRegexp = regexp.MustCompile(`/api/bids/(.*)/status\?`)

func (h *Handler) Register(router *http.ServeMux) {
	router.HandleFunc(http.MethodGet+" /api/bids/{tenderID}/list", h.FetchListByTender)
	router.HandleFunc(http.MethodGet+" /api/bids/my", h.FetchListByUser)
	router.HandleFunc(http.MethodGet+" /api/bids/{bidID}/status", h.FetchStatus)
	router.HandleFunc(http.MethodGet+" /api/bids/{tenderID}/reviews", h.FetchReviews)
}

func (h *Handler) FetchListByTender(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rq := r.URL.Query()
	user := "username"

	if rq.Get(user) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := tenderIdRegexp.FindStringSubmatch(r.RequestURI)
	if id == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, _ := h.b.FetchListByTender(r.Context(), &fetch.BidsRequestFetchListV1{Username: rq.Get(user), TenderID: id[1]})

	marshalOptions := protojson.MarshalOptions{
		Multiline:       true, // для красивого форматирования
		Indent:          "  ", // отступ двумя пробелами
		UseProtoNames:   true, // использовать оригинальные названия из .proto
		EmitUnpopulated: true, // включать пустые поля
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

	user := "username"
	if len(rq) >= 1 && rq.Get(user) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, _ := h.b.FetchListByUser(r.Context(), &fetch.BidsRequestFetchListByUserV1{Username: rq.Get(user)})

	marshalOptions := protojson.MarshalOptions{
		Multiline:       true, // для красивого форматирования
		Indent:          "  ", // отступ двумя пробелами
		UseProtoNames:   true, // использовать оригинальные названия из .proto
		EmitUnpopulated: true, // включать пустые поля
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

	if len(rq) > 1 || rq.Get(user) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bidID := bidIdStatusRegexp.FindStringSubmatch(r.RequestURI)
	if bidID == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, _ := h.b.FetchStatus(r.Context(), &fetch.BidsRequestFetchStatusV1{Username: rq.Get(user), BidID: bidID[1]})

	marshalOptions := protojson.MarshalOptions{
		Multiline:       true, // для красивого форматирования
		Indent:          "  ", // отступ двумя пробелами
		UseProtoNames:   true, // использовать оригинальные названия из .proto
		EmitUnpopulated: true, // включать пустые поля
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

func (h *Handler) FetchReviews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rq := r.URL.Query()
	username := "username"
	organizationID := "organizationId"
	authorUser := "authorUsername"
	if len(rq) < 3 || rq.Get(username) == "" || rq.Get(organizationID) == "" || rq.Get(authorUser) == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tenderID := tenderIdReviewRegexp.FindStringSubmatch(r.RequestURI)
	if tenderID == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, _ := h.b.FetchReviews(r.Context(), &fetch.BidsRequestFetchReviewsV1{Username: rq.Get(username), TenderId: tenderID[1], OrganizationId: rq.Get(organizationID), AuthorUsername: rq.Get(authorUser)})

	marshalOptions := protojson.MarshalOptions{
		Multiline:       true, // для красивого форматирования
		Indent:          "  ", // отступ двумя пробелами
		UseProtoNames:   true, // использовать оригинальные названия из .proto
		EmitUnpopulated: true, // включать пустые поля
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
