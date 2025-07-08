package service

import (
	"github.com/victorlavila/my-first-crud-golang/src/configuration/rest_err"
	"github.com/victorlavila/my-first-crud-golang/src/model"
)

func (ud *userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
