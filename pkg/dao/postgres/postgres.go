// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package postgres

import (
	"sewadong-service/pkg/dao"

	"github.com/jinzhu/gorm"
)

// DB DAO implementation with postgres database driver
type DB struct {
	conn *gorm.DB
}

// NewDB create new DAO instance with postgres database driver
func NewDB(conn *gorm.DB) dao.Server {
	return &DB{conn: conn}
}
