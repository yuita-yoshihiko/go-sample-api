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

type txKeyType struct{}

var txKey = txKeyType{}

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

type DBManager interface {
	DoInTx(ctx context.Context, f func(context.Context) (any, error)) (any, error)
}

type DBManagerImpl struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) DBManager {
	return &DBManagerImpl{db: db}
}

func (m *DBManagerImpl) DoInTx(ctx context.Context, f func(context.Context) (any, error)) (any, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	ctx = context.WithValue(ctx, txKey, tx)
	v, err := f(ctx)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return v, nil
}

type DBUtils interface {
	Error(error) error

	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
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
