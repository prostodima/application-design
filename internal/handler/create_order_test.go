package handler_test

import (
	"booking/internal/handler"
	"booking/internal/repository"
	"booking/internal/service"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderCreate(t *testing.T) {
	t.Parallel()

	reqBody := []byte(`{"hotel_id": "reddison","room_id": "lux","email": "guest@mail.ru","from": "2024-01-02T00:00:00Z","to": "2024-01-04T00:00:00Z"}`)

	req, err := http.NewRequest("POST", "/orders", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	orderRepository := repository.NewOrderRepo()
	orderService := service.NewOrderService(orderRepository)
	handler := handler.NewHandler(orderService)
	handler.CreateOrder(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	orders, err := orderRepository.SelectOrders()
	require.NoError(t, err)
	require.Len(t, orders, 1)
}
