package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Voce chamoua  rota de forma errada")
	c.JSON(err.Code, err)
}