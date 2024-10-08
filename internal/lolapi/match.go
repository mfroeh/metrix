package lolapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mfroeh/metrix/internal/helpers"
)

type Match struct {
	Info     MatchInfo     `json:"info"`
	Metadata MatchMetadata `json:"metadata"`
}

type MatchMetadata struct {
	DataVersion  string   `json:"data_version"`
	MatchId      string   `json:"match_id"`
	Participants []string `json:"participants"`
}

type MatchInfo struct {
	Participants []Participant `json:"participants"`

	GameDatetime    int64   `json:"game_datetime"`
	GameCreation    int64   `json:"gameCreation"`
	GameLength      float64 `json:"game_length"`
	GameVariation   string  `json:"game_variation"`
	GameVersion     string  `json:"game_version"`
	QueueId         int     `json:"queue_id"`
	QueueIdOld      int     `json:"queueId"`
	TftSetNumber    int     `json:"tft_set_number"`
	TftGameType     string  `json:"tft_game_type"`
	TftSetCoreName  string  `json:"tft_set_core_name"`
	EndOfGameResult string  `json:"endOfGameResult"`
	MapId           int     `json:"mapId"`
	GameId          int64   `json:"gameId"`
}

type Participant struct {
	Companion Companion `json:"companion"`
	Traits    []Trait   `json:"traits"`
	Units     []Unit    `json:"units"`
	Augments  []string  `json:"augments"`

	GoldLeft             int     `json:"gold_left"`
	LastRound            int     `json:"last_round"`
	Level                int     `json:"level"`
	Placement            int     `json:"placement"`
	PlayerEliminated     int     `json:"players_eliminated"`
	Puuid                string  `json:"puuid"`
	TimeEliminated       float64 `json:"time_eliminated"`
	TotalDamageToPlayers int     `json:"total_damage_to_players"`

	Missions map[string]any `json:"missions"`
}

type Companion struct {
	ContentId string `json:"content_id"`
	ItemId    int    `json:"item_id"`
	SkinId    int    `json:"skin_id"`
	Species   string `json:"species"`
}

type Trait struct {
	Name        string `json:"name"`
	NumUnits    int    `json:"num_units"`
	Style       int    `json:"style"`
	TierCurrent int    `json:"tier_current"`
	TierTotal   int    `json:"tier_total"`
}

type Unit struct {
	CharacterId string   `json:"character_id"`
	ItemNames   []string `json:"itemNames"`
	Chosen      string   `json:"chosen"`
	Name        string   `json:"name"`
	Rarity      int      `json:"rarity"`
	Tier        int      `json:"tier"`
}

type MatchesRequestOptions struct {
	Start     *int
	StartTime *time.Time
	EndTime   *time.Time
	Count     *int
}

func (p *Client) GetPlayerMatches(puuid string, options MatchesRequestOptions) ([]string, error) {
	url := p.makeUrlRegion(fmt.Sprintf("/tft/match/v1/matches/by-puuid/%s/ids", puuid))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	if options.StartTime != nil {
		q.Set("startTime", fmt.Sprintf("%d", options.StartTime.Unix()))
	}
	if options.EndTime != nil {
		q.Set("endTime", fmt.Sprintf("%d", options.EndTime.Unix()))
	}
	if options.Start != nil {
		q.Set("start", fmt.Sprintf("%d", *options.Start))
	}
	if options.Count != nil {
		q.Set("count", fmt.Sprintf("%d", *options.Count))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := p.executeRequest(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var matches []string
		err = helpers.ReadJSON(resp.Body, &matches)
		if err != nil {
			return nil, err
		}
		return matches, nil
	default:
		return nil, fmt.Errorf("lolapi: unknown response status code: %d", resp.StatusCode)
	}
}

func (p *Client) GetMatch(matchId string) (*Match, error) {
	url := p.makeUrlPlatform(fmt.Sprintf("/tft/match/v1/matches/%s", matchId))
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
		var match Match
		err = helpers.ReadJSON(resp.Body, &match)
		if err != nil {
			return nil, err
		}
		return &match, nil
	default:
		return nil, fmt.Errorf("lolapi: unknown response status code: %d", resp.StatusCode)
	}
}
