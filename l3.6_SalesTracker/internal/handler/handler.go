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
	router.GET("/items", h.GetAnalytics)

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

func (h *Handler) GetAnalytics(c *ginext.Context) {
	var req dto.AnalyticsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": err.Error()})
		return
	}

	filter := model.ItemsFilter{}

	if req.From != nil {
		fromDate, err := time.Parse("2006-01-02", *req.From)
		if err != nil {
			c.JSON(http.StatusBadRequest, ginext.H{"error": "invalid 'from' date format"})
			return
		}
		filter.From = &fromDate
	}

	if req.To != nil {
		toDate, err := time.Parse("2006-01-02", *req.To)
		if err != nil {
			c.JSON(http.StatusBadRequest, ginext.H{"error": "invalid 'to' date format"})
			return
		}
		filter.To = &toDate
	}

	filter.Category = req.Category

	analytics, err := h.service.GetAnalytics(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ginext.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, analytics)
}
