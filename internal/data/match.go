package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/mfroeh/metrix/internal/lolapi"
)

type Match struct {
	// tft-match-v1
	Participants    []Participant `json:"participants"`
	DataVersion     string        `json:"dataVersion"`
	MatchId         string        `json:"matchId"`
	GameCreation    time.Time     `json:"gameCreation"`
	GameDatetime    time.Time     `json:"gameDatetime"`
	GameLength      float64       `json:"gameLength"`
	GameVersion     string        `json:"gameVersion"`
	GameId          int64         `json:"gameId"`
	EndOfGameResult string        `json:"endOfGameResult"`
	MapId           int           `json:"mapId"`

	QueueId int `json:"queueId"`

	TftGameType  string `json:"tftGameType"`
	TftSetNumber int    `json:"tftSetNumber"`
	TftSetName   string `json:"tftSetName"`

	// Database internal
	Id        int       `json:"-"`
	CreatedAt time.Time `json:"-"`
}

type Participant struct {
	Puuid                string    `json:"puuid"`
	Augments             []string  `json:"augments"`
	Companion            Companion `json:"companion"`
	GoldLeft             int       `json:"goldLeft"`
	LastRound            int       `json:"lastRound"`
	Level                int       `json:"level"`
	Placement            int       `json:"placement"`
	PlayerEliminated     int       `json:"playerEliminated"`
	TotalDamageToPlayers int       `json:"totalDamageToPlayers"`
	TimeEliminated       float64   `json:"timeEliminated"`
	Traits               []Trait   `json:"traits"`
	Units                []Unit    `json:"units"`

	// Database internal
	Id int `json:"-"`
}

type Companion struct {
	ContentId string `json:"contentId"`
	ItemId    int    `json:"itemId"`
	SkinId    int    `json:"skinId"`
	Species   string `json:"species"`
}

type Trait struct {
	Name        string `json:"name"`
	NumUnits    int    `json:"numUnits"`
	Style       int    `json:"style"`
	TierCurrent int    `json:"tierCurrent"`
	TierTotal   int    `json:"tierTotal"`

	// Database internal
	Id int `json:"-"`
}

type Unit struct {
	CharacterId string   `json:"characterId"`
	Items       []string `json:"items"`
	Name        string   `json:"name"`
	Rarity      int      `json:"rarity"`
	Tier        int      `json:"tier"`

	// Database internal
	Id int `json:"-"`
}

type MatchModel struct {
	DB *sql.DB
}

func MatchFromApiMatch(match *lolapi.Match) *Match {
	participants := make([]Participant, len(match.Info.Participants))
	for i, participant := range match.Info.Participants {

		traits := make([]Trait, len(participant.Traits))
		for j, trait := range participant.Traits {
			traits[j] = Trait{
				Name:        trait.Name,
				NumUnits:    trait.NumUnits,
				Style:       trait.Style,
				TierCurrent: trait.TierCurrent,
				TierTotal:   trait.TierTotal,
			}
		}

		units := make([]Unit, len(participant.Units))
		for j, unit := range participant.Units {
			units[j] = Unit{
				CharacterId: unit.CharacterId,
				Name:        unit.Name,
				Rarity:      unit.Rarity,
				Tier:        unit.Tier,
				Items:       unit.ItemNames,
			}
		}

		participants[i] = Participant{
			Puuid:    participant.Puuid,
			Augments: participant.Augments,
			Companion: Companion{
				ContentId: participant.Companion.ContentId,
				ItemId:    participant.Companion.ItemId,
				SkinId:    participant.Companion.SkinId,
				Species:   participant.Companion.Species,
			},
			GoldLeft:             participant.GoldLeft,
			LastRound:            participant.LastRound,
			Level:                participant.Level,
			Placement:            participant.Placement,
			PlayerEliminated:     participant.PlayerEliminated,
			TotalDamageToPlayers: participant.TotalDamageToPlayers,
			TimeEliminated:       participant.TimeEliminated,
			Traits:               traits,
			Units:                units,
		}
	}

	return &Match{
		MatchId:      match.Metadata.MatchId,
		DataVersion:  match.Metadata.DataVersion,
		GameDatetime: time.UnixMilli(match.Info.GameDatetime),
		GameLength:   match.Info.GameLength,
		GameVersion:  match.Info.GameVersion,
		GameId:       match.Info.GameId,
		QueueId:      match.Info.QueueId,
		TftGameType:  match.Info.TftGameType,
		TftSetNumber: match.Info.TftSetNumber,
		TftSetName:   match.Info.TftSetCoreName,
		Participants: participants,
	}

}

