package repository

import (
	"booking/internal/model"
	"sync"
	"time"
)

type OrderRepo struct {
	data []model.Order
	mu   sync.Mutex
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{}
}

func (r *OrderRepo) InsertOrder(order model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data = append(r.data, order)

	return nil
}

func (r *OrderRepo) SelectOrders() ([]model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.data, nil
}

func (r *OrderRepo) IsRoomAvailable(hotel model.Hotel, room model.Room, from, to time.Time) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO: validate dates order
	from = from.Truncate(24 * time.Hour)
	to = to.Truncate(24 * time.Hour)

	for _, o := range r.data {
		if o.Hotel != hotel || o.Room != room {
			continue
		}

		// found booked days in order interval
		minTo := minDate(o.To, to)
		maxFrom := maxDate(o.From, from)
		if maxFrom.Before(minTo) || maxFrom.Equal(minTo) {
			return false, nil
		}
	}

	return true, nil
}

func minDate(start, end time.Time) time.Time {
	if start.Before(end) {
		return start
	}

	return end
}

func maxDate(start, end time.Time) time.Time {
	if start.After(end) {
		return start
	}

	return end
}
