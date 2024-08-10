package service_test

import (
	"booking/internal/repository"
	"booking/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOrderCreate(t *testing.T) {
	t.Parallel()

	const hotel = "hotel-1"
	const room = "room-1"
	const user = "testuser@localhost"

	repo := repository.NewOrderRepo()
	s := service.NewOrderService(repo)

	order, err := s.Create(hotel, room, time.Now(), time.Now().Add(24*time.Hour), user)
	require.NoError(t, err)
	require.NotEmpty(t, order.ID)

	order, err = s.Create(hotel, room, time.Now().Add(4*24*time.Hour), time.Now().Add(5*24*time.Hour), user)
	require.NoError(t, err)
	require.NotEmpty(t, order.ID)

	orders, err := repo.SelectOrders()
	require.NoError(t, err)

	require.Len(t, orders, 2)

	// we can't book same room twice
	order, err = s.Create(hotel, room, time.Now(), time.Now().Add(24*time.Hour), user)
	require.Equal(t, err, service.ErrRoomNotAvailable)
}
