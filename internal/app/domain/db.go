package domain

import "database/sql"

type (
	SQLDatabase interface {
		BeginTx() error
		Exec(query string, args ...interface{}) (sql.Result, error)
		Commit() error
		Rollback() error
		ReadQuery(query string) (string, error)
		Ping() error
		QueryRow(query string, args ...interface{}) *sql.Row
	}
)
