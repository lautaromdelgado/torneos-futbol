package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Player struct {
	Name         string    `json:"name" binding:"required"`
	Age          int       `json:"age" binding:"required"`
	CreationTime time.Time `json:"-"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	ginEngine.POST("/players", func(c *gin.Context) {
		var player Player
		if err := c.BindJSON(&player); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		player.CreationTime = time.Now().UTC()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		collection := client.Database("go-l").Collection("players")
		insertResult, err := collection.InsertOne(ctx, player)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, gin.H{"player_id": insertResult.InsertedID})
	})

	log.Fatal(ginEngine.Run(":8001"))
}
