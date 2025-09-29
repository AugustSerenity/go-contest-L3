package handler

import (
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/dto"
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/handler/tools"
	"github.com/wb-go/wbf/ginext"
)

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

func (h *Handler) ShortenCreate(c *ginext.Context) {
	var urlRequest dto.RequestURL
	err := c.BindJSON(&urlRequest)
	if err != nil {
		tools.SendError(c, 400, err.Error())
		return
	}

}

func (h *Handler) ClickShortLink(c *ginext.Context) {}
func (h *Handler) Getanalytics(c *ginext.Context)   {}
