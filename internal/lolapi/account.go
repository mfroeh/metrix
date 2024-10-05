package lolapi

import (
	"fmt"
	"net/http"
)

type RiotAccount struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func (p *Platform) GetAccount(gameName, tagLine string) (*RiotAccount, error) {
	url := p.makeUrl(fmt.Sprintf("/riot/account/v1/accounts/by-riot-id/%s/%s", gameName, tagLine))
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
		err = readJSON(resp.Body, &account)
		if err != nil {
			return nil, err
		}
		return &account, nil
	case http.StatusNotFound:
		return nil, fmt.Errorf("lolapi: account not found for %s#%s", gameName, tagLine)
	default:
		return nil, fmt.Errorf("lolapi: unknown error: %d", resp.StatusCode)
	}
}
