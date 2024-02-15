package model

import (
	"fmt"

	"github.com/EnzoGamaDS/Go-Platform-Project/src/configuration/logger"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
