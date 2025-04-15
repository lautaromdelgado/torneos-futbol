package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lautaromdelgado/torneos-futbol/cmd/api/handlers/player"
	"github.com/lautaromdelgado/torneos-futbol/internal/repositories/mongo"
	player_repository "github.com/lautaromdelgado/torneos-futbol/internal/repositories/mongo/player"
	services "github.com/lautaromdelgado/torneos-futbol/internal/services/player"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	playerRepo := player_repository.NewPlayerRepository(client)
	playerSer := services.NewService(playerRepo)
	playerHandler := player.NewHandler(playerSer)

	ginEngine.POST("/players", playerHandler.CreatePlayer)

	log.Fatal(ginEngine.Run(":8001"))
}
