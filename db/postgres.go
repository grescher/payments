package db

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"playground/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeebo/errs"
)

var dbErr = errs.Class("database error")

func NewPostgresDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL())
	if err != nil {
		return nil, dbErr.Wrap(err)
	}

	if err = db.Ping(); err != nil {
		return nil, dbErr.Wrap(err)
	}

	if err = createSchema(db); err != nil {
		return nil, dbErr.Wrap(err)
	}
	return db, nil
}

func databaseURL() (url string) {
	url = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DBUser(),
		config.DBPassword(),
		config.DBAddress(),
		config.DBPort(),
		config.DBName(),
	)
	return url
}

func createSchema(db *sql.DB) error {
	log.Println("creating database schema")
	schemaErr := errs.Class("schema error")

	file, err := os.Open(config.DBSchemaPath())
	if err != nil {
		return schemaErr.Wrap(err)
	}
	defer file.Close()

	var query string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		query += scanner.Text() + "\n"
	}
	if err = scanner.Err(); err != nil {
		return schemaErr.Wrap(err)
	}

	if _, err := db.Exec(query); err != nil {
		return schemaErr.Wrap(err)
	}
	return nil
}
