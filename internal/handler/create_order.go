package handler

import (
	"booking/internal/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateOrderRequest struct {
	Hotel string    `json:"hotel_id" validate:"required"`
	Room  string    `json:"room_id" validate:"required"`
	Email string    `json:"email" validate:"required,email"`
	From  time.Time `json:"from" validate:"required"`
	To    time.Time `json:"to" validate:"required"`
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.setErrorResponse(w, http.StatusUnprocessableEntity, []error{err})
		return
	}

	isValid, errs := h.validateCreateOrder(req)
	if !isValid {
		h.setErrorResponse(w, http.StatusUnprocessableEntity, errs)
		return
	}

	_, err = h.orderService.Create(model.Hotel(req.Hotel), model.Room(req.Room), req.From, req.To, model.User(req.Email))
	if err != nil {
		h.setErrorResponse(w, http.StatusUnprocessableEntity, []error{err})
		return
	}

	h.setSuccessResponse(w)
}

// TODO: validate hotels, rooms and users
func (h *Handler) validateCreateOrder(req CreateOrderRequest) (bool, []error) {
	var errs []error
	err := h.validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, err)
		}

		return false, errs
	}
	return true, nil
}
