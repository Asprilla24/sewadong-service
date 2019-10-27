// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package postgres

import (
	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"

	"github.com/jinzhu/gorm"
)

// CreateProduct create new product
func (db *DB) CreateProduct(product models.Product) (*models.Product, error) {
	err := db.conn.Create(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// UpdateProduct update existing product
func (db *DB) UpdateProduct(product models.Product) (*models.Product, error) {
	dbResult := db.conn.Model(&product).Update(&product)
	if dbResult.RowsAffected == 0 {
		return nil, dao.ErrRecordNotFound
	}

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return &product, nil
}

// DeleteProduct delete existing product
func (db *DB) DeleteProduct(product models.Product) error {
	dbResult := db.conn.Delete(&product)
	if dbResult.RowsAffected == 0 {
		return dao.ErrRecordNotFound
	}

	if dbResult.Error != nil {
		return dbResult.Error
	}

	return nil
}

// GetProductByProductID get product from database by ID
func (db *DB) GetProductByProductID(productID string) (*models.Product, error) {
	var product models.Product
	err := db.conn.Where("product_id=?", productID).Find(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	return &product, err
}

// GetProductByCategoryID get products from database by category ID
func (db *DB) GetProductByCategoryID(categoryID string) ([]models.Product, error) {
	var products []models.Product
	err := db.conn.Where("category_id=?", categoryID).Find(&products).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	return products, err
}

// GetProductLikeName get products from database where name like identifier
func (db *DB) GetProductLikeName(identifier string) ([]models.Product, error) {
	var products []models.Product
	err := db.conn.Where("name LIKE ?", "%"+identifier+"%").Find(&products).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	return products, err
}

// GetAllProduct get all product from database
func (db *DB) GetAllProduct() ([]models.Product, error) {
	var allProduct []models.Product
	err := db.conn.Find(&allProduct).Error

	return allProduct, err
}
