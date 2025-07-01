package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/victorlavila/my-first-crud-golang/src/configuration/validate"
	"github.com/victorlavila/my-first-crud-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marchal object, error=%s\n", err.Error())
		errRest := validate.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
