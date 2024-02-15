package controller

import (
	"net/http"

	"github.com/EnzoGamaDS/Go-Platform-Project/src/configuration/logger"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/configuration/validation"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/controller/model/request"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String(
			"journey",
			"CreateUser",
		))
	c.String(http.StatusOK, "")
}
