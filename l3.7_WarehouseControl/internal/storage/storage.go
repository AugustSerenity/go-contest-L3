package storage

import (
	"github.com/AugustSerenity/go-contest-L3/l3.7_WarehouseControl/internal/model"
	"github.com/wb-go/wbf/dbpg"
)

type Storage struct {
	db *dbpg.DB
}

func New(db *dbpg.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) ListItems() ([]model.Item, error) {
	rows, err := s.db.Master.Query(`SELECT id, name, quantity, updated_at FROM items ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []model.Item
	for rows.Next() {
		var it model.Item
		if err := rows.Scan(&it.ID, &it.Name, &it.Quantity, &it.UpdatedAt); err != nil {
			return nil, err
		}
		res = append(res, it)
	}
	return res, nil
}

func (s *Storage) CreateItem(name string, qty int, username string) error {
	s.db.Master.Exec("SET app.user = '" + username + "'")
	_, err := s.db.Master.Exec(`INSERT INTO items(name, quantity) VALUES($1,$2)`, name, qty)
	return err
}
