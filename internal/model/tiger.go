package model

import (
	"time"

	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
)

type CreateTigerRequest struct {
	LastLat     float64   `form:"last_lat"`
	LastLong    float64   `form:"last_long"`
	LastSeen    time.Time `form:"last_seen"`
	Name        string    `form:"name"`
	DateOfBirth time.Time
	LastPhoto   string
}

type ListTigerResponse struct {
	Data       []entity.Tiger     `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type CreateSightingRequest struct {
	TigerID int     `param:"tigerID"`
	Lat     float64 `form:"lat"`
	Long    float64 `form:"long"`
	Photo   string
	UserID  string
}

type ListSightingRequest struct {
	PaginationRequest
	TigerID int `param:"tigerID"`
}

type ListSightingResponse struct {
	Data       []entity.Sighting  `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
