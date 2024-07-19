package handler

import (
	"encoding/json"
	contr "example/zota_assignment/controller"
	"example/zota_assignment/mapper"
	"example/zota_assignment/model"
	"net/http"
	"regexp"
)

var (
	payOrderRe    = regexp.MustCompile(`^\/createOrder[\/]*$`)
	statusCheckRe = regexp.MustCompile(`^\/check\/(\w+)$`)
)

type Handler interface {
	CreateOrder(http.ResponseWriter, *http.Request)
	CheckStatus(http.ResponseWriter, *http.Request)
}

type handler struct {
	controller contr.Controller
}

func NewHandler() *handler {
	return &handler{
		controller: *contr.NewController(),
	}
}

func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := &model.Order{}

	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	req := mapper.OrderToZotaEntity(*order)
	body, err := h.controller.DepositRequest(req)
	if err != nil {
		w.Write([]byte(`{"error": "Something went wrong!"}`))
		return
	}

	w.Write(body)
}

func (h *handler) CheckStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Checking status..."))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "Not found!"}`))
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost && payOrderRe.MatchString(r.URL.Path):
		h.CreateOrder(w, r)
		return
	case r.Method == http.MethodGet && statusCheckRe.MatchString(r.URL.Path):
		h.CheckStatus(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}
