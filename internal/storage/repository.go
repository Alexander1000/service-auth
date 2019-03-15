package storage

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
	credTypeMap map[string]int
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
		credTypeMap: map[string]int{"login": 0, "email": 1, "phone": 2},
	}
}
