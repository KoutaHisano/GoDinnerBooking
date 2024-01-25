package models

import "time"

type Reservation struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	RestaurantID     int       `json:"restaurant_id"`
	ReservationTime  time.Time `json:"reservation_time"`
	NumberOfPeople   int       `json:"number_of_people"`
	CancellationFlag bool      `json:"cancellation_flag"`
	VisitedFlag      bool      `json:"visited_flag"`
}
