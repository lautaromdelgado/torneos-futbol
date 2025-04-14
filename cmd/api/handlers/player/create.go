package player

import (
	"net/http"

	domain "github.com/lautaromdelgado/torneos-futbol/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreatePlayer(c *gin.Context) {
	// Traducir el request
	var playerCreateParams domain.Player
	if err := c.BindJSON(&playerCreateParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Consumir un servicio externo
	insertedID, err := h.Services.CreatePlayerService(playerCreateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops, something went wrong"})
		return
	}

	// Traducir el response
	c.JSON(200, gin.H{"player_id": insertedID})
}
