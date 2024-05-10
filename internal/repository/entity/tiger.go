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

type TigerSighting struct {
	ID        int       `db:"id"`
	UserID    string    `db:"user_id"`
	TigerID   int       `db:"tiger_id"`
	Photo     string    `db:"photo"`
	Lat       float64   `db:"lat"`
	Long      float64   `db:"long"`
	CreatedAt time.Time `db:"created_at"`
}

type Sighting struct {
	Username  string    `json:"uploaded_by" db:"username"`
	TigerName string    `json:"name" db:"tiger_name"`
	Photo     string    `json:"photo" db:"photo"`
	Lat       float64   `json:"lat" db:"lat"`
	Long      float64   `json:"long" db:"long"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
