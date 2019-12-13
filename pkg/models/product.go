// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package models

import (
	"time"
)

// Product data models
type Product struct {
	ProductID   string    `json:"product_id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	PaymentID   string    `json:"payment_id"`
	CategoryID  string    `json:"category_id"`
	UserID      string    `json:"user_id"`
	Image       string    `json:"image"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// ProductRequest request format for create and update product
type ProductRequest struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	PaymentID   string  `json:"payment_id"`
	CategoryID  string  `json:"category_id"`
	UserID      string  `json:"user_id"`
	Image       string  `json:"image"`
}

// ProductResponse response format to send to requester
type ProductResponse struct {
	ProductID   string           `json:"product_id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float32          `json:"price"`
	PaymentID   string           `json:"payment_id"`
	Category    CategoryResponse `json:"category"`
	User        UserResponse     `json:"user"`
	Image       string           `json:"image"`
	UpdatedAt   time.Time        `json:"updated_at"`
	CreatedAt   time.Time        `json:"created_at"`
}

// CreateProductResponse create product response format to send to requester
type CreateProductResponse struct {
	Result ProductResponse `json:"result"`
}

// UpdateProductResponse update product response format to send to requester
type UpdateProductResponse struct {
	Result ProductResponse `json:"result"`
}

// DeleteProductResponse delete product response format to send to requester
type DeleteProductResponse struct {
	Result ProductResponse `json:"result"`
}

// GetProductResponse get product response format to send to requester
type GetProductResponse struct {
	Result ProductResponse `json:"result"`
}

// GetProductsResponse get products response format to send to requester
type GetProductsResponse struct {
	Result []ProductResponse `json:"result"`
}
