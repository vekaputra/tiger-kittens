// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type CreateSighting struct {
	TigerID string         `json:"tigerID"`
	Lat     float64        `json:"lat"`
	Long    float64        `json:"long"`
	Photo   graphql.Upload `json:"photo"`
}

type CreateTiger struct {
	DateOfBirth string         `json:"dateOfBirth"`
	LastLat     float64        `json:"lastLat"`
	LastLong    float64        `json:"lastLong"`
	LastPhoto   graphql.Upload `json:"lastPhoto"`
	LastSeen    time.Time      `json:"lastSeen"`
	Name        string         `json:"name"`
}

type ListSighting struct {
	Data       []*Sighting `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

type ListTiger struct {
	Data       []*Tiger    `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

type Login struct {
	AccessToken string `json:"accessToken"`
	Timestamp   string `json:"timestamp"`
}

type LoginUser struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Message struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type Mutation struct {
}

type Pagination struct {
	Page      int `json:"page"`
	PerPage   int `json:"perPage"`
	TotalPage int `json:"totalPage"`
	TotalItem int `json:"totalItem"`
}

type PaginationInput struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type Query struct {
}

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Sighting struct {
	UploadedBy string    `json:"uploadedBy"`
	TigerName  string    `json:"tigerName"`
	Photo      string    `json:"photo"`
	Lat        float64   `json:"lat"`
	Long       float64   `json:"long"`
	CreatedAt  time.Time `json:"createdAt"`
}

type Tiger struct {
	ID          string    `json:"ID"`
	DateOfBirth string    `json:"dateOfBirth"`
	LastLat     float64   `json:"lastLat"`
	LastLong    float64   `json:"lastLong"`
	LastSeen    time.Time `json:"lastSeen"`
	LastPhoto   string    `json:"lastPhoto"`
	Name        *string   `json:"name,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
