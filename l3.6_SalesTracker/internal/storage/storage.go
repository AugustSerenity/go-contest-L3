package storage

import (
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.6_SalesTracker/internal/model"
	"github.com/wb-go/wbf/dbpg"
	"github.com/wb-go/wbf/ginext"
)

type Storage struct{ db *dbpg.DB }

func New(db *dbpg.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (st *Storage) SaveItem(c *ginext.Context, item model.Item) (model.Item, error) {
	if item.CreatedAt.IsZero() {
		item.CreatedAt = time.Now()
	}

	query := `
		INSERT INTO items (type, category, amount, data, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	err := st.db.Master.QueryRow(query,
		item.Type,
		item.Category,
		item.Amount,
		item.Date,
		item.CreatedAt,
	).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		return model.Item{}, err
	}

	return item, nil
}
