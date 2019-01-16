package db

import (
	"database/sql"
)

// Store describes set of functionalities needed to interact with database
type Store interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(string, ...interface{}) (sql.Result, error)
	Begin() (*sql.Tx, error)
	Close() error
}
