package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func NewDatabase(connectionString string) (*Database, error) {
	db, error := sql.Open("postgres", connectionString)

	if error != nil {
		return nil, error
	}

	return &Database{db}, nil
}

func (d *Database) Close() {
	d.Close()
}
