package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-wsl-project/internal/api/handlers"
	"go-wsl-project/internal/api/middleware"
	"go-wsl-project/internal/core/services"
	"go-wsl-project/internal/infrastructure/config"
	"go-wsl-project/internal/infrastructure/database"
	"go-wsl-project/internal/infrastructure/repository"
	"go-wsl-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	// init logger
	log := logger.New()

	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	// connect DB
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer db.Close()

	// deps
	itemRepo := repository.NewPostgresItemRepository(db)
	itemSvc := services.NewItemService(itemRepo)

	// router
	router := gin.New()
	router.Use(middleware.Logger(log), gin.Recovery())

	// handlers
	handlers.RegisterPing(router)
	handlers.RegisterHealth(router, db)
	handlers.RegisterList(router, itemSvc)
	handlers.RegisterAdd(router, itemSvc)

	// server
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	go func() {
		log.Infof("listening on :%s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	log.Info("server exiting")
}
