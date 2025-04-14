package ports

import "github.com/lautaromdelgado/torneos-futbol/internal/domain"

type PlayerService interface {
	CreatePlayerService(player domain.Player) (id interface{}, err error)
}
