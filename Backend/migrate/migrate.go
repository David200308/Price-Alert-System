package main

import (
	"github.com/David200308/go-api/Backend/initializers"
	"github.com/David200308/go-api/Backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Payment{})
	initializers.DB.AutoMigrate(&models.Order{})
}
