package services

import "github.com/lautaromdelgado/torneos-futbol/internal/ports"

type Services struct {
	Repo ports.PlayerRepository
}

func NewService(repo ports.PlayerRepository) *Services {
	return &Services{Repo: repo}
}
