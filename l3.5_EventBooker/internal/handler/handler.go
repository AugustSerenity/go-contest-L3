package handler

import (
	"net/http"
	"strconv"

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
	router.POST("/events/:id/book", h.BookEvent)
	router.POST("/events/:id/confirm", h.ConfirmBooking)
	router.GET("/events/:id", h.GetEvent)

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

func (h *Handler) BookEvent(c *ginext.Context) {
	var req dto.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": err.Error()})
		return
	}

	eventID, _ := strconv.Atoi(c.Param("id"))
	booking, err := h.service.BookEvent(eventID, req.Seats)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.BookingResponse{
		ID:        booking.ID,
		EventID:   booking.EventID,
		Seats:     booking.Seats,
		Paid:      booking.Paid,
		CreatedAt: booking.CreatedAt,
		ExpiresAt: booking.ExpiresAt,
	})
}

func (h *Handler) ConfirmBooking(c *ginext.Context) {
	bookingID, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.ConfirmBooking(bookingID); err != nil {
		c.JSON(http.StatusInternalServerError, ginext.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ginext.H{"status": "confirmed"})
}

func (h *Handler) GetEvent(c *ginext.Context) {
	eventID, _ := strconv.Atoi(c.Param("id"))

	event, err := h.service.GetEvent(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, ginext.H{"error": "event not found"})
		return
	}

	c.JSON(http.StatusOK, dto.EventResponse{
		ID:        event.ID,
		Name:      event.Name,
		Date:      event.Date,
		Capacity:  event.Capacity,
		FreeSeats: event.FreeSeats,
	})
}
