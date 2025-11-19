package memberships

import (
	"net/http"

	"github.com/arashiaslan/music-catalog-go/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {
	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.service.Login(request)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, memberships.LoginResponse{AccessToken: accessToken})
}
