package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/database"
	"github.com/jebog/stuble/models"
	"github.com/jebog/stuble/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()

	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Entry{})
}

func serveApplication() {
	router := gin.Default()

	routes.NewUserRoute(router)
	routes.NewEntryRoute(router)
	routes.NewAuthRoute(router)

	err := router.Run(os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal("Error loading using this port")
	}

	fmt.Println("Server running on port 8000")
}
