package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bi6o/aroundhome-challenge/internal/database"
	"github.com/bi6o/aroundhome-challenge/internal/model"
	"github.com/bi6o/aroundhome-challenge/internal/repo"
	"github.com/bi6o/aroundhome-challenge/pkg/partner"

	_ "github.com/bi6o/aroundhome-challenge/docs"

	"github.com/caarlos0/env"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"

	"go.uber.org/zap"
)

// @title			Partners API
// @version		1.0
// @description	This service is responsible for managing partners.
// @host			localhost:8080
func main() {
	cfg := model.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(
			"unable to read env var 'SERVER_READ_TIMEOUT'",
			zap.Error(err),
		)
		return
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := database.Connect(ctx, cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	partnerRepo := repo.NewRepo(db)
	partnerController := partner.NewController(partnerRepo, logger)

	logger.Info("starting http server", zap.String("address", fmt.Sprintf(":%s", cfg.Port)))
	go httpServer(&cfg, logger, partnerController).ListenAndServe()

	errChan := make(chan error)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
	case err = <-errChan:
		logger.Error("got server error", zap.Error(err))
	}

	cancel()
}
