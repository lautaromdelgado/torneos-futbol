package services

import (
	"fmt"
	"log"
	"time"

	domain "github.com/lautaromdelgado/torneos-futbol/internal/domain"
)

func (s Services) Create(player domain.Player) (id interface{}, err error) {
	// Lógica de negocio
	// Guardar en el repositorio
	// Responder con el identificador del nuevo recurso creado
	player.CreationTime = time.Now().UTC()

	InsertedID, err := s.Repo.Insert(player)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating player: %w", err)
	}

	return InsertedID, nil
}
