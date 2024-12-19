package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		fmt.Printf("Error opening database")
		return nil, err
	}

	sqlFile, err := os.ReadFile("sql/users.sql")

	if err != nil {
		fmt.Printf("Error reading sql file")
		return nil, err
	}

	_, err = db.Exec(string(sqlFile))

	if err != nil {
		fmt.Printf("Error executing sql file")
		return nil, err
	}

	return db, nil
}
