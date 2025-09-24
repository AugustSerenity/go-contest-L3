package handler

import "github.com/wb-go/wbf/ginext"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Router() *ginext.Engine {
	router := ginext.New()

	router.POST("/notify", h.Notify)
	router.GET("/notify/:id", h.NotifyGetID)
	router.DELETE("/notify/:id", h.Delete)

	return router
}

func (h *Handler) Notify(c *ginext.Context) {

}

func (h *Handler) NotifyGetID(c *ginext.Context) {}

func (h *Handler) Delete(c *ginext.Context) {}
