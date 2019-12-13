// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package models

import "time"

// User data models
type User struct {
	UserID         string    `json:"user_id" gorm:"primary_key"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	Password       string    `json:"password,omitempty" gorm:"-"`
	HashedPassword string    `json:"-" gorm:"column:password"`
	RoleID         string    `json:"role_id" gorm:"column:role_id"`
	Gender         string    `json:"gender" gorm:"column:gender"`
	PhoneNumber    string    `json:"phone_number" gorm:"column:phone_number"`
	Address        string    `json:"address" gorm:"column:address"`
	Image          string    `json:"image" gorm:"column:image"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

// UserRequest request format for create and update user
type UserRequest struct {
	UserID      string `json:"user_id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	RoleID      string `json:"role_id"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Image       string `json:"image"`
}

// UserResponse response format to send to requester
type UserResponse struct {
	UserID      string       `json:"user_id"`
	Email       string       `json:"email"`
	Username    string       `json:"username"`
	Role        RoleResponse `json:"role"`
	Gender      string       `json:"gender"`
	PhoneNumber string       `json:"phone_number"`
	Address     string       `json:"address"`
	Image       string       `json:"image"`
	UpdatedAt   time.Time    `json:"updated_at"`
	CreatedAt   time.Time    `json:"created_at"`
}

// CreateUserResponse create user response format to send to requester
type CreateUserResponse struct {
	Result UserResponse `json:"result"`
}

// UpdateUserResponse update user response format to send to requester
type UpdateUserResponse struct {
	Result UserResponse `json:"result"`
}

// GetUserResponse get user response format to send to requester
type GetUserResponse struct {
	Result UserResponse `json:"result"`
}

// GetUsersResponse get users response format to send to requester
type GetUsersResponse struct {
	Result []UserResponse `json:"result"`
}

// LoginRequest request format for login
type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

// LoginResponse response format to send to requester
type LoginResponse struct {
	Result UserResponse `json:"result"`
}
