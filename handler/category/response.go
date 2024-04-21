package category

import (
	"time"
)

type ResponseData struct {
	ID        uint      `json:"id"`
	CreateAt  time.Time `json:"createAt"`
	UpdateAt  time.Time `json:"updateAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Title     string    `json:"title"`
}

type ResponseCategory struct {
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseCategories struct {
	Message string         `json:"message"`
	Data    []ResponseData `json:"data"`
}
