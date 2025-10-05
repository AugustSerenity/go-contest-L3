package storages

import (
	"fmt"

	"github.com/AugustSerenity/go-contest-L3/tree/main/l3.3_CommentTree/internal/model"
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

func (st *Storage) InsertComment(comment model.Comment) (int64, error) {
	query := `
		INSERT INTO comments (text, parent_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	var id int64

	err := st.db.Master.QueryRow(query,
		comment.Text,
		comment.ParentID,
		comment.CreatedAt,
		comment.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("insert comment failed: %w", err)
	}

	return id, nil
}
