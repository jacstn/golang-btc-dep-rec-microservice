package config

import (
	"database/sql"
)

type AppConfig struct {
	Production bool
	DB         *sql.DB
}
