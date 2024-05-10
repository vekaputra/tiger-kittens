package entity

import "time"

type Tiger struct {
	ID          int       `json:"id" db:"id"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	LastLat     float64   `json:"last_lat" db:"last_lat"`
	LastLong    float64   `json:"last_long" db:"last_long"`
	LastSeen    time.Time `json:"last_seen" db:"last_seen"`
	LastPhoto   string    `json:"last_photo" db:"last_photo"`
	Name        string    `json:"name" db:"name"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
