// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package models

// Role data models
type Role struct {
	RoleID string `json:"role_id" gorm:"primary_key"`
	Name   string `json:"name"`
}

// RoleResponse response format to send to requester
type RoleResponse struct {
	RoleID string `json:"role_id"`
	Name   string `json:"name"`
}

// GetRoleResponse get role response format to send to requester
type GetRoleResponse struct {
	Result RoleResponse `json:"result"`
}

// GetRolesResponse get roles response format to send to requester
type GetRolesResponse struct {
	Result []RoleResponse `json:"result"`
}
