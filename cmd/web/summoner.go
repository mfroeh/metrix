package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mfroeh/metrix/internal/data"
	"github.com/mfroeh/metrix/internal/helpers"
	"github.com/mfroeh/metrix/internal/lolapi"
	"github.com/mfroeh/metrix/internal/validator"
)

func (app *application) createOrGetSummoner(w http.ResponseWriter, r *http.Request) {
	nameAndTag := struct {
		Name string `json:"name"`
		Tag  string `json:"tag"`
	}{}

	err := helpers.ReadJSON(r.Body, &nameAndTag)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(nameAndTag.Name != "", "name", "must be provided")
	v.Check(nameAndTag.Tag != "", "tag", "must be provided")
	v.Check(validator.InRange(len(nameAndTag.Name), 3, 16), "name", "must be 3-16 characters long")
	v.Check(validator.InRange(len(nameAndTag.Tag), 3, 5), "tag", "must be 3-5 characters long")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// try to get summoner from db, otherwise kick off go routine to get from Riot API
	summoner, err := app.models.Summoners.GetByName(nameAndTag.Name, nameAndTag.Tag)
	if nil == err {
		// app.writeJSON(w, http.StatusOK, envelope{"summoner": summoner}, nil)
		// return
	}

	if !errors.Is(err, data.ErrRecordNotFound) {
		// app.serverErrorResponse(w, r, err)
		// return
	}

	fmt.Println("summoner not found, getting from Riot API")

	fmt.Println("getting account from Riot API")
	account, err := app.lolapi.GetAccountByName(nameAndTag.Name, nameAndTag.Tag)
	if err != nil {
		switch {
		case errors.Is(err, lolapi.ErrResourceNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	fmt.Println("getting summoner from Riot API")
	apiSummoner, err := app.lolapi.GetSummonerByPuuid(account.Puuid)
	if err != nil {
		switch {
		case errors.Is(err, lolapi.ErrResourceNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	fmt.Println("getting league from Riot API")
	apiLeagues, err := app.lolapi.GetLeagueBySummonerID(apiSummoner.Id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	leagues := make([]*data.League, 0)
	for _, league := range apiLeagues {
		leagues = append(leagues, &data.League{
			Puuid:        apiSummoner.Puuid,
			SummonerId:   apiSummoner.Id,
			QueueType:    league.QueueType,
			Tier:         league.Tier,
			Rank:         helpers.Rntoi(league.Rank),
			Wins:         league.Wins,
			Losses:       league.Losses,
			HotStreak:    league.HotStreak,
			Veteran:      league.Veteran,
			FreshBlood:   league.FreshBlood,
			Inactive:     league.Inactive,
			LeaguePoints: league.LeaguePoints,
			RatedRating:  league.RatedRating,
			MiniSeries: data.MiniSeries{
				Wins:     league.MiniSeries.Wins,
				Losses:   league.MiniSeries.Losses,
				Target:   league.MiniSeries.Target,
				Progress: league.MiniSeries.Progress,
			},
		})
	}

	summoner = &data.Summoner{
		Name:          account.GameName,
		Tag:           account.TagLine,
		Puuid:         account.Puuid,
		AccountId:     apiSummoner.AccountId,
		ProfileIconId: apiSummoner.ProfileIconId,
		RevisionDate:  time.UnixMilli(apiSummoner.RevisionDate),
		SummonerLevel: apiSummoner.SummonerLevel,
		SummonerId:    apiSummoner.Id,
		Leagues:       leagues,
	}

	fmt.Println("inserting summoner into db")
	summoner, err = app.models.Summoners.Insert(summoner)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	fmt.Println("writing summoner to client")
	app.writeJSON(w, http.StatusOK, envelope{"summoner": summoner}, nil)
}
