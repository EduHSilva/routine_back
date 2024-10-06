package category

import (
	"github.com/Eduardo681/go_routine/helper"
)

type CreateCategoryRequest struct {
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
}

func (r CreateCategoryRequest) Validate() error {

	if r.Title == "" {
		return helper.ErrParamIsRequired("title", "string")
	}

	if r.UserID == 0 {
		return helper.ErrParamIsRequired("user_id", "uint")
	}

	return nil
}

type UpdateCategoryRequest struct {
	Title string `json:"title"`
}

func (r UpdateCategoryRequest) Validate() error {

	if r.Title == "" {
		return helper.ErrParamIsRequired("title", "string")
	}

	return nil
}
