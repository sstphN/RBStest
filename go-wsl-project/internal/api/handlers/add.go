package handlers

import (
	"go-wsl-project/internal/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddItemRequest struct {
	Name string `json:"name" binding:"required"`
}

func RegisterAdd(r *gin.Engine, svc *services.ItemService) {
	r.POST("/add", func(c *gin.Context) {
		var req AddItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item, err := svc.Create(c.Request.Context(), req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, item)
	})
}
