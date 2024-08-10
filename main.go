package main

import (
	"booking/internal/handler"
	"booking/internal/repository"
	"booking/internal/service"
	"errors"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	orderRepository := repository.NewOrderRepo()
	orderService := service.NewOrderService(orderRepository)
	handler := handler.NewHandler(orderService)

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handler.CreateOrder)

	zap.L().Info("Server listening on localhost:8080", zap.Int("port", 8080))
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		zap.L().Info("Server closed", zap.Int("port", 8080))
	} else if err != nil {
		zap.L().Error("Server failed", zap.Error(err))
		os.Exit(1)
	}
}
