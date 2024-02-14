package main

import (
	"net/http"
	"time"

	"github.com/bi6o/aroundhome-challenge/internal/model"
	"github.com/bi6o/aroundhome-challenge/pkg/partner"

	brotli "github.com/anargu/gin-brotli"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	serverWriteTimeout = 15
	serverReadTimeout  = 15
)

func httpServer(cfg *model.Config, logger *zap.Logger, partnerController *partner.Controller) *http.Server {
	var router *gin.Engine

	gin.SetMode(gin.DebugMode)
	router = gin.Default()

	router.Use(Compress())
	router.Use(gin.Recovery())

	router.POST("/partners/match", partnerController.Match)
	router.GET("/partners/:id", partnerController.Get)

	return &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		WriteTimeout: time.Duration(serverWriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(serverReadTimeout) * time.Second,
	}
}

func Compress() gin.HandlerFunc {
	return func(c *gin.Context) {
		brotliMiddleware := brotli.Brotli(brotli.DefaultCompression)
		brotliMiddleware(c)
	}
}
