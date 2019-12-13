// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package postgres

import (
	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"

	"github.com/jinzhu/gorm"
)

// CreateCategory create new category
func (db *DB) CreateCategory(category models.Category) (*models.CategoryResponse, error) {
	err := db.conn.Create(&category).Error
	if err != nil {
		return nil, err
	}

	categoryResponse := models.CategoryResponse(category)

	return &categoryResponse, nil
}

// UpdateCategory update existing category
func (db *DB) UpdateCategory(category models.Category) (*models.CategoryResponse, error) {
	dbResult := db.conn.Model(&category).Update(&category)
	if dbResult.RowsAffected == 0 {
		return nil, dao.ErrRecordNotFound
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	categoryResponse := models.CategoryResponse(category)

	return &categoryResponse, nil
}

// DeleteCategory delete existing category
func (db *DB) DeleteCategory(category models.Category) error {
	dbResult := db.conn.Delete(&category)
	if dbResult.RowsAffected == 0 {
		return dao.ErrRecordNotFound
	}
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

// GetCategoryByCategoryID get category from database by ID
func (db *DB) GetCategoryByCategoryID(categoryID string) (*models.CategoryResponse, error) {
	var category models.Category
	err := db.conn.Where("category_id=?", categoryID).Find(&category).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	categoryResponse := models.CategoryResponse(category)

	return &categoryResponse, nil
}

// GetAllCategory get all category from database
func (db *DB) GetAllCategory() ([]models.CategoryResponse, error) {
	var allCategory []models.Category
	err := db.conn.Find(&allCategory).Error

	var categoriesResponse []models.CategoryResponse
	for _, category := range allCategory {
		categoryResponse := models.CategoryResponse(category)
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse, err
}
