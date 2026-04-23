package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	conn *sql.DB
}

func NewDatabase(connectionString string) (*Database, error) {
	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	db := &Database{
		conn: conn,
	}

	return db, nil
}

func (d *Database) Close() {
	d.Close()
}

func (d *Database) Ping() error {
	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}
