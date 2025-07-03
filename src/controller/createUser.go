package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorlavila/my-first-crud-golang/src/configuration/logger"
	"github.com/victorlavila/my-first-crud-golang/src/configuration/validate"
	"github.com/victorlavila/my-first-crud-golang/src/controller/model/request"
	"github.com/victorlavila/my-first-crud-golang/src/controller/model/response"
	"go.uber.org/zap"
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

	logger.Info("User created successfully",
		zap.String("jouney", "createUser"),
	)

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, response)
}
