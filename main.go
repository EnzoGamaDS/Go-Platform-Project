package main 

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/EnzoGamaDS/Go-Platform-Project/src/controller/routes"
	"github.com/gin-gonic/gin"
)

func  main()  {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}