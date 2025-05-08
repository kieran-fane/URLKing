package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"backend/db"
	"database/sql"
	// "github.com/mattn/go-sqlite3"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)
type App struct {
	DB *sql.DB
}

func main() {
	pool := db.Connect()
	defer pool.Close()

	app := &App{DB: pool}
	e := echo.New()

	e.POST("/shorten", app.shortenURL)
	e.GET("/:hashslug", app.getURL)

	e.Logger.Fatal(e.Start(":8080"))
}

func (app *App) shortenURL(c echo.Context) error {
	type ShortenRequest struct {
		URL string `json:"url"`
	}
	type ShortenResponse struct {
		Slug string `json:"hashslug"`
		URL  string `json:"url"`
	}

	var req ShortenRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
	}
	if req.URL == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing url field"})
	}

	hash := sha256.Sum256([]byte(req.URL))
	slug := hex.EncodeToString(hash[:])[:8]

	err := db.InsertURL(app.DB, slug, req.URL)
	if err != nil {
		slug = xid.New().String()
		err = db.InsertURL(app.DB, slug, req.URL)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				map[string]string{"error": "inserting into db failed"})
		}
	}

	return c.JSON(http.StatusOK, ShortenResponse{Slug: slug, URL: req.URL})
}

func (app *App) getURL(c echo.Context) error {
	slug := c.Param("hashslug")

	originalURL, err := db.Lookup(app.DB, slug)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "slug not found"})
	}

	return c.Redirect(http.StatusFound, originalURL)
}

