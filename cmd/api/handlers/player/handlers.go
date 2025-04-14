package player

import ports "github.com/lautaromdelgado/torneos-futbol/internal/ports"

type Handler struct {
	Services ports.PlayerService
}

func NewHandler(port ports.PlayerService) *Handler {
	return &Handler{Services: port}
}
