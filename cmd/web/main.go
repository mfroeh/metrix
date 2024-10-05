package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mfroeh/lol-metrix/internal/lolapi"
)

type config struct {
	port       int
	env        string
	riotAPIKey string
}

type application struct {
	config config
	lolapi *lolapi.Platform
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.riotAPIKey, "riot-api-key", "", "Riot API key")

	flag.Parse()

	app := &application{
		config: cfg,
		lolapi: lolapi.NewPlatform(cfg.riotAPIKey, "europe"),
	}

	account, err := app.lolapi.GetAccount("Dr Orange", "Caps")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", account)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	log.Printf("Starting server on %s", srv.Addr)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
