package database

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/lib/pq"

	"github.com/Alexander1000/service-auth/internal/utils"
)

const (
	DBMaxOpenConns = 8
	DBMaxIdleConns = 2
)

// Settings содержит настройки подключения
type Settings struct {
	Host     string             `json:"host"`
	Port     int                `json:"port"`
	User     string             `json:"user"`
	Password utils.SecretString `json:"password"`
	Database string             `json:"database"`
}

// Connect осуществляет подключение
func Connect(setting Settings) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn(setting))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	db.SetMaxOpenConns(DBMaxOpenConns)
	db.SetMaxIdleConns(DBMaxIdleConns)

	return db, err
}

func dsn(settings Settings) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		settings.User,
		url.PathEscape(settings.Password.Value()),
		settings.Host,
		settings.Port,
		settings.Database,
	)
}
