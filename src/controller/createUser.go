package controller

import (
	"net/http"

	"github.com/EnzoGamaDS/Go-Platform-Project/src/configuration/logger"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/configuration/validation"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/controller/model/request"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init create user controller",
		zap.String(
			"journey",
			"CreateUser",
		))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String(
				"journey",
				"CreateUser",
			))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	response := response.UserResponse{
		ID:       "test",
		Email:    userRequest.Email,
		Name:     userRequest.Name,
		Password: userRequest.Password,
		Age:      userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String(
			"journey",
			"CreateUser",
		))
	c.JSON(http.StatusOK, response)
}
