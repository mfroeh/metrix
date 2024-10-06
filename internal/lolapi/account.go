package lolapi

import (
	"fmt"
	"net/http"

	"github.com/mfroeh/lol-metrix/internal/helpers"
)

type RiotAccount struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func (p *Client) GetAccountByName(gameName, tagLine string) (*RiotAccount, error) {
	url := p.makeUrlPlatform(fmt.Sprintf("/riot/account/v1/accounts/by-riot-id/%s/%s", gameName, tagLine))
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
		var account RiotAccount
		err = helpers.ReadJSON(resp.Body, &account)
		if err != nil {
			return nil, err
		}
		return &account, nil
	case http.StatusNotFound:
		return nil, ErrResourceNotFound
	default:
		return nil, fmt.Errorf("lolapi: unknown response status code: %d", resp.StatusCode)
	}
}

func (p *Client) GetAccountByPuuid(puuid string) (*RiotAccount, error) {
	url := p.makeUrlPlatform(fmt.Sprintf("/riot/account/v1/accounts/by-puuid/%s", puuid))
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
		var account RiotAccount
		err = helpers.ReadJSON(resp.Body, &account)
		if err != nil {
			return nil, err
		}
		return &account, nil
	default:
		return nil, fmt.Errorf("lolapi: unknown response status code: %d", resp.StatusCode)
	}
}
