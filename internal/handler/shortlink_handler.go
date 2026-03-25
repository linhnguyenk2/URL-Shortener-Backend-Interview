package handler

import (
	"errors"
	"net/http"

	"urlshortener/internal/service"

	"github.com/gin-gonic/gin"
)

type ShortlinkHandler struct {
	svc  service.ShortlinkService
	host string
}

func NewShortlinkHandler(svc service.ShortlinkService, host string) *ShortlinkHandler {
	return &ShortlinkHandler{svc: svc, host: host}
}

type CreateShortlinkRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

type CreateShortlinkResponse struct {
	ID       string `json:"id"`
	ShortURL string `json:"short_url"`
}

func (h *ShortlinkHandler) CreateShortlink(c *gin.Context) {
	var req CreateShortlinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	shortlink, err := h.svc.CreateShortlink(req.OriginalURL)
	if err != nil {
		if errors.Is(err, service.ErrInvalidURL) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid original url format"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create shortlink"})
		return
	}

	resp := CreateShortlinkResponse{
		ID:       shortlink.ID,
		ShortURL: h.host + "/shortlinks/" + shortlink.ID,
	}
	c.JSON(http.StatusCreated, resp)
}
