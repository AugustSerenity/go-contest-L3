package handler

import (
	"net/http"
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.6_SalesTracker/internal/handler/dto"
	"github.com/AugustSerenity/go-contest-L3/l3.6_SalesTracker/internal/model"
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
	router := ginext.New("release")
	router.Use(ginext.Logger(), ginext.Recovery())

	router.POST("/items", h.CreateItem)

	return router
}

func (h *Handler) CreateItem(c *ginext.Context) {
	var req dto.ItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": "invalid date format, must be YYYY-MM-DD"})
		return
	}

	createdItem, err := h.service.CreateItem(c, model.Item{
		Type:     req.Type,
		Category: req.Category,
		Amount:   req.Amount,
		Date:     date,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ginext.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdItem)
}
