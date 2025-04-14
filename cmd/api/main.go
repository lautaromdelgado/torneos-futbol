package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lautaromdelgado/torneos-futbol/cmd/api/handlers/player"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	playerHandler := player.Handler{}

	ginEngine.POST("/players", playerHandler.CreatePlayer)

	log.Fatal(ginEngine.Run(":8001"))
}
