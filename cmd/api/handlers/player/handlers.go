package player

import ports "github.com/lautaromdelgado/torneos-futbol/internal/ports"

type Handler struct {
	Services ports.PlayerService
}

func NewHandler(services ports.PlayerService) *Handler {
	return &Handler{Services: services}
}
