package handlers

import (
	"net/http"

	"go-wsl-project/internal/core/services"

	"github.com/gin-gonic/gin"
)

func RegisterList(r *gin.Engine, svc *services.ItemService) {
	r.GET("/list", func(c *gin.Context) {
		items, err := svc.List(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	})
}
