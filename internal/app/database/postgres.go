package database

import (
	"database/sql"
	"errors"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

const errorConnectionMsg = "error, connection with postgres not found"

var (
	once sync.Once

	Client *sql.DB
)

func init() {
	once.Do(func() {
		Client = buildDBClient()
	})
}

func buildDBClient() *sql.DB {
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/aug7?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed connection with DB, it is not available: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed connection with Ping to DB: %v", err)
	}
	log.Print("Connected to postgres")
	return db
}

// Postgres model
type Postgres struct {
	db *sql.DB
	tx *sql.Tx
}

// BeginTx start transaction
func (p *Postgres) BeginTx() error {
	var err error

	p.tx, err = p.db.Begin()
	if err != nil {
		return err
	}
	return nil
}

// Exec queries
func (p *Postgres) Exec(query string, args ...interface{}) (sql.Result, error) {
	if p.tx == nil {
		return nil, errors.New(errorConnectionMsg)
	}
	result, err := p.tx.Exec(query, args...)
	return result, err
}

// Commit ....
func (p *Postgres) Commit() error {
	if p.tx == nil {
		return errors.New(errorConnectionMsg)
	}
	err := p.tx.Commit()
	return err
}

// Rollback ....
func (p *Postgres) Rollback() error {
	if p.tx == nil {
		return errors.New(errorConnectionMsg)
	}
	err := p.tx.Rollback()
	return err
}

// ReadQuery utils.ReadQuery
func (p *Postgres) ReadQuery(query string) (string, error) {
	// TODO: Review the implemtation of this
	// file, err := utils.ReadQuery(query)
	return "", nil
}

// Ping pings
func (p *Postgres) Ping() error {
	return p.db.Ping()
}

// QueryRow rows
func (p *Postgres) QueryRow(query string, args ...interface{}) *sql.Row {
	return p.tx.QueryRow(query, args...)
}
