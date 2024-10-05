package lolapi

import (
	"fmt"
	"net/http"
	"time"
)

type Match struct {
	Info     MatchInfo     `json:"info"`
	Metadata MatchMetadata `json:"metadata"`
}

type MatchMetadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type MatchInfo struct {
	Participants []Participant `json:"participants"`
	Teams        []Team        `json:"teams"`

	EndOfGameResult    string `json:"endOfGameResult"`
	GameCreation       int64  `json:"gameCreation"`
	GameDuration       int64  `json:"gameDuration"`
	GameEndTimestamp   int64  `json:"gameEndTimestamp"`
	GameId             int64  `json:"gameId"`
	GameMode           string `json:"gameMode"`
	GameName           string `json:"gameName"`
	GameStartTimestamp int64  `json:"gameStartTimestamp"`
	GameType           string `json:"gameType"`
	GameVersion        string `json:"gameVersion"`
	MapId              int    `json:"mapId"`
	PlatformId         string `json:"platformId"`
	QueueId            int    `json:"queueId"`
	TournamentCode     string `json:"tournamentCode"`
}

type Participant struct {
	Challenges map[string]any `json:"challenges"`
	Missions   Missions       `json:"missions"`
	Perks      Perks          `json:"perks"`

	AllInPings                     int    `json:"allInPings"`
	Assists                        int    `json:"assists"`
	AssistMePings                  int    `json:"assistMePings"`
	BaronKills                     int    `json:"baronKills"`
	BasicPings                     int    `json:"basicPings"`
	BountyLevel                    int    `json:"bountyLevel"`
	ChampExperience                int    `json:"champExperience"`
	ChampLevel                     int    `json:"champLevel"`
	ChampionId                     int    `json:"championId"`
	ChampionName                   string `json:"championName"`
	ChampionTransform              int    `json:"championTransform"`
	CommandPings                   int    `json:"commandPings"`
	ConsumablesPurchased           int    `json:"consumablesPurchased"`
	DamageDealtToBuildings         int    `json:"damageDealtToBuildings"`
	DamageDealtToObjectives        int    `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int    `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int    `json:"damageSelfMitigated"`
	DangerPings                    int    `json:"dangerPings"`
	Deaths                         int    `json:"deaths"`
	DetectorWardsPlaced            int    `json:"detectorWardsPlaced"`
	DoubleKills                    int    `json:"doubleKills"`
	DragonKills                    int    `json:"dragonKills"`
	EligibleForProgression         bool   `json:"eligibleForProgression"`
	EnemyMissingPings              int    `json:"enemyMissingPings"`
	EnemyVisionPings               int    `json:"enemyVisionPings"`
	FirstBloodAssist               bool   `json:"firstBloodAssist"`
	FirstBloodKill                 bool   `json:"firstBloodKill"`
	FirstTowerAssist               bool   `json:"firstTowerAssist"`
	FirstTowerKill                 bool   `json:"firstTowerKill"`
	GameEndedInEarlySurrender      bool   `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender           bool   `json:"gameEndedInSurrender"`
	GetBackPings                   int    `json:"getBackPings"`
	GoldEarned                     int    `json:"goldEarned"`
	GoldSpent                      int    `json:"goldSpent"`
	HoldPings                      int    `json:"holdPings"`
	IndividualPosition             string `json:"individualPosition"`
	InhibitorKills                 int    `json:"inhibitorKills"`
	InhibitorTakedowns             int    `json:"inhibitorTakedowns"`
	InhibitorsLost                 int    `json:"inhibitorsLost"`
	Item0                          int    `json:"item0"`
	Item1                          int    `json:"item1"`
	Item2                          int    `json:"item2"`
	Item3                          int    `json:"item3"`
	Item4                          int    `json:"item4"`
	Item5                          int    `json:"item5"`
	Item6                          int    `json:"item6"`
	ItemsPurchased                 int    `json:"itemsPurchased"`
	KillingSprees                  int    `json:"killingSprees"`
	Kills                          int    `json:"kills"`
	Lane                           string `json:"lane"`
	LargestCriticalStrike          int    `json:"largestCriticalStrike"`
	LargestKillingSpree            int    `json:"largestKillingSpree"`
	LargestMultiKill               int    `json:"largestMultiKill"`
	LongestTimeSpentLiving         int    `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int    `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int    `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int    `json:"magicDamageTaken"`
	NeedVisionPings                int    `json:"needVisionPings"`
	NeutralMinionsKilled           int    `json:"neutralMinionsKilled"`
	NexusKills                     int    `json:"nexusKills"`
	NexusLost                      int    `json:"nexusLost"`
	NexusTakedowns                 int    `json:"nexusTakedowns"`
	ObjectivesStolen               int    `json:"objectivesStolen"`
	ObjectivesStolenAssists        int    `json:"objectivesStolenAssists"`
	OnMyWayPings                   int    `json:"onMyWayPings"`
	ParticipantId                  int    `json:"participantId"`
	PentaKills                     int    `json:"pentaKills"`
	PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
	Placement                      int    `json:"placement"`
	PlayerAugment1                 int    `json:"playerAugment1"`
	PlayerAugment2                 int    `json:"playerAugment2"`
	PlayerAugment3                 int    `json:"playerAugment3"`
	PlayerAugment4                 int    `json:"playerAugment4"`
	PlayerAugment5                 int    `json:"playerAugment5"`
	PlayerAugment6                 int    `json:"playerAugment6"`
	PlayerSubteamId                int    `json:"playerSubteamId"`
	ProfileIcon                    int    `json:"profileIcon"`
	PushPings                      int    `json:"PushPings"`
	Puuid                          string `json:"puuid"`
	QuadraKills                    int    `json:"quadraKills"`
	RiotIdGameName                 string `json:"riotIdGameName"`
	RiotIdName                     string `json:"riotIdName"`
	RiotIdTagline                  string `json:"riotIdTagline"`
	Role                           string `json:"role"`
	SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int    `json:"spell1Casts"`
	Spell2Casts                    int    `json:"spell2Casts"`
	Spell3Casts                    int    `json:"spell3Casts"`
	Spell4Casts                    int    `json:"spell4Casts"`
	SubteamPlacement               int    `json:"subteamPlacement"`
	Summoner1Casts                 int    `json:"summoner1Casts"`
	Summoner1Id                    int    `json:"summoner1Id"`
	Summoner2Casts                 int    `json:"summoner2Casts"`
	Summoner2Id                    int    `json:"summoner2Id"`
	SummonerId                     string `json:"summonerId"`
	SummonerLevel                  int    `json:"summonerLevel"`
	SummonerName                   string `json:"summonerName"`
	TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
	TeamId                         int    `json:"teamId"`
	TeamPosition                   string `json:"teamPosition"`
	TimeCCingOthers                int    `json:"timeCCingOthers"`
	TimePlayed                     int    `json:"timePlayed"`
	TotalAllyJungleMinionsKilled   int    `json:"totalAllyJungleMinionsKilled"`
	TotalDamageDealt               int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int    `json:"totalDamageTaken"`
	TotalEnemyJungleMinionsKilled  int    `json:"totalEnemyJungleMinionsKilled"`
	TotalHeal                      int    `json:"totalHeal"`
	TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int    `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int    `json:"totalUnitsHealed"`
	TripleKills                    int    `json:"tripleKills"`
	TrueDamageDealt                int    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int    `json:"trueDamageTaken"`
	TurretKills                    int    `json:"turretKills"`
	TurretTakedowns                int    `json:"turretTakedowns"`
	TurretsLost                    int    `json:"turretsLost"`
	UnrealKills                    int    `json:"unrealKills"`
	VisionClearedPings             int    `json:"visionClearedPings"`
	VisionScore                    int    `json:"visionScore"`
	VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
	WardsKilled                    int    `json:"wardsKilled"`
	WardsPlaced                    int    `json:"wardsPlaced"`
	Win                            bool   `json:"win"`
}

type Missions struct {
	PlayerScore0  int `json:"playerScore0"`
	PlayerScore1  int `json:"playerScore1"`
	PlayerScore10 int `json:"playerScore10"`
	PlayerScore11 int `json:"playerScore11"`
	PlayerScore2  int `json:"playerScore2"`
	PlayerScore3  int `json:"playerScore3"`
	PlayerScore4  int `json:"playerScore4"`
	PlayerScore5  int `json:"playerScore5"`
	PlayerScore6  int `json:"playerScore6"`
	PlayerScore7  int `json:"playerScore7"`
	PlayerScore8  int `json:"playerScore8"`
	PlayerScore9  int `json:"playerScore9"`
}

type Perks struct {
	StatPerks PerkStats   `json:"statPerks"`
	Styles    []PerkStyle `json:"styles"`
}

type PerkStats struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type PerkStyle struct {
	Description string `json:"description"`
	Selections  []struct {
		Perk int `json:"perk"`
		Var1 int `json:"var1"`
		Var2 int `json:"var2"`
		Var3 int `json:"var3"`
	} `json:"selections"`
	Style int `json:"style"`
}

type Team struct {
	Bans       []Ban      `json:"bans"`
	Objectives Objectives `json:"objectives"`
	TeamId     int        `json:"teamId"`
	Win        bool       `json:"win"`
}

type Ban struct {
	ChampionId int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type Objectives struct {
	Baron      Objective `json:"baron"`
	Champion   Objective `json:"champion"`
	Dragon     Objective `json:"dragon"`
	Horde      Objective `json:"horde"`
	Inhibitor  Objective `json:"inhibitor"`
	RiftHerald Objective `json:"riftHerald"`
	Tower      Objective `json:"tower"`
}

type Objective struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

type MatchesRequestOptions struct {
	StartTime *time.Time
	EndTime   *time.Time
	QueueId   *int
	Type      *string
	Start     *int
	Count     *int
}

func (p *Platform) GetPlayerMatches(puuid string, options MatchesRequestOptions) ([]string, error) {
	url := p.makeUrl(fmt.Sprintf("/lol/match/v5/matches/by-puuid/%s/ids", puuid))
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
	if options.QueueId != nil {
		q.Set("queue", fmt.Sprintf("%d", *options.QueueId))
	}
	if options.Type != nil {
		q.Set("type", *options.Type)
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
		err = readJSON(resp.Body, &matches)
		if err != nil {
			return nil, err
		}
		return matches, nil
	default:
		return nil, fmt.Errorf("lolapi: unknown response status code: %d", resp.StatusCode)
	}
}

func (p *Platform) GetMatch(matchId string) (*Match, error) {
	url := p.makeUrl(fmt.Sprintf("/lol/match/v5/matches/%s", matchId))
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
		err = readJSON(resp.Body, &match)
		if err != nil {
			return nil, err
		}
		return &match, nil
	case http.StatusNotFound:
		return nil, fmt.Errorf("lolapi: match %s not found", matchId)
	default:
		return nil, fmt.Errorf("lolapi: unknown response status code: %d", resp.StatusCode)
	}
}
