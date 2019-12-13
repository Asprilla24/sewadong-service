// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package models

import (
	"time"
)

// Transaction data models
type Transaction struct {
	TransactionID string    `json:"transaction_id" gorm:"primary_key"`
	UserID        string    `json:"user_id"`
	ProductID     string    `json:"product_id"`
	Status        string    `json:"status"`
	DateFrom      time.Time `json:"date_from"`
	DateTo        time.Time `json:"date_to"`
	TotalPrice    float32   `json:"total_price"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}

// TransactionRequest request format for create and update transaction
type TransactionRequest struct {
	TransactionID string    `json:"transaction_id"`
	UserID        string    `json:"user_id"`
	ProductID     string    `json:"product_id"`
	Status        string    `json:"status"`
	DateFrom      time.Time `json:"date_from"`
	DateTo        time.Time `json:"date_to"`
	TotalPrice    float32   `json:"total_price"`
}

// TransactionResponse response format to send to requester
type TransactionResponse struct {
	TransactionID string          `json:"transaction_id"`
	User          UserResponse    `json:"user"`
	Product       ProductResponse `json:"product"`
	Status        string          `json:"status"`
	DateFrom      time.Time       `json:"date_from"`
	DateTo        time.Time       `json:"date_to"`
	TotalPrice    float32         `json:"total_price"`
	UpdatedAt     time.Time       `json:"updated_at"`
	CreatedAt     time.Time       `json:"created_at"`
}

// CreateTransactionResponse create transaction response format to send to requester
type CreateTransactionResponse struct {
	Result TransactionResponse `json:"result"`
}

// UpdateTransactionResponse update transaction response format to send to requester
type UpdateTransactionResponse struct {
	Result TransactionResponse `json:"result"`
}

// GetTransactionResponse get transaction response format to send to requester
type GetTransactionResponse struct {
	Result TransactionResponse `json:"result"`
}

// GetTransactionsResponse get transactions response format to send to requester
type GetTransactionsResponse struct {
	Result []TransactionResponse `json:"result"`
}
