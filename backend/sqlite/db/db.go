package db

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./shortener.db")
	if err != nil {
		log.Fatal("Failed to connect to SQLite:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS urls (
		slug TEXT UNIQUE NOT NULL,
		long_url TEXT UNIQUE NOT NULL
	);`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Connected to SQLite DB.")
	return db
}

func InsertURL(db *sql.DB, slug string, url string) error {
	_, err := db.Exec(`INSERT INTO urls (slug, long_url) VALUES (?, ?)`, slug, url)
	return err
}

func Lookup(db *sql.DB, slug string) (string, error) {
	var url string
	err := db.QueryRow(`SELECT long_url FROM urls WHERE slug = ?`, slug).Scan(&url)
	if err != nil {
		return "", errors.New("not found")
	}
	return url, nil
}

