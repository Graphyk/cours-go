package database

import (
	"database/sql"

	_ "github.com/go-sql-drive/mysql"
)

type Database struct {
	Read  *sql.DB
	Write *sql.DB
}

type DBConfig struct {
	Read  string `json:"read"`
	Write string `json:"write"`
}

func Open(cfg DBConfig) (*Database, error) {
	db := &Database{}
	var err error

	db.Read, err = OpenDB(cfg.Read)
	if err != nil {
		return nil, err
	}

	db.Write, err = OpenDB(cfg.Read)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
