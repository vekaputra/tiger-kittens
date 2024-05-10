package model

import (
	"time"

	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
)

type CreateTigerRequest struct {
	LastLat     float64   `form:"last_lat" validate:"required,min=-90,max=90"`
	LastLong    float64   `form:"last_long" validate:"required,min=-180,max=180"`
	LastSeen    time.Time `form:"last_seen" validate:"required"`
	Name        string    `form:"name" validate:"required,min=3,max=64"`
	DateOfBirth time.Time `validate:"required"`
	LastPhoto   string
}

type ListTigerResponse struct {
	Data       []entity.Tiger     `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type CreateSightingRequest struct {
	TigerID int     `param:"tigerID" validate:"required"`
	Lat     float64 `form:"lat" validate:"required,min=-90,max=90"`
	Long    float64 `form:"long"  validate:"required,min=-180,max=180"`
	UserID  string  `validate:"required"`
	Photo   string
}

type ListSightingRequest struct {
	PaginationRequest
	TigerID int `param:"tigerID" validate:"required"`
}

type ListSightingResponse struct {
	Data       []entity.Sighting  `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
