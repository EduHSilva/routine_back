package category

import (
	"github.com/Eduardo681/go_routine/helper"
)

type CreateCategoryRequest struct {
	Title string `json:"title"`
}

func (r CreateCategoryRequest) Validate() error {

	if r.Title == "" {
		return helper.ErrParamIsRequired("title", "string")
	}
	
	return nil
}
