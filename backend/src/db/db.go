package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DbConn *sql.DB

func InitDB(dataSourceName string) (*sql.DB, error) {
	var err error
	DbConn, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = DbConn.Ping(); err != nil {
		return nil, err
	}
	return DbConn, nil
}
