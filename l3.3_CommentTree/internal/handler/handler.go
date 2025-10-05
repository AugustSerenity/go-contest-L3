package handler

import (
	"github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/handler/dto"
	"github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/handler/tools"
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

	router.POST("/comments", h.CreateComment)

	return router
}

func (h *Handler) CreateComment(c *ginext.Context) {
	var commentIn dto.CommentRequest
	err := c.BindJSON(&commentIn)
	if err != nil {
		tools.SendError(c, 400, err.Error())
		return
	}

	comment, err := h.service.CreateComment(commentIn)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to create cooment")
		tools.SendError(c, 500, "failed to create cooment")
		return
	}

	tools.SendSuccess(c, 202, ginext.H{
		"comment": comment,
	})

}
