package lolapi

import (
	"fmt"
	"net/http"
	"strings"
)

type Platform struct {
	apiKey   string
	platform string
}

func NewPlatform(apiKey, platform string) *Platform {
	return &Platform{
		apiKey:   apiKey,
		platform: platform,
	}
}

func (p *Platform) makeUrl(relativeUrl string) string {
	server := fmt.Sprintf("%s.api.riotgames.com", p.platform)
	relativeUrl = strings.TrimPrefix(relativeUrl, "/")
	return fmt.Sprintf("https://%s/%s", server, relativeUrl)
}

func (p *Platform) executeRequest(r *http.Request) (*http.Response, error) {
	r.Header.Add("X-Riot-Token", p.apiKey)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, ErrTooManyRequests
	}

	return resp, nil
}
