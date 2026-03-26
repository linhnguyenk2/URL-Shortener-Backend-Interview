package main

import (
	"os"
	"time"

	"urlshortener/internal/config"
	"urlshortener/internal/db"
	"urlshortener/internal/handler"
	"urlshortener/internal/repository"
	"urlshortener/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	// Initialize structured logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load config")
	}

	database, err := db.InitDB(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	logger.Info().Msg("Connected to PostgreSQL successfully")

	// Dependencies Injection
	repo := repository.NewShortlinkRepository(database)
	svc := service.NewShortlinkService(repo)

	host := "http://localhost:8080"
	hdl := handler.NewShortlinkHandler(svc, host)

	// Setup Gin
	r := gin.New()
	r.Use(ZerologLogger(logger))
	r.Use(gin.Recovery())

	// API Routes
	api := r.Group("/api")
	{
		api.POST("/shortlinks", hdl.CreateShortlink)
		api.GET("/shortlinks/:id", hdl.GetDetail)
	}

	// Redirect Route
	r.GET("/shortlinks/:id", hdl.Redirect)

	// Start server
	logger.Info().Msgf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Fatal().Err(err).Msg("Server failed")
	}
}

func ZerologLogger(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Str("client_ip", c.ClientIP()).
			Msg("Incoming request")
	}
}
