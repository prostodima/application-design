package handler

import (
	"booking/internal/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type OrderService interface {
	Create(hotel model.Hotel, room model.Room, from, to time.Time, user model.User) (model.Order, error)
}

type Handler struct {
	orderService OrderService
	validate     *validator.Validate
}

func NewHandler(orderService OrderService) *Handler {
	return &Handler{
		orderService: orderService,
		validate:     validator.New(validator.WithRequiredStructEnabled()),
	}
}

// TODO: handle errors properly
func (h *Handler) setErrorResponse(w http.ResponseWriter, statusCode int, errs []error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errResp := make([]string, len(errs))
	for _, e := range errs {
		errResp = append(errResp, e.Error())
	}

	json.NewEncoder(w).Encode(errResp)
}

func (h *Handler) setSuccessResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
