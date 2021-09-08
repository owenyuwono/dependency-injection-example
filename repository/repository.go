package repository

import (
	"database/sql"
	"dependency-injection-example/model"
	"errors"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Insert(data model.Data) (string, error) {
	id := uuid.New().String()
	res, err := r.db.Exec(`
	INSERT INTO testtable (id, name) VALUES ($1, $2)
	`, id, data.Name)
	if err != nil {
		return "", err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if affected < 1 {
		return "", errors.New("no rows affected")
	}
	return id, nil
}
