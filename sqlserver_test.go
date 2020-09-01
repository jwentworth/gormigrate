// +build sqlserver

package gormigrate

import (
	_ "gorm.io/gorm/dialects/mssql"
)

func init() {
	databases = append(databases, database{
		name:    "mssql",
		connEnv: "SQLSERVER_CONN_STRING",
	})
}
