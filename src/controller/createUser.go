package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorlavila/my-first-crud-golang/src/configuration/logger"
	"github.com/victorlavila/my-first-crud-golang/src/configuration/validate"
	"github.com/victorlavila/my-first-crud-golang/src/controller/model/request"
	"github.com/victorlavila/my-first-crud-golang/src/model"
	"github.com/victorlavila/my-first-crud-golang/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init create user controller",
		zap.String("jouney", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info",
			err,
			zap.String("jouney", "createUser"),
		)
		errRest := validate.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewDOmainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("jouney", "createUser"),
	)

	c.String(http.StatusOK, "")
}
