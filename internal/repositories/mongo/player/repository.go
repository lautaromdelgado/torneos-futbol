package player_repository

import "go.mongodb.org/mongo-driver/mongo"

type PlayerRepository struct {
	Client *mongo.Client
}

func NewPlayerRepository(client *mongo.Client) *PlayerRepository {
	return &PlayerRepository{Client: client}
}
