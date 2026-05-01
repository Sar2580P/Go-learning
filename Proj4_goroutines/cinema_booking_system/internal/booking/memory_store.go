package booking

import (
	"context"
)

var _ BookingStore = (*MemoryStore)(nil)  // compile-time check for interface implementation

type MemoryStore struct{
	bookings map[string]Booking   // in-memory storage for bookings
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: make(map[string]Booking),
	}
}

func (s *MemoryStore) Book(b Booking) (Booking, error) {
	// check if seat taken, return err else book the seat

	if _, exists:= s.bookings[b.SeatID]; exists {
		return  Booking{}, ErrSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return b, nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	
	var result []Booking
	for _, b := range s.bookings{
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}


func (s *MemoryStore) Release(ctx context.Context, sessionID string, userID string) error {
	return nil  // just for sake
}

func (s *MemoryStore) Confirm(ctx context.Context, sessionID string, userID string) (Booking, error) {
	return Booking{}, nil  // just for sake
}
