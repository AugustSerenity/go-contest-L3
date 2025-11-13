package handler

import (
	"net/http"

	"github.com/AugustSerenity/go-contest-L3/l3.5_EventBooker/internal/handler/dto"
	"github.com/AugustSerenity/go-contest-L3/l3.5_EventBooker/internal/model"
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

	router.POST("/events", h.CreateEvent)

	return router
}

func (h *Handler) CreateEvent(c *ginext.Context) {
	var req dto.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": err.Error()})
		return
	}

	event := model.Event{
		Name:       req.Name,
		Date:       req.Date,
		Capacity:   req.Capacity,
		FreeSeats:  req.Capacity,
		PaymentTTL: req.PaymentTTL,
	}

	if err := h.service.CreateEvent(c.Request.Context(), &event); err != nil {
		c.JSON(http.StatusInternalServerError, ginext.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.EventResponse{
		ID:        event.ID,
		Name:      event.Name,
		Date:      event.Date,
		Capacity:  event.Capacity,
		FreeSeats: event.FreeSeats,
	})

}
