// +build mysql

package gormigrate

import (
	"os"

	"gorm.io/driver/mysql"
)


func init() {
	databases = append(databases, database{
		name: "mysql",
		db:   mysql.Open(os.Getenv("MYSQL_CONN_STRING")),
	})
}
