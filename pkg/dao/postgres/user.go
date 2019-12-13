// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package postgres

import (
	"github.com/jinzhu/gorm"
	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"
)

// CreateUser create new user
func (db *DB) CreateUser(user models.User) (*models.UserResponse, error) {
	err := db.conn.Create(&user).Error
	if err != nil {
		return nil, err
	}

	roleDetail, err := db.GetRoleByRoleID(user.RoleID)
	if err != nil {
		return nil, err
	}

	userResponse := models.UserResponse{
		UserID:      user.UserID,
		Email:       user.Email,
		Username:    user.Username,
		Role:        *roleDetail,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Image:       user.Image,
		UpdatedAt:   user.UpdatedAt,
		CreatedAt:   user.CreatedAt,
	}

	return &userResponse, nil
}

// UpdateUser update existing user
func (db *DB) UpdateUser(user models.User) (*models.UserResponse, error) {
	dbResult := db.conn.Model(&user).Update(&user)
	if dbResult.RowsAffected == 0 {
		return nil, dao.ErrRecordNotFound
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	roleDetail, err := db.GetRoleByRoleID(user.RoleID)
	if err != nil {
		return nil, err
	}

	userResponse := models.UserResponse{
		UserID:      user.UserID,
		Email:       user.Email,
		Username:    user.Username,
		Role:        *roleDetail,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Image:       user.Image,
		UpdatedAt:   user.UpdatedAt,
		CreatedAt:   user.CreatedAt,
	}

	return &userResponse, nil
}

// GetUserByUserID get user from database by ID
func (db *DB) GetUserByUserID(userID string) (*models.UserResponse, error) {
	var user models.User
	err := db.conn.Where("user_id=?", userID).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	roleDetail, err := db.GetRoleByRoleID(user.RoleID)
	if err != nil {
		return nil, err
	}

	userResponse := models.UserResponse{
		UserID:      user.UserID,
		Email:       user.Email,
		Username:    user.Username,
		Role:        *roleDetail,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Image:       user.Image,
		UpdatedAt:   user.UpdatedAt,
		CreatedAt:   user.CreatedAt,
	}

	return &userResponse, err
}

// GetUserByRoleID get user from database by role ID
func (db *DB) GetUserByRoleID(roleID string) ([]models.UserResponse, error) {
	var users []models.User
	err := db.conn.Where("role_id=?", roleID).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	roleDetail, err := db.GetRoleByRoleID(roleID)
	if err != nil {
		return nil, err
	}

	var usersResponse []models.UserResponse
	for _, user := range users {
		userResponse := models.UserResponse{
			UserID:      user.UserID,
			Email:       user.Email,
			Username:    user.Username,
			Role:        *roleDetail,
			Gender:      user.Gender,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
			Image:       user.Image,
			UpdatedAt:   user.UpdatedAt,
			CreatedAt:   user.CreatedAt,
		}

		usersResponse = append(usersResponse, userResponse)
	}

	return usersResponse, err
}

// GetUserByEmailOrUsername get user from database by email or username
func (db *DB) GetUserByEmailOrUsername(identifier string) (*models.UserResponse, error) {
	var user models.User
	err := db.conn.Where("email=? OR username=?", identifier, identifier).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	roleDetail, err := db.GetRoleByRoleID(user.RoleID)
	if err != nil {
		return nil, err
	}

	userResponse := models.UserResponse{
		UserID:      user.UserID,
		Email:       user.Email,
		Username:    user.Username,
		Role:        *roleDetail,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Image:       user.Image,
		UpdatedAt:   user.UpdatedAt,
		CreatedAt:   user.CreatedAt,
	}

	return &userResponse, err
}

// GetUserByEmailOrUsernameReturnUser get user from database by email or username
func (db *DB) GetUserByEmailOrUsernameReturnUser(identifier string) (*models.User, error) {
	var user models.User
	err := db.conn.Where("email=? OR username=?", identifier, identifier).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	return &user, err
}

// GetAllUser get all user from database
func (db *DB) GetAllUser() ([]models.UserResponse, error) {
	var allUser []models.User
	err := db.conn.Find(&allUser).Error

	var usersResponse []models.UserResponse
	for _, user := range allUser {
		roleDetail, errGetRole := db.GetRoleByRoleID(user.RoleID)
		if errGetRole != nil {
			return nil, errGetRole
		}

		userResponse := models.UserResponse{
			UserID:      user.UserID,
			Email:       user.Email,
			Username:    user.Username,
			Role:        *roleDetail,
			Gender:      user.Gender,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
			Image:       user.Image,
			UpdatedAt:   user.UpdatedAt,
			CreatedAt:   user.CreatedAt,
		}

		usersResponse = append(usersResponse, userResponse)
	}

	return usersResponse, err
}
