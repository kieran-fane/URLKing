package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"backend/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/rs/xid"
)

type App struct {
	DB *pgxpool.Pool
}

func main() {
	pool := db.Connect()
	defer pool.Close()

	app := &App{DB: pool}
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
	}))

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

	_ = db.InsertURL(app.DB, slug, req.URL)
	return c.JSON(http.StatusOK, ShortenResponse{Slug: slug, URL: req.URL})
}

func (app *App) getURL(c echo.Context) error {
	slug := c.Param("hashslug")

	originalURL, err := db.Lookup(app.DB, slug)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "slug not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"url": originalURL})
}

