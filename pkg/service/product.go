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

// CreateProduct handle create new product
func (service *Service) CreateProduct(request *restful.Request, response *restful.Response) {
	var req models.ProductRequest
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

	product := models.Product{
		ProductID:   newUUID.String(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		PaymentID:   req.PaymentID,
		CategoryID:  req.CategoryID,
		UserID:      req.UserID,
		Image:       req.Image,
		UpdatedAt:   time.Now().UTC(),
		CreatedAt:   time.Now().UTC(),
	}

	createdProduct, err := service.server.CreateProduct(product)
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to create product")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to create product:", err)

		return
	}

	productResponse := models.ProductResponse(*createdProduct)

	result := &models.CreateProductResponse{
		Result: productResponse,
	}

	writeResponse(response, http.StatusCreated, result)
}

// UpdateProduct handle update product request
func (service *Service) UpdateProduct(request *restful.Request, response *restful.Response) {
	var req models.ProductRequest
	if err := request.ReadEntity(&req); err != nil {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"unable parse request body")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to parse request body:", err)

		return
	}

	product := models.Product{
		ProductID:   req.ProductID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		PaymentID:   req.PaymentID,
		CategoryID:  req.CategoryID,
		UserID:      req.UserID,
		Image:       req.Image,
	}

	updatedProduct, err := service.server.UpdateProduct(product)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"product not found")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to update product:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to update product")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to update product:", err)

		return
	}

	productResponse := models.ProductResponse(*updatedProduct)

	result := &models.UpdateProductResponse{
		Result: productResponse,
	}

	writeResponse(response, http.StatusAccepted, result)
}

// DeleteProduct handle delete product request
func (service *Service) DeleteProduct(request *restful.Request, response *restful.Response) {
	var req models.ProductRequest
	if err := request.ReadEntity(&req); err != nil {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"unable parse request body")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to parse request body:", err)

		return
	}

	product := models.Product{
		ProductID:   req.ProductID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		PaymentID:   req.PaymentID,
		CategoryID:  req.CategoryID,
		UserID:      req.UserID,
		Image:       req.Image,
	}

	err := service.server.DeleteProduct(product)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"product not found")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to update product:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to update product")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to update product:", err)

		return
	}

	productResponse := models.ProductResponse(product)

	result := &models.DeleteProductResponse{
		Result: productResponse,
	}

	writeResponse(response, http.StatusNoContent, result)
}

// GetAllProduct handle get all product
func (service *Service) GetAllProduct(request *restful.Request, response *restful.Response) {
	dbResult, err := service.server.GetAllProduct()
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to get all product")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to get all product:", err)

		return
	}

	var productsResponse []models.ProductResponse
	for _, product := range dbResult {
		productResponse := models.ProductResponse(product)
		productsResponse = append(productsResponse, productResponse)
	}

	result := &models.GetProductsResponse{
		Result: productsResponse,
	}

	writeResponse(response, http.StatusOK, result)
}

// GetProductByProductID handle get product by product id
func (service *Service) GetProductByProductID(request *restful.Request, response *restful.Response) {
	productID := request.PathParameter("productid")

	dbResult, err := service.server.GetProductByProductID(productID)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve product")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Unable to retrieve product:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to retrieve product")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to retrieve product:", err)

		return
	}

	productResponse := models.ProductResponse(*dbResult)

	result := &models.GetProductResponse{
		Result: productResponse,
	}

	writeResponse(response, http.StatusOK, result)
}

// GetProductByCategoryID handle get product by category id
func (service *Service) GetProductByCategoryID(request *restful.Request, response *restful.Response) {
	categoryID := request.PathParameter("categoryid")

	dbResult, err := service.server.GetProductByCategoryID(categoryID)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve products")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Unable to retrieve products:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to retrieve products")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to retrieve products:", err)

		return
	}

	var productsResponse []models.ProductResponse
	for _, product := range dbResult {
		productResponse := models.ProductResponse(product)
		productsResponse = append(productsResponse, productResponse)
	}

	result := &models.GetProductsResponse{
		Result: productsResponse,
	}

	writeResponse(response, http.StatusOK, result)
}

// GetProductLikeName handle get products like name
func (service *Service) GetProductLikeName(request *restful.Request, response *restful.Response) {
	identifier := request.PathParameter("identifier")

	dbResult, err := service.server.GetProductLikeName(identifier)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve products")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Unable to retrieve products:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to retrieve products")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to retrieve products:", err)

		return
	}

	var productsResponse []models.ProductResponse
	for _, product := range dbResult {
		productResponse := models.ProductResponse(product)
		productsResponse = append(productsResponse, productResponse)
	}

	result := &models.GetProductsResponse{
		Result: productsResponse,
	}

	writeResponse(response, http.StatusOK, result)
}
