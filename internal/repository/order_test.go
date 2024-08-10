package repository_test

import (
	"booking/internal/model"
	"booking/internal/repository"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOrderInsert(t *testing.T) {
	t.Parallel()

	repo := repository.NewOrderRepo()

	err := repo.InsertOrder(model.Order{ID: "id-1"})
	require.NoError(t, err)

	err = repo.InsertOrder(model.Order{ID: "id-2"})
	require.NoError(t, err)

	orders, err := repo.SelectOrders()
	require.NoError(t, err)

	require.Len(t, orders, 2)
	require.Equal(t, "id-1", orders[0].ID)
	require.Equal(t, "id-2", orders[1].ID)
}

func TestRoomAvailability(t *testing.T) {
	t.Parallel()

	const hotel = "hotel-1"
	const room = "room-1"

	repo := repository.NewOrderRepo()

	// create sparse orders
	err := repo.InsertOrder(model.Order{
		ID:    "id-1",
		Hotel: hotel,
		Room:  room,
		From:  date(2024, 8, 10),
		To:    date(2024, 8, 12),
	})
	require.NoError(t, err)

	err = repo.InsertOrder(model.Order{
		ID:    "id-2",
		Hotel: hotel,
		Room:  room,
		From:  date(2024, 8, 14),
		To:    date(2024, 8, 17),
	})
	require.NoError(t, err)

	orders, err := repo.SelectOrders()
	require.NoError(t, err)
	require.Len(t, orders, 2)

	tests := []struct {
		from     time.Time
		to       time.Time
		expected bool
	}{
		{
			from:     date(2024, 8, 8),
			to:       date(2024, 8, 9),
			expected: true,
		},
		{
			from:     date(2024, 8, 8),
			to:       date(2024, 8, 10),
			expected: false,
		},
		{
			from:     date(2024, 8, 17),
			to:       date(2024, 8, 19),
			expected: false,
		},
		{
			from:     date(2024, 8, 13),
			to:       date(2024, 8, 13),
			expected: true,
		},
		{
			from:     date(2024, 8, 13),
			to:       date(2024, 8, 14),
			expected: false,
		},
		{
			from:     date(2024, 8, 15),
			to:       date(2024, 8, 16),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("test-availability-%s-%s", tt.from.Format("20060102"), tt.from.Format("20060102")), func(t *testing.T) {
			ok, err := repo.IsRoomAvailable(hotel, room, tt.from, tt.to)
			require.NoError(t, err)
			require.Equal(t, tt.expected, ok)
		})
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
