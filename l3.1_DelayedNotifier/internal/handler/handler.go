package handler

import (
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/handler/dto"
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/handler/tools"
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
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

	router.POST("/notify", h.Notify)
	router.GET("/notify/:id", h.NotifyGetID)
	router.DELETE("/notify/:id", h.Delete)

	return router
}

func (h *Handler) Notify(c *ginext.Context) {
	var notifyRequest dto.NotificationRequest
	err := c.BindJSON(&notifyRequest)
	if err != nil {
		tools.SendError(c, 400, err.Error())
		return
	}

	notif := model.CastToNotification(notifyRequest)

	err = h.service.CreateNotification(c, notif)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to create notification")
		tools.SendError(c, 503, "failed to schedule notification")
		return
	}

	tools.SendSuccess(c, 202, ginext.H{
		"id":     notif.ID,
		"status": notif.Status,
	})
}

func (h *Handler) NotifyGetID(c *ginext.Context) {}

func (h *Handler) Delete(c *ginext.Context) {}
