package storage

import (
	"strconv"
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
		INSERT INTO items (type, category, amount, date, created_at)
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

func (st *Storage) AnalyticsCalculate(c *ginext.Context, filter model.ItemsFilter) (model.AnalyticsResponse, error) {
	query := `
        SELECT 
            COALESCE(SUM(amount), 0) as total_sum,
            COALESCE(AVG(amount), 0) as average,
            COUNT(amount) as count,
            COALESCE(PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY amount), 0) AS median,
            COALESCE(PERCENTILE_CONT(0.9) WITHIN GROUP (ORDER BY amount), 0) AS percentile_90
        FROM items 
        WHERE 1=1
    `

	args := []any{}
	argCounter := 1

	if filter.From != nil {
		query += " AND date >= $" + strconv.Itoa(argCounter)
		args = append(args, *filter.From)
		argCounter++
	}
	if filter.To != nil {
		query += " AND date <= $" + strconv.Itoa(argCounter)
		args = append(args, *filter.To)
		argCounter++
	}
	if filter.Category != nil && *filter.Category != "" {
		query += " AND category = $" + strconv.Itoa(argCounter)
		args = append(args, *filter.Category)
		argCounter++
	}

	var response model.AnalyticsResponse
	err := st.db.Master.QueryRow(query, args...).Scan(
		&response.Sum,
		&response.Avg,
		&response.Count,
		&response.Median,
		&response.P90,
	)
	if err != nil {
		return model.AnalyticsResponse{}, err
	}

	return response, nil
}
