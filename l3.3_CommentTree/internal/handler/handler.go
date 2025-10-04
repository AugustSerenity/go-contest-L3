package handler

import "github.com/wb-go/wbf/ginext"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Router() *ginext.Engine {
	router := ginext.New()

	router.POST("/comments", h.CreateComment)

	return router
}

func (h *Handler) CreateComment(c *ginext.Context) {

}
