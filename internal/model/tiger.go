package model

import "time"

type CreateTigerRequest struct {
	LastLat     float64   `form:"last_lat"`
	LastLong    float64   `form:"last_long"`
	LastSeen    time.Time `form:"last_seen"`
	Name        string    `form:"name"`
	DateOfBirth time.Time
	LastPhoto   string
}
