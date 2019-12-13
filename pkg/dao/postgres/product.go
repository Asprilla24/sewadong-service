// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package postgres

import (
	"github.com/jinzhu/gorm"
	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"
	"strings"
)

// CreateProduct create new product
func (db *DB) CreateProduct(product models.Product) (*models.ProductResponse, error) {
	err := db.conn.Create(&product).Error
	if err != nil {
		return nil, err
	}

	userResponse, err := db.GetUserByUserID(product.UserID)
	if err != nil {
		return nil, err
	}

	categoryResponse, err := db.GetCategoryByCategoryID(product.CategoryID)
	if err != nil {
		return nil, err
	}

	productResponse := models.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		PaymentID:   product.PaymentID,
		Category:    *categoryResponse,
		User:        *userResponse,
		Image:       product.Image,
		UpdatedAt:   product.UpdatedAt,
		CreatedAt:   product.CreatedAt,
	}

	return &productResponse, nil
}

// UpdateProduct update existing product
func (db *DB) UpdateProduct(product models.Product) (*models.ProductResponse, error) {
	dbResult := db.conn.Model(&product).Update(&product)
	if dbResult.RowsAffected == 0 {
		return nil, dao.ErrRecordNotFound
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	userResponse, err := db.GetUserByUserID(product.UserID)
	if err != nil {
		return nil, err
	}

	categoryResponse, err := db.GetCategoryByCategoryID(product.CategoryID)
	if err != nil {
		return nil, err
	}

	productResponse := models.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		PaymentID:   product.PaymentID,
		Category:    *categoryResponse,
		User:        *userResponse,
		Image:       product.Image,
		UpdatedAt:   product.UpdatedAt,
		CreatedAt:   product.CreatedAt,
	}

	return &productResponse, nil
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
func (db *DB) GetProductByProductID(productID string) (*models.ProductResponse, error) {
	var product models.Product
	err := db.conn.Where("product_id=?", productID).Find(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	userResponse, err := db.GetUserByUserID(product.UserID)
	if err != nil {
		return nil, err
	}

	categoryResponse, err := db.GetCategoryByCategoryID(product.CategoryID)
	if err != nil {
		return nil, err
	}

	productResponse := models.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		PaymentID:   product.PaymentID,
		Category:    *categoryResponse,
		User:        *userResponse,
		Image:       product.Image,
		UpdatedAt:   product.UpdatedAt,
		CreatedAt:   product.CreatedAt,
	}

	return &productResponse, nil
}

// GetProductByCategoryID get products from database by category ID
func (db *DB) GetProductByCategoryID(categoryID string) ([]models.ProductResponse, error) {
	var products []models.Product
	err := db.conn.Where("category_id=?", categoryID).Find(&products).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	var productsResponse []models.ProductResponse
	for _, product := range products {
		userResponse, errGetUser := db.GetUserByUserID(product.UserID)
		if errGetUser != nil {
			return nil, errGetUser
		}

		categoryResponse, errGetCategory := db.GetCategoryByCategoryID(product.CategoryID)
		if errGetCategory != nil {
			return nil, errGetCategory
		}

		productResponse := models.ProductResponse{
			ProductID:   product.ProductID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			PaymentID:   product.PaymentID,
			Category:    *categoryResponse,
			User:        *userResponse,
			Image:       product.Image,
			UpdatedAt:   product.UpdatedAt,
			CreatedAt:   product.CreatedAt,
		}

		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse, err
}

// GetProductByUserID get products from database by user ID
func (db *DB) GetProductByUserID(userID string) ([]models.ProductResponse, error) {
	var products []models.Product
	err := db.conn.Where("user_id=?", userID).Find(&products).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	var productsResponse []models.ProductResponse
	for _, product := range products {
		userResponse, errGetUser := db.GetUserByUserID(product.UserID)
		if errGetUser != nil {
			return nil, errGetUser
		}

		categoryResponse, errGetCategory := db.GetCategoryByCategoryID(product.CategoryID)
		if errGetCategory != nil {
			return nil, errGetCategory
		}

		productResponse := models.ProductResponse{
			ProductID:   product.ProductID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			PaymentID:   product.PaymentID,
			Category:    *categoryResponse,
			User:        *userResponse,
			Image:       product.Image,
			UpdatedAt:   product.UpdatedAt,
			CreatedAt:   product.CreatedAt,
		}

		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse, err
}

// GetProductLikeName get products from database where name like identifier
func (db *DB) GetProductLikeName(identifier string) ([]models.ProductResponse, error) {
	var products []models.Product
	err := db.conn.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(identifier)+"%").Find(&products).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	var productsResponse []models.ProductResponse
	for _, product := range products {
		userResponse, errGetUser := db.GetUserByUserID(product.UserID)
		if errGetUser != nil {
			return nil, errGetUser
		}

		categoryResponse, errGetCategory := db.GetCategoryByCategoryID(product.CategoryID)
		if errGetCategory != nil {
			return nil, errGetCategory
		}

		productResponse := models.ProductResponse{
			ProductID:   product.ProductID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			PaymentID:   product.PaymentID,
			Category:    *categoryResponse,
			User:        *userResponse,
			Image:       product.Image,
			UpdatedAt:   product.UpdatedAt,
			CreatedAt:   product.CreatedAt,
		}

		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse, err
}

// GetAllProduct get all product from database
func (db *DB) GetAllProduct() ([]models.ProductResponse, error) {
	var allProduct []models.Product
	err := db.conn.Find(&allProduct).Error

	var productsResponse []models.ProductResponse
	for _, product := range allProduct {
		userResponse, errGetUser := db.GetUserByUserID(product.UserID)
		if errGetUser != nil {
			return nil, errGetUser
		}

		categoryResponse, errGetCategory := db.GetCategoryByCategoryID(product.CategoryID)
		if errGetCategory != nil {
			return nil, errGetCategory
		}

		productResponse := models.ProductResponse{
			ProductID:   product.ProductID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			PaymentID:   product.PaymentID,
			Category:    *categoryResponse,
			User:        *userResponse,
			Image:       product.Image,
			UpdatedAt:   product.UpdatedAt,
			CreatedAt:   product.CreatedAt,
		}

		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse, err
}
