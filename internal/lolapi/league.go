package lolapi

import (
	"fmt"
	"net/http"

	"github.com/mfroeh/lol-metrix/internal/helpers"
)

type LeagueEntry struct {
	Puuid        string     `json:"puuid"`
	LeagueId     string     `json:"leagueId"`
	SummonerId   string     `json:"summonerId"`
	SummonerName string     `json:"summonerName"`
	QueueType    string     `json:"queueType"`
	RatedTier    string     `json:"ratedTier"`
	Tier         string     `json:"tier"`
	Rank         string     `json:"rank"`
	LeaguePoints int        `json:"leaguePoints"`
	RatedRating  int        `json:"ratedRating"`
	Wins         int        `json:"wins"`
	Losses       int        `json:"losses"`
	HotStreak    bool       `json:"hotStreak"`
	Veteran      bool       `json:"veteran"`
	FreshBlood   bool       `json:"freshBlood"`
	Inactive     bool       `json:"inactive"`
	MiniSeries   MiniSeries `json:"miniSeries"`
}

type MiniSeries struct {
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Progress string `json:"progress"`
}

func (c *Client) GetLeagueBySummonerID(summonerId string) ([]LeagueEntry, error) {
	url := c.makeUrlRegion(fmt.Sprintf("tft/league/v1/entries/by-summoner/%s", summonerId))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.executeRequest(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var leagueEntries []LeagueEntry
		err = helpers.ReadJSON(resp.Body, &leagueEntries)
		if err != nil {
			return nil, err
		}

		return leagueEntries, nil
	default:
		return nil, fmt.Errorf("lolapi: unknown status code: %d", resp.StatusCode)
	}
}
