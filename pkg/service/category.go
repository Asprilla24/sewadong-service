// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"net/http"
	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"

	"github.com/emicklei/go-restful"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CreateCategory handle create new category
func (service *Service) CreateCategory(request *restful.Request, response *restful.Response) {
	var req models.CategoryRequest
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

	category := models.Category{
		CategoryID: newUUID.String(),
		Name:       req.Name,
	}

	createdCategory, err := service.server.CreateCategory(category)
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to create category")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to create category:", err)

		return
	}

	result := &models.CreateCategoryResponse{
		Result: *createdCategory,
	}

	writeResponse(response, http.StatusCreated, result)
}

// UpdateCategory handle update category request
func (service *Service) UpdateCategory(request *restful.Request, response *restful.Response) {
	var req models.CategoryRequest
	if err := request.ReadEntity(&req); err != nil {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"unable parse request body")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to parse request body:", err)

		return
	}

	category := models.Category(req)

	updatedCategory, err := service.server.UpdateCategory(category)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"category not found")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to update category:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to update category")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to update category:", err)

		return
	}

	result := &models.UpdateCategoryResponse{
		Result: *updatedCategory,
	}

	writeResponse(response, http.StatusAccepted, result)
}

// DeleteCategory handle delete product request
func (service *Service) DeleteCategory(request *restful.Request, response *restful.Response) {
	var req models.CategoryRequest
	if err := request.ReadEntity(&req); err != nil {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"unable parse request body")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to parse request body:", err)

		return
	}

	category := models.Category(req)

	err := service.server.DeleteCategory(category)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"category not found")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to update category:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to update category")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to update category:", err)

		return
	}

	result := &models.DeleteCategoryResponse{
		Result: models.CategoryResponse{},
	}

	writeResponse(response, http.StatusNoContent, result)
}

// GetAllCategory handle get all category
func (service *Service) GetAllCategory(request *restful.Request, response *restful.Response) {
	dbResult, err := service.server.GetAllCategory()
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to get all category")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to get all category:", err)

		return
	}

	result := &models.GetCategoriesResponse{
		Result: dbResult,
	}

	writeResponse(response, http.StatusOK, result)
}

// GetCategoryByCategoryID handle get category by category id
func (service *Service) GetCategoryByCategoryID(request *restful.Request, response *restful.Response) {
	categoryID := request.PathParameter("categoryid")

	dbResult, err := service.server.GetCategoryByCategoryID(categoryID)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve category")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Unable to retrieve category:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to retrieve category")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to retrieve category:", err)

		return
	}

	result := &models.GetCategoryResponse{
		Result: *dbResult,
	}

	writeResponse(response, http.StatusOK, result)
}
