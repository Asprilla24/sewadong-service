// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"sewadong-service/pkg/dao"
)

const (
	messageServerError   = "SERVER_ERROR"
	messageDatabaseError = "DB_ERROR"
	messageBadRequest    = "BAD_REQUEST"
)

// Service contains service configuration
type Service struct {
	server dao.Server
}

// New create new service instance
func New(server dao.Server) *Service {
	return &Service{server: server}
}
