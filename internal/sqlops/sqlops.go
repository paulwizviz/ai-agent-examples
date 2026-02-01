package sqlops

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrDBConn = errors.New("connection error")
)

var ()

// NewSQLiteMem instantiate a connection to SQLite
func NewSQLiteMem() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConn, err)
	}
	return db, nil
}

// NewSQLiteFile instantiate a file based SQLite
func NewSQLiteFile(f string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", f)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConn, err)
	}
	return db, nil
}
