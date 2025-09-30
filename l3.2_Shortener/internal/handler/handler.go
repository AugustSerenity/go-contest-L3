package handler

import (
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/dto"
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/handler/tools"
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

	router.POST("/shorten", h.ShortenCreate)
	router.GET("/s/:short_url", h.ClickShortLink)
	router.GET("/analytics/:short_url", h.GetAnalytics)

	router.Static("/ui", "./web")
	router.GET("/", func(c *ginext.Context) {
		c.Redirect(302, "/ui/index.html")
	})
	return router
}

func (h *Handler) ShortenCreate(c *ginext.Context) {
	var urlRequest dto.RequestURL
	err := c.BindJSON(&urlRequest)
	if err != nil {
		tools.SendError(c, 400, err.Error())
		return
	}
	urlShorten, err := h.service.Shorten(c.Request.Context(), urlRequest)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to shorten url")
		tools.SendError(c, 500, "failed to shorten url")
		return
	}

	tools.SendSuccess(c, 202, ginext.H{
		"short_url": urlShorten.ShortURL,
	})

}

func (h *Handler) ClickShortLink(c *ginext.Context) {
	shortURL := c.Param("short_url")

	originalUrl, err := h.service.GetOriginalURL(c.Request.Context(), shortURL)
	if err != nil {
		zlog.Logger.Error().Err(err).Str("short_url", shortURL).Msg("Short URL not found")
		tools.SendError(c, 404, "Short URL not found")
		return
	}

	userAgent := c.Request.UserAgent()
	if err := h.service.TrackClick(c.Request.Context(), shortURL, userAgent); err != nil {
		zlog.Logger.Error().Err(err).Str("short_url", shortURL).Msg("Failed to record click")
	}

	c.Redirect(302, originalUrl)
}

func (h *Handler) GetAnalytics(c *ginext.Context) {
	shortURL := c.Param("short_url")

	clicks, err := h.service.GetAnalytics(c.Request.Context(), shortURL)
	if err != nil {
		zlog.Logger.Error().Err(err).Str("short_url", shortURL).Msg("Failed to get analytics")
		tools.SendError(c, 500, "Failed to get analytics")
		return
	}

	tools.SendSuccess(c, 200, clicks)
}
