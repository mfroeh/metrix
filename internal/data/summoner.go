package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Summoner struct {
	// tft-summoner-v1
	Puuid         string    `json:"puuid"`
	AccountId     string    `json:"accountId"`
	ProfileIconId int       `json:"profileIconId"`
	RevisionDate  time.Time `json:"revisionDate"`
	SummonerLevel int       `json:"summonerLevel"`
	SummonerId    string    `json:"summonerId"`

	// account-v1
	Name string `json:"name"`
	Tag  string `json:"tag"`

	// tft-league-v1 (stored in league table)
	Leagues []*League `json:"leagues"`

	// Database internal
	Id        int       `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type League struct {
	QueueType    string     `json:"queueType"`
	Tier         string     `json:"tier"`
	Rank         int        `json:"rank"`
	Wins         int        `json:"wins"`
	Losses       int        `json:"losses"`
	LeaguePoints int        `json:"leaguePoints"`
	RatedRating  int        `json:"ratedRating"`
	HotStreak    bool       `json:"hotStreak"`
	Veteran      bool       `json:"veteran"`
	FreshBlood   bool       `json:"freshBlood"`
	Inactive     bool       `json:"inactive"`
	MiniSeries   MiniSeries `json:"miniSeries"`

	Puuid      string `json:"puuid"`
	SummonerId string `json:"summonerId"`

	// Database internal
	Id        int       `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type MiniSeries struct {
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Progress string `json:"progress"`
}

type SummonerModel struct {
	DB *sql.DB
}

func (m *SummonerModel) Insert(summoner *Summoner) (*Summoner, error) {
	// TODO: optimistic locking
	query := `
		INSERT INTO summoners (puuid, account_id, profile_icon_id, revision_date, summoner_level, summoner_id, name, tag)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (puuid)
		DO UPDATE SET
			profile_icon_id = EXCLUDED.profile_icon_id,
			revision_date = EXCLUDED.revision_date,
			summoner_level = EXCLUDED.summoner_level,
			name = EXCLUDED.name,
			tag = EXCLUDED.tag,
			updated_at = now()
		RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	args := []any{summoner.Puuid, summoner.AccountId, summoner.ProfileIconId, summoner.RevisionDate, summoner.SummonerLevel, summoner.SummonerId, summoner.Name, summoner.Tag}

	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Err()
	if err != nil {
		return nil, err
	}

	err = row.Scan(&summoner.Id, &summoner.CreatedAt, &summoner.UpdatedAt)
	if err != nil {
		return nil, err
	}

	for _, league := range summoner.Leagues {
		query = `
			INSERT INTO leagues (puuid, summoner_id, queue_type, tier, rank, wins, losses, league_points, rated_rating, hot_streak, veteran, fresh_blood, inactive, mini_series_wins, mini_series_losses, mini_series_target, mini_series_progress)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
			ON CONFLICT (puuid, queue_type)
			DO UPDATE SET
				tier = EXCLUDED.tier,
				rank = EXCLUDED.rank,
				wins = EXCLUDED.wins,
				losses = EXCLUDED.losses,
				league_points = EXCLUDED.league_points,
				rated_rating = EXCLUDED.rated_rating,
				hot_streak = EXCLUDED.hot_streak,
				veteran = EXCLUDED.veteran,
				fresh_blood = EXCLUDED.fresh_blood,
				inactive = EXCLUDED.inactive,
				mini_series_wins = EXCLUDED.mini_series_wins,
				mini_series_losses = EXCLUDED.mini_series_losses,
				mini_series_target = EXCLUDED.mini_series_target,
				mini_series_progress = EXCLUDED.mini_series_progress,
				updated_at = now()
			RETURNING id, created_at, updated_at
			`

		args := []any{league.Puuid, league.SummonerId, league.QueueType, league.Tier, league.Rank, league.Wins, league.Losses, league.LeaguePoints, league.RatedRating, league.HotStreak, league.Veteran, league.FreshBlood, league.Inactive, league.MiniSeries.Wins, league.MiniSeries.Losses, league.MiniSeries.Target, league.MiniSeries.Progress}
		row := tx.QueryRowContext(ctx, query, args...)
		err = row.Err()
		if err != nil {
			return nil, err
		}

		err = row.Scan(&league.Id, &league.CreatedAt, &league.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
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

	summoner := Summoner{}
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
