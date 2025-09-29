package storage

import "github.com/wb-go/wbf/dbpg"

type Storage struct {
	db *dbpg.DB
}

func New(db *dbpg.DB) *Storage {
	return &Storage{
		db: db,
	}
}
