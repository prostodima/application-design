package main

import (
	"booking/internal/handler"
	"booking/internal/model"
	"booking/internal/repository"
	"booking/internal/service"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Orders = []model.Order{}

func main() {
	orderRepository := repository.NewOrderRepo()
	orderService := service.NewOrderService(orderRepository)
	handler := handler.NewHandler(orderService)

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handler.CreateOrder)

	LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		LogInfo("Server closed")
	} else if err != nil {
		LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}

var logger = log.Default()

func LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Error]: %s\n", msg)
}

func LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Info]: %s\n", msg)
}
