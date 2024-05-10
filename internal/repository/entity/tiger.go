package entity

import "time"

type Tiger struct {
	ID          int       `db:"id"`
	DateOfBirth time.Time `db:"date_of_birth"`
	LastLat     float64   `db:"last_lat"`
	LastLong    float64   `db:"last_long"`
	LastSeen    time.Time `db:"last_seen"`
	LastPhoto   string    `db:"last_photo"`
	Name        string    `db:"name"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
