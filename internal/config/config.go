package config

import (
	"github.com/Alexander1000/service-auth/internal/database"
)

type Config struct {
	Port int `json:"port"`
	Database database.Settings `json:"database"`
}
