// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"net/http"

	"sewadong-service/pkg/models"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func writeResponse(response *restful.Response, statusCode int, entity interface{}) {
	err := response.WriteHeaderAndJson(statusCode, entity, restful.MIME_JSON)
	if err != nil {
		respondErr(response, http.StatusInternalServerError, "Internal Server Error", "")
		logrus.WithFields(logrus.Fields{"status_code": statusCode, "entity": entity, "mime": restful.MIME_JSON}).
			Error("unable to write response: ", err)
	}
}

func respondErr(response *restful.Response, statusCode int, message, description string) {
	err := response.WriteHeaderAndJson(statusCode, models.ErrorResponse{
		StatusCode:  statusCode,
		Message:     message,
		Description: description,
	}, restful.MIME_JSON)
	if err != nil {
		logrus.WithFields(logrus.Fields{"status_code": statusCode, "mime": restful.MIME_JSON}).
			Error("unable to write error response: ", err)
	}
}

func encryptPassword(password []byte) (string, error) {
	encrypt, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encrypt[:]), nil
}

func comparePassword(password []byte, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
	return err
}
