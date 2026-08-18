package models

import "time"

type Fight struct {
	ID          int64     `json:"id"`
	EventID     int64     `json:"event_id"`
	Fighter1ID  int64     `json:"fighter_1_id"`
	Fighter2ID  int64     `json:"fighter_2_id"`
	WinnerID    *int64    `json:"winner_id"`
	WeightClass string    `json:"weight_class"`
	CreatedAt   time.Time `json:"created_at"`
}