package booking

import (
	"cinema-booking-system/internal/utils"
	"encoding/json"
	"net/http"
	"time"
	"log"
)
type handler struct{
	svc *Service
}

func NewHandler(svc *Service) *handler{
	return &handler{svc:svc}
}

type holdSeatRequest struct {
	UserID string `json:"user_id"`
}

func (h *handler) ListSeats(w http.ResponseWriter, r *http.Request){

	movieID:= r.PathValue("movieID")
	bookings:= h.svc.ListBookings(movieID)

	seats:= make([]seatInfo, 0, len(bookings))
	for _, b := range bookings{
		seats = append(seats, seatInfo{
			SeatID: b.SeatID, 
			UserID: b.UserID, 
			Booked: true, 
		})
	}
	utils.WriteJSON(w, http.StatusOK, seats)

}

func (h *handler) HoldSeat(w http.ResponseWriter, r *http.Request){

	// comes as query-params
	movieID:= r.PathValue("movieID")
	seatID:= r.PathValue("seatID")


	// comes as payload --> userID
	// we can use struct to decide what all we want to decode from the request payload
	type holdRequest struct{
		UserID string `json:"user_id"`
	}
	var req holdRequest
	if err:= json.NewDecoder(r.Body).Decode(&req); err!= nil{
		return 
	}


	data:= Booking{
		MovieID :  movieID,
		SeatID:  seatID,
		UserID:   req.UserID,
	}
	session, err:= h.svc.Book(data)
	if err!= nil {
		return
	}

	response:= holdResponse{
		SeatID: seatID,
		MovieID: movieID,
		SessionID: session.ID,
		ExpiresAt: session.ExpiresAt.Format(time.RFC3339),
	}
	utils.WriteJSON(w, http.StatusOK, response)
}



func (h *handler) ConfirmSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionID")

	var req holdSeatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	if req.UserID == "" {
		return
	}

	session, err := h.svc.ConfirmSeat(r.Context(), sessionID, req.UserID)
	if err != nil {
		return
	}

	utils.WriteJSON(w, http.StatusOK, sessionResponse{
		SessionID: session.ID,
		MovieID:   session.MovieID,
		SeatID:    session.SeatID,
		UserID:    req.UserID,
		Status:    session.Status,
	})
}



func (h *handler) ReleaseSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionID")

	var req holdSeatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		return
	}
	if req.UserID == "" {
		return
	}

	err := h.svc.ReleaseSeat(r.Context(), sessionID, req.UserID)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type seatInfo struct{
	SeatID string  `json:"seat_id"`
	UserID string `json:"user_id"`
	Booked bool `json:"booked"`
}

type holdResponse struct{
	SessionID string `json:"session_id"`
	MovieID string `json:"movie_id"`
	SeatID string `json:"seat_id"`
	ExpiresAt string   `json:"expires_at"`
}

type sessionResponse struct {
	SessionID string `json:"session_id"`
	MovieID   string `json:"movie_id"`
	SeatID    string `json:"seat_id"`
	UserID    string `json:"user_id"`
	Status    string `json:"status"`
	ExpiresAt string `json:"expires_at,omitempty"`
}

