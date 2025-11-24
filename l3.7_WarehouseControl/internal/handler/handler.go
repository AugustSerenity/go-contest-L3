package handler

import "github.com/wb-go/wbf/ginext"

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

	router.Static("/static", "./web")
	router.GET("/", func(c *ginext.Context) {
		c.File("./web/index.html")
	})

	return router
}

func (h *Handler) CreateItem(c *ginext.Context) {
	
}
