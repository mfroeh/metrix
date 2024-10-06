package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/mfroeh/lol-metrix/internal/data"
	"github.com/mfroeh/lol-metrix/internal/lolapi"

	_ "github.com/lib/pq"
)

type config struct {
	port       int
	env        string
	riotAPIKey string
	dsn        string
}

type application struct {
	config config
	lolapi *lolapi.Client
	models data.Models
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.riotAPIKey, "riot-api-key", os.Getenv("RGAPI"), "Riot API key")
	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("METRIX_DB_DSN"), "PostgreSQL DSN")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	db, err := openDB(&cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		lolapi: lolapi.NewClient(cfg.riotAPIKey, "europe", "EUW1"),
		models: data.NewModels(db),
		logger: logger,
	}

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", cfg.port),
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	log.Printf("Starting server on %s", srv.Addr)

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func openDB(cfg *config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
