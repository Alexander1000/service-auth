package storage

import (
	"database/sql"
)

const (
	AccessTokenStatusActive = 0
	AccessTokenStatusRefreshed = 1
	AccessTokenStatusDisabled = 2

	RefreshTokenStatusActive = 0
	RefreshTokenStatusRefreshed = 1
	RefreshTokenStatusDisabled = 2

	AuthOk = 0
	AuthExpired = 1
	AuthNotFound = 2
	AuthRefreshed = 3
	AuthDisabled = 4
	AuthInternalError = 5
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
