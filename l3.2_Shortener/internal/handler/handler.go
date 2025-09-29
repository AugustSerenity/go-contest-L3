package handler

import "github.com/wb-go/wbf/ginext"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Router() *ginext.Engine {
	router := ginext.New()

	router.POST("/shorten", h.ShortenCreate)
	router.GET("/s/:short_url", h.ClickShortLink)
	router.GET("/analytics/:short_url", h.Getanalytics)
	return router
}

func (h *Handler) ShortenCreate(c *ginext.Context) {}

func (h *Handler) ClickShortLink(c *ginext.Context) {}
func (h *Handler) Getanalytics(c *ginext.Context)   {}
