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
	CreateUser(user models.User) (*models.UserResponse, error)
	GetUserByUserID(userID string) (*models.UserResponse, error)
	GetUserByRoleID(roleID string) ([]models.UserResponse, error)
	GetUserByEmailOrUsername(identifier string) (*models.UserResponse, error)
	GetUserByEmailOrUsernameReturnUser(identifier string) (*models.User, error)
	GetAllUser() ([]models.UserResponse, error)
	UpdateUser(user models.User) (*models.UserResponse, error)

	// Role Dao Interface
	GetRoleByRoleID(roleID string) (*models.RoleResponse, error)
	GetAllRole() ([]models.RoleResponse, error)

	// Product Dao Interface
	CreateProduct(product models.Product) (*models.ProductResponse, error)
	GetProductByProductID(productID string) (*models.ProductResponse, error)
	GetProductByCategoryID(categoryID string) ([]models.ProductResponse, error)
	GetProductByUserID(userID string) ([]models.ProductResponse, error)
	GetProductLikeName(identifier string) ([]models.ProductResponse, error)
	GetAllProduct() ([]models.ProductResponse, error)
	UpdateProduct(product models.Product) (*models.ProductResponse, error)
	DeleteProduct(product models.Product) error

	// Category Dao Interface
	CreateCategory(category models.Category) (*models.CategoryResponse, error)
	GetCategoryByCategoryID(categoryID string) (*models.CategoryResponse, error)
	GetAllCategory() ([]models.CategoryResponse, error)
	UpdateCategory(category models.Category) (*models.CategoryResponse, error)
	DeleteCategory(category models.Category) error

	// Transaction Dao Interface
	CreateTransaction(transaction models.Transaction) (*models.TransactionResponse, error)
	GetTransactionByTransactionID(transactionID string) (*models.TransactionResponse, error)
	GetTransactionByUserID(userID string) ([]models.TransactionResponse, error)
	UpdateTransaction(transaction models.Transaction) (*models.TransactionResponse, error)
}
