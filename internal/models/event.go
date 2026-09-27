package models

import "time"

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	EventDate   time.Time `json:"event_date"`
	Location    *string   `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
}