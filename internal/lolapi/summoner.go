package lolapi

import (
	"fmt"
	"net/http"

	"github.com/mfroeh/metrix/internal/helpers"
)

type Summoner struct {
	AccountId     string `json:"accountId"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	Id            string `json:"id"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
}

func (p *Client) GetSummonerByPuuid(puuid string) (*Summoner, error) {
	url := p.makeUrlRegion(fmt.Sprintf("/tft/summoner/v1/summoners/by-puuid/%s", puuid))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.executeRequest(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var summoner Summoner
		err = helpers.ReadJSON(resp.Body, &summoner)
		if err != nil {
			return nil, err
		}
		return &summoner, nil
	default:
		return nil, fmt.Errorf("lolapi: unknown response status code: %d", resp.StatusCode)
	}
}
