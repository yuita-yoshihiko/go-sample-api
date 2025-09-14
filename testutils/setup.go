package testutils

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/joho/godotenv"
)

func LoadFixture(t *testing.T, fixtureDir string) *sql.DB {
	db := loadTestDB()
	fixture, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(fixtureDir),
	)
	if err != nil {
		panic(err)
	}
	if err := fixture.Load(); err != nil {
		panic(err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Log("failed to close database connection: %w", err)
		}
	})
	return db
}

func loadTestDB() *sql.DB {
	if err := godotenv.Load(filepath.Join("../..", ".env-sample")); err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Could not connect to database")
		panic(err.Error())
	}
	return db
}
