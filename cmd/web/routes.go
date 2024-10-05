package main

import (
	"net/http"

	"github.com/mfroeh/lol-metrix/frontend"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServerFS(frontend.Files))

	return mux
}
