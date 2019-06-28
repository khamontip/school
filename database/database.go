package database

import (
	"database/sql"
	"os"

	//"os"

	_ "github.com/lib/pq"
)

func ConnDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	// db, err := sql.Open("postgres", "postgres://dbuvonnf:R2v1DtbVijedmzNCTHnIWAWqtE_FYw8_@satao.db.elephantsql.com:5432/dbuvonnf")
	if err != nil {
		return nil, err
	}
	return db, nil

}