func (m *MatchModel) Insert(match *Match) (*Match, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO matches (
			match_id,
			data_version,
			end_of_game_result,
			game_creation,
			game_datetime,
			game_id,
			game_length,
			game_version,
			map_id,
			queue_id,
			tft_game_type,
			tft_set_number,
			tft_set_name
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		)
		RETURNING id
	`

	args := [13]any{
		match.MatchId,
		match.DataVersion,
		match.EndOfGameResult,
		match.GameCreation,
		match.GameDatetime,
		match.GameId,
		match.GameLength,
		match.GameVersion,
		match.MapId,
		match.QueueId,
		match.TftGameType,
		match.TftSetNumber,
		match.TftSetName,
	}

	err = tx.QueryRow(query, args[:]...).Scan(&match.Id)
	if err != nil {
		return nil, err
	}

	for i := range match.Participants {
		participant := &match.Participants[i]

		query = `
			INSERT INTO matches_participants (
				match_id,
				puuid,
				augments,
				gold_left,
				last_round,
				level,
				placement,
				player_eliminated,
				total_damage_to_players,
				time_eliminated,
				companion_content_id,
				companion_item_id,
				companion_skin_id,
				companion_species
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
			)
			RETURNING id
		`

		args := [14]any{
			match.Id,
			participant.Puuid,
			pq.Array(participant.Augments),
			participant.GoldLeft,
			participant.LastRound,
			participant.Level,
			participant.Placement,
			participant.PlayerEliminated,
			participant.TotalDamageToPlayers,
			participant.TimeEliminated,
			participant.Companion.ContentId,
			participant.Companion.ItemId,
			participant.Companion.SkinId,
			participant.Companion.Species,
		}
		err = tx.QueryRow(query, args[:]...).Scan(&participant.Id)
		if err != nil {
			return nil, err
		}

		for j := range participant.Traits {
			trait := &participant.Traits[j]

			query = `
				INSERT INTO matches_participants_traits (
					match_id,
					participant_id,
					name,
					num_units,
					style,
					tier_current,
					tier_total
				) VALUES (
					$1, $2, $3, $4, $5, $6, $7
				)
				RETURNING id
			`

			args := [7]any{
				match.Id,
				participant.Id,
				trait.Name,
				trait.NumUnits,
				trait.Style,
				trait.TierCurrent,
				trait.TierTotal,
			}

			err = tx.QueryRow(query, args[:]...).Scan(&trait.Id)
			if err != nil {
				return nil, err
			}
		}

		for j := range participant.Units {
			unit := &participant.Units[j]

			query = `
				INSERT INTO matches_participants_units (
					match_id,
					participant_id,
					character_id,
					items,
					name,
					rarity,
					tier
				) VALUES (
					$1, $2, $3, $4, $5, $6, $7
				)
				RETURNING id
			`

			args := [7]any{
				match.Id,
				participant.Id,
				unit.CharacterId,
				pq.Array(unit.Items),
				unit.Name,
				unit.Rarity,
				unit.Tier,
			}

			err = tx.QueryRow(query, args[:]...).Scan(&unit.Id)
			if err != nil {
				return nil, err
			}
		}
	}

	return match, tx.Commit()
}
