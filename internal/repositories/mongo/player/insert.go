package player_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/lautaromdelgado/torneos-futbol/internal/domain"
)

func (p PlayerRepository) Insert(player domain.Player) (id interface{}, err error) {
	collection := p.Client.Database("go-l").Collection("players")
	insertResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting player: %w", err)
	}

	return insertResult.InsertedID, nil
}
