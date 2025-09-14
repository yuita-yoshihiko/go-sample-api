package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/lib/pq"
)

var ErrNotFound = errors.New("record not found")

func Init() (*sql.DB, error) {
	connection, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}

	db, err := sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(300 * time.Second)
	return db, nil
}

type DBUtils interface {
	Error(error) error

	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type dbUtil struct {
	*sql.DB
}

func NewDBUtil(db *sql.DB) DBUtils {
	return &dbUtil{DB: db}
}

func (u *dbUtil) Error(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w: %v", ErrNotFound, err)
	}
	return err
}
