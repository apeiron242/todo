package db

import (
	"database/sql"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	rootDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path.Join(rootDir, ".todoData.db"))
	if err != nil {
		file, _ = os.Create(path.Join(rootDir, ".todoData.db"))
	}

	db, err := sql.Open("sqlite3", file.Name())
	if err != nil {
		return nil, err
	}

	createTableQuery := `
					CREATE TABLE IF NOT EXISTS data (
						id integer PRIMARY KEY AUTOINCREMENT,
						title text,
						details text,
						time text
					)
	`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}
