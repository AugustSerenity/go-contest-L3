package handler

import (
	"net/http"

	"github.com/wb-go/wbf/ginext"
)

type Handler struct {
	service Service
}

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Router() *ginext.Engine {
	router := ginext.New()

	router.POST("/upload", h.UploadImage)

	return router
}

func (h *Handler) UploadImage(c *ginext.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": "image file not loaded"})
		return
	}
	defer file.Close()

	imageID, err := h.service.UploadImage(c.Request.Context(), file, header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ginext.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ginext.H{
		"id":     imageID,
		"status": "pending",
	})
}
