package repository

import (
	"context"
	"database/sql"

	"go-wsl-project/internal/core/domain"
)

type PostgresItemRepository struct {
	db *sql.DB
}

func NewPostgresItemRepository(db *sql.DB) *PostgresItemRepository {
	return &PostgresItemRepository{db: db}
}

func (r *PostgresItemRepository) List(ctx context.Context) ([]domain.Item, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name FROM items ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		var it domain.Item
		if err := rows.Scan(&it.ID, &it.Name); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, rows.Err()
}

func (r *PostgresItemRepository) Create(ctx context.Context, name string) (domain.Item, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, "INSERT INTO items (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		return domain.Item{}, err
	}
	return domain.Item{ID: id, Name: name}, nil
}
