// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"net/http"

	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/models"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
)

// GetAllRole handle get all role
func (service *Service) GetAllRole(request *restful.Request, response *restful.Response) {
	dbResult, err := service.server.GetAllRole()
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to get all role")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to get all role:", err)

		return
	}

	result := &models.GetRolesResponse{
		Result: dbResult,
	}

	writeResponse(response, http.StatusOK, result)
}

// GetRoleByID handle get role by role id
func (service *Service) GetRoleByID(request *restful.Request, response *restful.Response) {
	roleID := request.PathParameter("roleid")

	dbResult, err := service.server.GetRoleByRoleID(roleID)
	if err == dao.ErrRecordNotFound {
		respondErr(response, http.StatusNotFound, messageDatabaseError,
			"unable to retrieve role")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusNotFound}).
			Error("Unable to retrieve role:", err)

		return
	}
	if err != nil {
		respondErr(response, http.StatusInternalServerError, messageDatabaseError,
			"unable to retrieve role")

		logrus.WithFields(logrus.Fields{"module": "service", "resp": http.StatusInternalServerError}).
			Error("Unable to retrieve role:", err)

		return
	}

	result := &models.GetRoleResponse{
		Result: *dbResult,
	}

	writeResponse(response, http.StatusOK, result)
}
