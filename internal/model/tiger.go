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
