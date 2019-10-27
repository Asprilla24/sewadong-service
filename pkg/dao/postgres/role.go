// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package postgres

import (
	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"

	"github.com/jinzhu/gorm"
)

// GetRoleByRoleID get role from database by ID
func (db *DB) GetRoleByRoleID(roleID string) (*models.Role, error) {
	var role models.Role
	err := db.conn.Where("role_id=?", roleID).Find(&role).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	return &role, err
}

// GetAllRole get all role from database
func (db *DB) GetAllRole() ([]models.Role, error) {
	var allRole []models.Role
	err := db.conn.Find(&allRole).Error

	return allRole, err
}
