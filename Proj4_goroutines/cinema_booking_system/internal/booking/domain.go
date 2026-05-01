package booking

import (
	"errors"
	"time"
	"context"
)
var (
	ErrSeatAlreadyBooked = errors.New("seat already booked")
)

type Booking struct {
	ID       string
	MovieID    string
	SeatID	 string
	UserID   string
	Status   string
	ExpiresAt time.Time
}

type BookingStore interface {
	Book(b Booking) (Booking, error)
	ListBookings(movieID string) []Booking
	Confirm(ctx context.Context, sessionID string, userID string) (Booking, error)
	Release(ctx context.Context, sessionID string, userID string) error

}