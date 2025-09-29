package main

import (
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/config"
	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/storage"
	"github.com/wb-go/wbf/zlog"
)

func main() {
	var cfg config.Config

	zlog.Init()

	loader := config.New()
	if err := loader.Load("config.yml"); err != nil {
		zlog.Logger.Fatal().
			Err(err).
			Msg("Failed to load config file")
	}

	if err := loader.Unmarshal(&cfg); err != nil {
		zlog.Logger.Fatal().
			Err(err).
			Msg("Failed to unmarshal config")
	}

	db := storage.InitDB(cfg.DB)
	defer storage.CloseDB(db)

}
