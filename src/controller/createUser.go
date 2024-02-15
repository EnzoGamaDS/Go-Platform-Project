package controller

import (
	"fmt"
	"net/http"

	"github.com/EnzoGamaDS/Go-Platform-Project/src/configuration/validation"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
	c.JSON(http.StatusOK, gin.H{"user": userRequest})
}
