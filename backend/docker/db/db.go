// Database connection made here.
package db

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {
	connStr := "postgres://devuser:devpass@localhost:5432/url_shortener"
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}
	log.Print("Connected to db.")
	return pool
}


func InsertURL(pool *pgxpool.Pool, slug string, url string) error {
	_, err := pool.Exec(context.Background(),
		`INSERT INTO urls (slug, long_url) VALUES ($1, $2)`, slug, url)
	return err
}

func Lookup(pool *pgxpool.Pool, slug string) (string, error) {
	var url string
	err := pool.QueryRow(context.Background(),
		`SELECT long_url FROM urls WHERE slug = $1`, slug).Scan(&url)
	if err != nil {
		return "NOT", errors.New("not found")
	}
	return url, nil
}
