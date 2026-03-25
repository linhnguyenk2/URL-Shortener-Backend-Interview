package main

import (
	"os"
	"time"

	"urlshortener/internal/config"

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

	// Setup Gin
	r := gin.New()
	r.Use(ZerologLogger(logger))
	r.Use(gin.Recovery())

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
