package service

import (
	"github.com/victorlavila/my-first-crud-golang/src/configuration/rest_err"
	"github.com/victorlavila/my-first-crud-golang/src/model"
)

func NewDOmainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
