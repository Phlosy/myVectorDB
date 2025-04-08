package dto

import "time"

// Vector is a struct that represents a vector in the database.
type Vector struct {
	ID        string    `json:"id"`
	Vector    []float64 `json:"vector"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
