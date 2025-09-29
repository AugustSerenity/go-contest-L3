package storage

import (
	"context"
	"database/sql"

	"github.com/AugustSerenity/go-contest-L3/l3.2/internal/model"
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

func (st *Storage) SaveLink(ctx context.Context, url *model.URL) error {
	query := `
		INSERT INTO links (original_url, short_url, created_at)
		VALUES ($1, $2, $3)
	`
	_, err := st.db.ExecContext(ctx, query, url.OriginalURL, url.ShortURL, url.CreateAt)
	return err
}

func (st *Storage) ExistsByShortCode(ctx context.Context, shortCode string) (bool, error) {
	query := `
		SELECT 1 FROM links WHERE short_url = $1 LIMIT 1
	`

	row := st.db.Master.QueryRowContext(ctx, query, shortCode)

	var exists int
	err := row.Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
