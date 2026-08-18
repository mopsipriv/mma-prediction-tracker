package models

type Fighter struct {
    ID        int64     `json:"id"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    Nickname  *string   `json:"nickname"`
    WeightClass string  `json:"weight_class"`
    CreatedAt time.Time `json:"created_at"`
}