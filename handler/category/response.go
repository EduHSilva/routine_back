package category

import (
	"github.com/Eduardo681/go_routine/schemas"
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

func ConvertCategoryToCategoryResponse(category *schemas.Category) ResponseData {
	return ResponseData{
		ID:        category.ID,
		CreateAt:  category.CreatedAt,
		UpdateAt:  category.UpdatedAt,
		DeletedAt: category.DeletedAt.Time,
		Title:     category.Title,
	}
}
