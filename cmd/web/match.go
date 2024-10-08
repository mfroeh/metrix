package main

import (
	"net/http"
	"time"

	"github.com/mfroeh/metrix/internal/data"
	"github.com/mfroeh/metrix/internal/helpers"
	"github.com/mfroeh/metrix/internal/lolapi"
)

func (app *application) getMatches(w http.ResponseWriter, r *http.Request) {
	input := struct {
		Puuid string    `json:"puuid"`
		After time.Time `json:"after"`
	}{}

	err := helpers.ReadJSON(r.Body, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	count := 200
	matches, err := app.lolapi.GetPlayerMatches(input.Puuid, lolapi.MatchesRequestOptions{
		StartTime: &input.After,
		Start:     &count,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	for _, match := range matches {
		go func() {
			apiMatch, err := app.lolapi.GetMatch(match)
			if err != nil {
				app.logError(r, err)
				return
			}

			_, err = app.models.Matches.Insert(data.MatchFromApiMatch(apiMatch))
			if err != nil {
				app.logError(r, err)
				return
			}
		}()
	}

}
