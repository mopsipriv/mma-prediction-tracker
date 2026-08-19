package models

import "time"

type Prediction struct {
	ID                int64     `json:"id"`
	PredictorID       int64     `json:"predictor_id"`
	FightID           int64     `json:"fight_id"`
	PredictedWinnerID int64     `json:"predicted_winner_id"`
	Probability       float64   `json:"probability"`
	CreatedAt         time.Time `json:"created_at"`
}