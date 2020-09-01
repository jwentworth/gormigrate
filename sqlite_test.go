// +build sqlite

package gormigrate

import (
	"os"

	"gorm.io/driver/sqlite"
)

func init() {
	databases = append(databases, database{
		name: "sqlite3",
		db:   sqlite.Open(os.Getenv("SQLITE_CONN_STRING")),
	})
}
