// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package postgres

import (
	"github.com/jinzhu/gorm"
	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"
)

// CreateTransaction create new transaction
func (db *DB) CreateTransaction(transaction models.Transaction) (*models.TransactionResponse, error) {
	err := db.conn.Create(&transaction).Error
	if err != nil {
		return nil, err
	}

	userDetail, err := db.GetUserByUserID(transaction.UserID)
	if err != nil {
		return nil, err
	}

	productDetail, err := db.GetProductByProductID(transaction.ProductID)
	if err != nil {
		return nil, err
	}

	transactionResponse := models.TransactionResponse{
		TransactionID: transaction.TransactionID,
		User:          *userDetail,
		Product:       *productDetail,
		Status:        transaction.Status,
		DateFrom:      transaction.DateFrom,
		DateTo:        transaction.DateTo,
		TotalPrice:    transaction.TotalPrice,
		UpdatedAt:     transaction.UpdatedAt,
		CreatedAt:     transaction.CreatedAt,
	}

	return &transactionResponse, nil
}

// UpdateTransaction update existing transaction
func (db *DB) UpdateTransaction(transaction models.Transaction) (*models.TransactionResponse, error) {
	dbResult := db.conn.Model(&transaction).Update(&transaction)
	if dbResult.RowsAffected == 0 {
		return nil, dao.ErrRecordNotFound
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	userDetail, err := db.GetUserByUserID(transaction.UserID)
	if err != nil {
		return nil, err
	}

	productDetail, err := db.GetProductByProductID(transaction.ProductID)
	if err != nil {
		return nil, err
	}

	transactionResponse := models.TransactionResponse{
		TransactionID: transaction.TransactionID,
		User:          *userDetail,
		Product:       *productDetail,
		Status:        transaction.Status,
		DateFrom:      transaction.DateFrom,
		DateTo:        transaction.DateTo,
		TotalPrice:    transaction.TotalPrice,
		UpdatedAt:     transaction.UpdatedAt,
		CreatedAt:     transaction.CreatedAt,
	}

	return &transactionResponse, nil
}

// GetTransactionByTransactionID get transaction from database by ID
func (db *DB) GetTransactionByTransactionID(transactionID string) (*models.TransactionResponse, error) {
	var transaction models.Transaction
	err := db.conn.Where("transaction_id=?", transactionID).Find(&transaction).Error
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	}

	userDetail, err := db.GetUserByUserID(transaction.UserID)
	if err != nil {
		return nil, err
	}

	productDetail, err := db.GetProductByProductID(transaction.ProductID)
	if err != nil {
		return nil, err
	}

	transactionResponse := models.TransactionResponse{
		TransactionID: transaction.TransactionID,
		User:          *userDetail,
		Product:       *productDetail,
		Status:        transaction.Status,
		DateFrom:      transaction.DateFrom,
		DateTo:        transaction.DateTo,
		TotalPrice:    transaction.TotalPrice,
		UpdatedAt:     transaction.UpdatedAt,
		CreatedAt:     transaction.CreatedAt,
	}

	return &transactionResponse, err
}

// GetTransactionByUserID get products from database by user ID
func (db *DB) GetTransactionByUserID(userID string) ([]models.TransactionResponse, error) {
	var transactions []models.Transaction
	user, err := db.GetUserByUserID(userID)
	if err == gorm.ErrRecordNotFound {
		return nil, dao.ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	isTenant := user.Role.RoleID == "1"

	if isTenant {
		products, errGetProductByUser := db.GetProductByUserID(userID)
		if errGetProductByUser == gorm.ErrRecordNotFound {
			return nil, dao.ErrRecordNotFound
		}

		var listOfProductID []string
		for _, product := range products {
			listOfProductID = append(listOfProductID, product.ProductID)
		}

		err = db.conn.Where("product_id IN (?)", listOfProductID).Find(&transactions).Error
		if err == gorm.ErrRecordNotFound {
			return nil, dao.ErrRecordNotFound
		}
	} else {
		err = db.conn.Where("user_id=?", userID).Find(&transactions).Error
		if err == gorm.ErrRecordNotFound {
			return nil, dao.ErrRecordNotFound
		}
	}

	var transactionsResponse []models.TransactionResponse
	for _, transaction := range transactions {
		userDetail, errGetUser := db.GetUserByUserID(transaction.UserID)
		if errGetUser != nil {
			return nil, errGetUser
		}

		productDetail, errGetProduct := db.GetProductByProductID(transaction.ProductID)
		if errGetProduct != nil {
			return nil, errGetProduct
		}

		transactionResponse := models.TransactionResponse{
			TransactionID: transaction.TransactionID,
			User:          *userDetail,
			Product:       *productDetail,
			Status:        transaction.Status,
			DateFrom:      transaction.DateFrom,
			DateTo:        transaction.DateTo,
			TotalPrice:    transaction.TotalPrice,
			UpdatedAt:     transaction.UpdatedAt,
			CreatedAt:     transaction.CreatedAt,
		}

		transactionsResponse = append(transactionsResponse, transactionResponse)
	}

	return transactionsResponse, err
}
