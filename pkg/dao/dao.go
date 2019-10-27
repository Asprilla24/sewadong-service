// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package dao

import (
	"errors"

	"sewadong-service/pkg/models"
)

var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = errors.New("record not found")
)

// Server contains server info interfaces
type Server interface {
	// User Dao Interface
	CreateUser(user models.User) (*models.User, error)
	GetUserByUserID(userID string) (*models.User, error)
	GetUserByEmailOrUsername(identifier string) (*models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateUser(user models.User) (*models.User, error)

	// Role Dao Interface
	GetRoleByRoleID(roleID string) (*models.Role, error)
	GetAllRole() ([]models.Role, error)

	// Product Dao Interface
	CreateProduct(product models.Product) (*models.Product, error)
	GetProductByProductID(productID string) (*models.Product, error)
	GetProductByCategoryID(categoryID string) ([]models.Product, error)
	GetProductLikeName(identifier string) ([]models.Product, error)
	GetAllProduct() ([]models.Product, error)
	UpdateProduct(product models.Product) (*models.Product, error)
	DeleteProduct(product models.Product) error

	// Category Dao Interface
	CreateCategory(category models.Category) (*models.Category, error)
	GetCategoryByCategoryID(categoryID string) (*models.Category, error)
	GetAllCategory() ([]models.Category, error)
	UpdateCategory(category models.Category) (*models.Category, error)
	DeleteCategory(category models.Category) error
}
