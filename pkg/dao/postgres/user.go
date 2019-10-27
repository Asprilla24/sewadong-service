// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package postgres

import (
	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"

	"github.com/jinzhu/gorm"
)

// CreateUser create new user
func (db *DB) CreateUser(user models.User) (*models.User, error) {
	err := db.conn.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser update existing user
func (db *DB) UpdateUser(user models.User) (*models.User, error) {
	dbResult := db.conn.Model(&user).Update(&user)
	if dbResult.RowsAffected == 0 {
		return nil, dao.ErrRecordNotFound
	}

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return &user, nil
}

// GetUserByUserID get user from database by ID
func (db *DB) GetUserByUserID(userID string) (*models.User, error) {
	var user models.User
	err := db.conn.Where("user_id=?", userID).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	return &user, err
}

// GetUserByEmailOrUsername get user from database by email or username
func (db *DB) GetUserByEmailOrUsername(identifier string) (*models.User, error) {
	var user models.User
	err := db.conn.Where("email=? OR username=?", identifier, identifier).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	return &user, err
}

// GetAllUser get all user from database
func (db *DB) GetAllUser() ([]models.User, error) {
	var allUser []models.User
	err := db.conn.Find(&allUser).Error

	return allUser, err
}
