package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Conn *sql.DB

func Connect() error {
	conn, err := sql.Open("sqlite3", "file:test.db")

	if err != nil {
		return err
	}

	err = conn.Ping()

	if err != nil {
		return err
	}

	Conn = conn

	return initDB()
}

func initDB() error {
	query := `
		create table if not exists connections(
			id text not null,
			address text not null,
			username text not null,
			password text not null
		);
	`

	_, err := Conn.Exec(query, nil)

	return err
}
