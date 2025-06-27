package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/victorlavila/my-first-crud-golang/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := gin.Default()

	routes.InitRoutes(&route.RouterGroup)

	if err := route.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
