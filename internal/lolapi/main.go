package lolapi

import (
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	apiKey   string
	platform string
	region   string
}

func NewClient(apiKey, platform, region string) *Client {
	return &Client{
		apiKey:   apiKey,
		platform: platform,
		region:   region,
	}
}

func (p *Client) makeUrlPlatform(relativeUrl string) string {
	server := fmt.Sprintf("%s.api.riotgames.com", p.platform)
	relativeUrl = strings.TrimPrefix(relativeUrl, "/")
	return fmt.Sprintf("https://%s/%s", server, relativeUrl)
}

func (p *Client) makeUrlRegion(relativeUrl string) string {
	server := fmt.Sprintf("%s.api.riotgames.com", p.region)
	relativeUrl = strings.TrimPrefix(relativeUrl, "/")
	return fmt.Sprintf("https://%s/%s", server, relativeUrl)
}

func (p *Client) executeRequest(r *http.Request) (*http.Response, error) {
	r.Header.Add("X-Riot-Token", p.apiKey)

	fmt.Printf("Executing request to %s, body: %s\n", r.URL.String(), r.Body)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	if err, ok := ErrorForStatusCode[resp.StatusCode]; ok {
		return nil, err
	}

	return resp, nil
}
