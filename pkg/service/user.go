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

// CreateUser handle create new user
func (service *Service) CreateUser(request *restful.Request, response *restful.Response) {
	var req models.UserRequest
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

	user := models.User{
		UserID:      newUUID.String(),
		Email:       req.Email,
		Username:    req.Username,
		RoleID:      req.RoleID,
		Gender:      req.Gender,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Image:       req.Image,
		UpdatedAt:   time.Now().UTC(),
		CreatedAt:   time.Now().UTC(),
	}

	user.HashedPassword, err = encryptPassword([]byte(req.Password))
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageServerError,
			"unable to encrypt password")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to encrypt password:", err)

		return
	}

	createdUser, err := service.server.CreateUser(user)
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to create user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to create user:", err)

		return
	}

	userResponse := models.UserResponse{
		UserID:      createdUser.UserID,
		Email:       createdUser.Email,
		Username:    createdUser.Username,
		RoleID:      createdUser.RoleID,
		Gender:      createdUser.Gender,
		PhoneNumber: createdUser.PhoneNumber,
		Address:     createdUser.Address,
		Image:       createdUser.Image,
		UpdatedAt:   createdUser.UpdatedAt,
		CreatedAt:   createdUser.CreatedAt,
	}

	result := &models.CreateUserResponse{
		Result: userResponse,
	}

	writeResponse(response, http.StatusCreated, result)
}

// UpdateUser handle update user request
func (service *Service) UpdateUser(request *restful.Request, response *restful.Response) {
	var req models.UserRequest
	if err := request.ReadEntity(&req); err != nil {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"unable parse request body")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to parse request body:", err)

		return
	}

	user := models.User{
		UserID:      req.UserID,
		Email:       req.Email,
		Username:    req.Username,
		RoleID:      req.RoleID,
		Gender:      req.Gender,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Image:       req.Image,
	}

	updatedUser, err := service.server.UpdateUser(user)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"user not found")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to update user:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to update user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to update user:", err)

		return
	}

	userResponse := models.UserResponse{
		UserID:      updatedUser.UserID,
		Email:       updatedUser.Email,
		Username:    updatedUser.Username,
		RoleID:      updatedUser.RoleID,
		Gender:      updatedUser.Gender,
		PhoneNumber: updatedUser.PhoneNumber,
		Address:     updatedUser.Address,
		Image:       updatedUser.Image,
		UpdatedAt:   updatedUser.UpdatedAt,
		CreatedAt:   updatedUser.CreatedAt,
	}

	result := &models.UpdateUserResponse{
		Result: userResponse,
	}

	writeResponse(response, http.StatusOK, result)
}

// GetAllUser handle get all user
func (service *Service) GetAllUser(request *restful.Request, response *restful.Response) {
	dbResult, err := service.server.GetAllUser()
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to get all user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to get all user:", err)

		return
	}

	var usersResponse []models.UserResponse
	for _, user := range dbResult {
		userResponse := models.UserResponse{
			UserID:      user.UserID,
			Email:       user.Email,
			Username:    user.Username,
			RoleID:      user.RoleID,
			Gender:      user.Gender,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
			Image:       user.Image,
			UpdatedAt:   user.UpdatedAt,
			CreatedAt:   user.CreatedAt,
		}
		usersResponse = append(usersResponse, userResponse)
	}

	result := &models.GetUsersResponse{
		Result: usersResponse,
	}

	writeResponse(response, http.StatusOK, result)
}

// GetUserByEmailOrUsername handle get user by email or username
func (service *Service) GetUserByEmailOrUsername(request *restful.Request, response *restful.Response) {
	identifier := request.PathParameter("identifier")

	dbResult, err := service.server.GetUserByEmailOrUsername(identifier)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Unable to retrieve user:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to retrieve user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to retrieve user:", err)

		return
	}

	userResponse := models.UserResponse{
		UserID:      dbResult.UserID,
		Email:       dbResult.Email,
		Username:    dbResult.Username,
		RoleID:      dbResult.RoleID,
		Gender:      dbResult.Gender,
		PhoneNumber: dbResult.PhoneNumber,
		Address:     dbResult.Address,
		Image:       dbResult.Image,
		UpdatedAt:   dbResult.UpdatedAt,
		CreatedAt:   dbResult.CreatedAt,
	}

	result := &models.GetUserResponse{
		Result: userResponse,
	}

	writeResponse(response, http.StatusOK, result)
}

// Login handle get login request
func (service *Service) Login(request *restful.Request, response *restful.Response) {
	var req models.LoginRequest
	if err := request.ReadEntity(&req); err != nil {
		respondErr(response, http.StatusBadRequest, messageBadRequest,
			"unable parse request body")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusBadRequest}).
			Error("Unable to parse request body:", err)

		return
	}

	dbResult, err := service.server.GetUserByEmailOrUsername(req.Identifier)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Unable to retrieve user:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to retrieve user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to retrieve user:", err)

		return
	}

	err = comparePassword([]byte(req.Password), dbResult.HashedPassword)
	if err != nil {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Password does not match:", err)

		return
	}

	userResponse := models.UserResponse{
		UserID:      dbResult.UserID,
		Email:       dbResult.Email,
		Username:    dbResult.Username,
		RoleID:      dbResult.RoleID,
		Gender:      dbResult.Gender,
		PhoneNumber: dbResult.PhoneNumber,
		Address:     dbResult.Address,
		Image:       dbResult.Image,
		UpdatedAt:   dbResult.UpdatedAt,
		CreatedAt:   dbResult.CreatedAt,
	}

	result := &models.LoginResponse{
		Result: userResponse,
	}

	writeResponse(response, http.StatusOK, result)
}

// GetUserByID handle get user by user id
func (service *Service) GetUserByID(request *restful.Request, response *restful.Response) {
	userID := request.PathParameter("userid")

	dbResult, err := service.server.GetUserByUserID(userID)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Unable to retrieve user:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to retrieve user")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to retrieve user:", err)

		return
	}

	userResponse := models.UserResponse{
		UserID:      dbResult.UserID,
		Email:       dbResult.Email,
		Username:    dbResult.Username,
		Gender:      dbResult.Gender,
		PhoneNumber: dbResult.PhoneNumber,
		Address:     dbResult.Address,
		Image:       dbResult.Image,
		UpdatedAt:   dbResult.UpdatedAt,
		CreatedAt:   dbResult.CreatedAt,
	}

	result := &models.GetUserResponse{
		Result: userResponse,
	}

	writeResponse(response, http.StatusOK, result)
}
