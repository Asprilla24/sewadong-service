// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"net/http"
	"time"

	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"

	"github.com/emicklei/go-restful"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CreateTransaction handle create new transaction
func (service *Service) CreateTransaction(request *restful.Request, response *restful.Response) {
	var req models.TransactionRequest
	if err := request.ReadEntity(&req); err != nil {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"unable parse request body")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to parse request body:", err)

		return
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageServerError,
			"unable to create uuid")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to create uuid:", err)

		return
	}

	transaction := models.Transaction{
		TransactionID: newUUID.String(),
		UserID:        req.UserID,
		ProductID:     req.ProductID,
		Status:        req.Status,
		DateFrom:      req.DateFrom,
		DateTo:        req.DateTo,
		TotalPrice:    req.TotalPrice,
		UpdatedAt:     time.Now().UTC(),
		CreatedAt:     time.Now().UTC(),
	}

	createdTransaction, err := service.server.CreateTransaction(transaction)
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to create transaction")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to create transaction:", err)

		return
	}

	result := &models.CreateTransactionResponse{
		Result: *createdTransaction,
	}

	writeResponse(response, http.StatusCreated, result)
}

// UpdateTransaction handle update transaction request
func (service *Service) UpdateTransaction(request *restful.Request, response *restful.Response) {
	var req models.TransactionRequest
	if err := request.ReadEntity(&req); err != nil {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"unable parse request body")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to parse request body:", err)

		return
	}

	transaction := models.Transaction{
		TransactionID: req.TransactionID,
		UserID:        req.UserID,
		ProductID:     req.ProductID,
		Status:        req.Status,
		DateFrom:      req.DateFrom,
		DateTo:        req.DateTo,
		TotalPrice:    req.TotalPrice,
		UpdatedAt:     time.Time{},
	}

	updatedTransaction, err := service.server.UpdateTransaction(transaction)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"transaction not found")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to update transaction:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to update transaction")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to update transaction:", err)

		return
	}

	result := &models.UpdateTransactionResponse{
		Result: *updatedTransaction,
	}

	writeResponse(response, http.StatusAccepted, result)
}

// GetTransactionByUserID handle get all transaction by user id
func (service *Service) GetTransactionByUserID(request *restful.Request, response *restful.Response) {
	userID := request.PathParameter("userid")
	dbResult, err := service.server.GetTransactionByUserID(userID)
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to get all transaction by user id")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to get all transaction by user id:", err)

		return
	}

	result := &models.GetTransactionsResponse{
		Result: dbResult,
	}

	writeResponse(response, http.StatusOK, result)
}

// GetTransactionByTransactionID handle get transaction by transaction id
func (service *Service) GetTransactionByTransactionID(request *restful.Request, response *restful.Response) {
	transactionID := request.PathParameter("transactionid")

	dbResult, err := service.server.GetTransactionByTransactionID(transactionID)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve transaction")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Unable to retrieve transaction:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to retrieve transaction")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to retrieve transaction:", err)

		return
	}

	result := &models.GetTransactionResponse{
		Result: *dbResult,
	}

	writeResponse(response, http.StatusOK, result)
}
