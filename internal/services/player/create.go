package services

import (
	"context"
	"log"
	"os"
	"time"

	domain "github.com/lautaromdelgado/torneos-futbol/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreatePlayerService(player domain.Player) (id interface{}, err error) {
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

	return insertResult.InsertedID, nil
}
