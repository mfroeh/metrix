package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Summoner struct {
	// tft-summoner-v1
	Puuid         string    `json:"puuid"`
	AccountId     string    `json:"-"`
	ProfileIconId int       `json:"profileIconId"`
	RevisionDate  time.Time `json:"-"`
	SummonerLevel int       `json:"summonerLevel"`
	SummonerId    string    `json:"-"`

	// account-v1
	Name string `json:"name"`
	Tag  string `json:"tag"`

	// Database internal
	Id        int       `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type SummonerModel struct {
	DB *sql.DB
}

func (m *SummonerModel) Insert(summoner *Summoner) (*Summoner, error) {
	// TODO: optimistic locking
	query := `
		INSERT INTO summoners (puuid, account_id, profile_icon_id, revision_date, summoner_level, summoner_id, name, tag)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at;
	`

	args := []any{summoner.Puuid, summoner.AccountId, summoner.ProfileIconId, summoner.RevisionDate, summoner.SummonerLevel, summoner.SummonerId, summoner.Name, summoner.Tag}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, args...)
	err := row.Err()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value"):
			return nil, fmt.Errorf("data: duplicate key value violates unique constraint, %w", err)
		default:
			return nil, err
		}
	}

	err = row.Scan(&summoner.Id, &summoner.CreatedAt, &summoner.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return summoner, nil
}

func (m *SummonerModel) GetByName(name string, tag string) (*Summoner, error) {
	query := `
		SELECT id, created_at, updated_at, puuid, account_id, profile_icon_id, revision_date, summoner_level, summoner_id, name, tag
		FROM summoners
		WHERE name = $1 AND tag = $2;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, name, tag)

	var summoner Summoner
	err := row.Scan(&summoner.Id, &summoner.CreatedAt, &summoner.UpdatedAt, &summoner.Puuid, &summoner.AccountId, &summoner.ProfileIconId, &summoner.RevisionDate, &summoner.SummonerLevel, &summoner.SummonerId, &summoner.Name, &summoner.Tag)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &summoner, nil
}
