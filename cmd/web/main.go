package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

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

	account, err := app.lolapi.GetAccountByName("Dr Orange", "Caps")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", account)

	startTime := time.Now().AddDate(0, 0, -2)
	matches, err := app.lolapi.GetPlayerMatches(account.Puuid, lolapi.MatchesRequestOptions{
		StartTime: &startTime,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", matches)

	match, err := app.lolapi.GetMatch(matches[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", match)

	json, err := json.Marshal(match)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", len(json))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	log.Printf("Starting server on %s", srv.Addr)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
