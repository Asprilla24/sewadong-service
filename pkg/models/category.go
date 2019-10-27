// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package models

// Category data models
type Category struct {
	CategoryID string `json:"category_id" gorm:"primary_key"`
	Name       string `json:"name"`
}

// CategoryRequest request format to send to requester
type CategoryRequest struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

// CategoryResponse response format to send to requester
type CategoryResponse struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

// GetCategoryResponse get category response format to send to requester
type GetCategoryResponse struct {
	Result CategoryResponse `json:"result"`
}

// GetCategoriesResponse get categories response format to send to requester
type GetCategoriesResponse struct {
	Result []CategoryResponse `json:"result"`
}

// CreateCategoryResponse create category response format to send to requester
type CreateCategoryResponse struct {
	Result CategoryResponse `json:"result"`
}

// UpdateCategoryResponse update category response format to send to requester
type UpdateCategoryResponse struct {
	Result CategoryResponse `json:"result"`
}

// DeleteCategoryResponse delete category response format to send to requester
type DeleteCategoryResponse struct {
	Result CategoryResponse `json:"result"`
}
