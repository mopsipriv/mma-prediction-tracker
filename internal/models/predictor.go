package models

import "time"

type Predictor struct {
	ID          int64     `json:"id"`
	Name		string	  `json:"name"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`	
}