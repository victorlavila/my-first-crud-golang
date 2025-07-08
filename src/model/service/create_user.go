package service

import (
	"fmt"

	"github.com/victorlavila/my-first-crud-golang/src/configuration/logger"
	"github.com/victorlavila/my-first-crud-golang/src/configuration/rest_err"
	"github.com/victorlavila/my-first-crud-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init create user model", zap.String("journey", "createUser"))

	userDomain.EncryptePassword()

	fmt.Printf(userDomain.Getpassword())

	return nil
}
