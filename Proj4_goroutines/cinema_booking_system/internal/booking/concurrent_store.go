package booking

import (
	"sync"
	"context"
)
// implements BookingStore interface

var _ BookingStore = (*ConcurrentStore)(nil)  // compile-time check for interface implementation

type ConcurrentStore struct{
	bookings map[string]Booking   //  storage for bookings

	sync.RWMutex 

}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: make(map[string]Booking),
	}
}

func (s *ConcurrentStore) Book(b Booking) (Booking, error) {
	// check if seat taken, return err else book the seat
	s.Lock()
	defer s.Unlock()  // ensure we release the lock even if we return early
	if _, exists:= s.bookings[b.SeatID]; exists {
		return Booking{}, ErrSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return b,  nil
}

func (s *ConcurrentStore) ListBookings(movieID string) []Booking {
	
	s.RLock()    // Blocks writers (Lock()) while reading is in progress. Does not block other readers
	defer s.RUnlock()    // Book() calls will wait until all readers finish

	var result []Booking
	for _, b := range s.bookings{
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}


func (s *ConcurrentStore) Release(ctx context.Context, sessionID string, userID string) error {
	return nil  // just for sake
}

func (s *ConcurrentStore) Confirm(ctx context.Context, sessionID string, userID string) (Booking, error) {
	return Booking{}, nil  // just for sake
}
