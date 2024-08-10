package service

import (
	"booking/internal/model"
	"errors"
	"sync"
	"time"

	"github.com/gofrs/uuid/v5"
)

var ErrRoomNotAvailable = errors.New("room is not available for requested days")

type OrderRepository interface {
	InsertOrder(order model.Order) error
	IsRoomAvailable(hotel model.Hotel, room model.Room, from, to time.Time) (bool, error)
}

type Order struct {
	repo OrderRepository
	mu   sync.Mutex
}

func NewOrderService(repo OrderRepository) *Order {
	return &Order{
		repo: repo,
	}
}

func (s *Order) Create(hotel model.Hotel, room model.Room, from, to time.Time, user model.User) (model.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	isAvailable, err := s.repo.IsRoomAvailable(hotel, room, from, to)
	if err != nil {
		return model.Order{}, err
	}

	if !isAvailable {
		return model.Order{}, ErrRoomNotAvailable
	}

	id, _ := uuid.NewV7()

	order := model.Order{
		ID:    id.String(),
		Hotel: hotel,
		Room:  room,
		User:  user,
		From:  from,
		To:    to,
	}

	err = s.repo.InsertOrder(order)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}
