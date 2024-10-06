package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mfroeh/metrix/frontend"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServerFS(frontend.Files)

	// Handle other routes by serving the index.html file
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fs.ServeHTTP(w, r)
			return
		}

		f, err := frontend.Files.Open(strings.TrimPrefix(r.URL.Path, "/"))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				r.URL.Path = "/"
			}
		} else {
			defer f.Close()
		}

		fmt.Println("Serving ", r.URL.Path)
		fs.ServeHTTP(w, r)
	})

	mux.HandleFunc("POST /api/v1/summoner", app.createOrGetSummoner)

	return app.logRequest(mux)
}
