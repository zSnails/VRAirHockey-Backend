package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

// Initializes the \[DB]
func Init() error {
	var err error
	db, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		return err
	}
	return nil
}

// Returns the [DB] handle.
func Get() *sql.DB {
	return db
}

// Close closes the database and prevents new queries from starting. Close then
// waits for all queries that have started processing on the server to finish.
// It is rare to Close a \[DB], as the \[DB] handle is meant to be long-lived
// and shared between many goroutines.
func Close() error {
	return db.Close()
}
