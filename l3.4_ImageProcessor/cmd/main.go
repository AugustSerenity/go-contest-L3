package main

import (
	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/config"
	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/handler"
	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/service"
	"github.com/AugustSerenity/go-contest-L3/l3.4_ImageProcessor/internal/storage"
	"github.com/wb-go/wbf/zlog"
)

func main() {
	zlog.Init()
	zlog.Logger.Info().Msg("Starting application...")

	var cfg config.Config
	cfgLoader := config.New()
	if err := cfgLoader.Load("config/config.yaml"); err != nil {
		zlog.Logger.Fatal().
			Err(err).
			Msg("Failed to load config file")
	}

	if err := cfgLoader.Unmarshal(&cfg); err != nil {
		zlog.Logger.Fatal().
			Err(err).
			Msg("Failed to unmarshal config")
	}

	db := storage.InitDB(cfg.DB)
	defer storage.CloseDB(db)

	store := storage.New(db)

	srv := service.New(store)

	h := handler.New(srv)
	zlog.Logger.Info().Str("addr", cfg.Server.Address).Msg("starting server")
	if err := h.Router().Run(cfg.Server.Address); err != nil {
		zlog.Logger.Fatal().Err(err).Msg("server failed to start")
	}

}
