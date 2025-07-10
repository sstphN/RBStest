package handlers

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterHealth(r *gin.Engine, db *sql.DB) {
	r.GET("/health", func(c *gin.Context) {
		start := time.Now()
		err := db.Ping()
		status := "ok"
		if err != nil {
			status = "db_error"
		}

		c.JSON(200, gin.H{
			"status":    status,
			"timestamp": time.Now().UTC(),
			"latency":   time.Since(start).String(),
		})
	})
}
