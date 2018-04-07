package ionia

import (
	"fmt"
	"net/http"
	"strconv"
)

// MatchService represents the Match-V3 API methods.
// https://developer.riotgames.com/api-methods/#match-v3
type MatchService service

// MatchDTO contains match data.
type MatchDTO struct {
	SeasonID              int                      `json:"seasonId"`
	QueueID               int                      `json:"queueId"`
	GameID                int64                    `json:"gameId"`
	ParticipantIdentities []ParticipantIdentityDTO `json:"participantIdentities"`
	GameVersion           string                   `json:"gameVersion"`
	PlatformID            string                   `json:"platformId"`
	GameMode              string                   `json:"gameMode"`
	MapID                 int                      `json:"mapId"`
	GameType              string                   `json:"gameType"`
	Teams                 []TeamStatsDTO           `json:"teams"`
	Participants          []ParticipantDTO         `json:"participants"`
	GameDuration          int64                    `json:"gameDuration"`
	GameCreation          int64                    `json:"gameCreation"`
}

// ParticipantIdentityDTO contains participant identity data.
type ParticipantIdentityDTO struct {
	Player        PlayerDTO `json:"player"`
	ParticipantID int       `json:"participantId"`
}

// PlayerDTO contains player data.
type PlayerDTO struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	PlatformID        string `json:"platformId"`
	CurrentAccountID  int64  `json:"currentAccountId"`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerID        int64  `json:"summonerId"`
	AccountID         int64  `json:"accountId"`
}

// TeamStatsDTO contains team stats data.
type TeamStatsDTO struct {
	FirstDragon          bool          `json:"firstDragon"`
	FirstInhibitor       bool          `json:"firstInhibitor"`
	Bans                 []TeamBansDTO `json:"bans"`
	BaronKills           int           `json:"baronKills"`
	FirstRiftHerald      bool          `json:"firstRiftHerald"`
	FirstBaron           bool          `json:"firstBaron"`
	RiftHeraldKills      int           `json:"riftHeraldKills"`
	FirstBlood           bool          `json:"firstBlood"`
	TeamID               int           `json:"teamId"`
	FirstTower           bool          `json:"firstTower"`
	VilemawKills         int           `json:"vilemawKills"`
	InhibitorKills       int           `json:"inhibitorKills"`
	TowerKills           int           `json:"towerKills"`
	DominionVictoryScore int           `json:"dominionVictoryScore"`
	Win                  string        `json:"win"`
	DragonKills          int           `json:"dragonKills"`
}

// TeamBansDTO contains team bans data.
type TeamBansDTO struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
}

// ParticipantDTO contains participant data.
type ParticipantDTO struct {
	Stats                     ParticipantStatsDTO    `json:"stats"`
	ParticipantID             int                    `json:"participantId"`
	Runes                     []RuneDTO              `json:"runes"`
	Timeline                  ParticipantTimelineDTO `json:"timeline"`
	TeamID                    int                    `json:"teamId"`
	Spell2ID                  int                    `json:"spell2Id"`
	Masteries                 []MasteryDTO           `json:"masteries"`
	HighestAchievedSeasonTier string                 `json:"highestAchievedSeasonTier"`
	Spell1ID                  int                    `json:"spell1Id"`
	ChampionID                int                    `json:"championId"`
}

// ParticipantStatsDTO contains participant stats data.
type ParticipantStatsDTO struct {
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
	NeutralMinionsKilledTeamJungle  int   `json:"neutralMinionsKilledTeamJungle"`
	MagicDamageDealt                int64 `json:"magicDamageDealt"`
	TotalPlayerScore                int   `json:"totalPlayerScore"`
	Deaths                          int   `json:"deaths"`
	Win                             bool  `json:"win"`
	NeutralMinionsKilledEnemyJungle int   `json:"neutralMinionsKilledEnemyJungle"`
	AltarsCaptured                  int   `json:"altarsCaptured"`
	LargestCriticalStrike           int   `json:"largestCriticalStrike"`
	TotalDamageDealt                int64 `json:"totalDamageDealt"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`
	VisionWardsBoughtInGame         int   `json:"visionWardsBoughtInGame"`
	DamageDealtToObjectives         int64 `json:"damageDealtToObjectives"`
	LargestKillingSpree             int   `json:"largestKillingSpree"`
	Item1                           int   `json:"item1"`
	QuadraKills                     int   `json:"quadraKills"`
	TeamObjective                   int   `json:"teamObjective"`
	TotalTimeCrowdControlDealt      int   `json:"totalTimeCrowdControlDealt"`
	LongestTimeSpentLiving          int   `json:"longestTimeSpentLiving"`
	WardsKilled                     int   `json:"wardsKilled"`
	FirstTowerAssist                bool  `json:"firstTowerAssist"`
	FirstTowerKill                  bool  `json:"firstTowerAssist"`
	Item2                           int   `json:"item2"`
	Item3                           int   `json:"item3"`
	Item0                           int   `json:"item0"`
	FirstBloodAssist                bool  `json:"firstBloodAssist"`
	VisionScore                     int64 `json:"visionScore"`
	WardsPlaced                     int   `json:"wardsPlaced"`
	Item4                           int   `json:"item4"`
	Item5                           int   `json:"item5"`
	Item6                           int   `json:"item6"`
	TurretKills                     int   `json:"turretKills"`
	TripleKills                     int   `json:"tripleKills"`
	DamageSelfMitigated             int64 `json:"damageSelfMitigated"`
	ChampLevel                      int   `json:"champLevel"`
	NodeNeutralizeAssist            int   `json:"nodeNeutralizeAssist"`
	FirstInhibitorKill              bool  `json:"FirstInhibitorKill"`
	GoldEarned                      int   `json:"goldEarned"`
	MagicalDamageTaken              int64 `json:"magicalDamageTaken"`
	Kills                           int   `json:"kills"`
	DoubleKills                     int   `json:"doubleKills"`
	NodeCaptureAssist               int   `json:"nodeCaptureAssist"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken"`
	NodeNeutralize                  int   `json:"nodeNeutralize"`
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`
	Assists                         int   `json:"assists"`
	UnrealKills                     int   `json:"unrealKills"`
	NeutralMinionsKilled            int   `json:"neutralMinionsKilled"`
	ObjectivePlayerScore            int   `json:"objectivePlayerScore"`
	CombatPlayerScore               int   `json:"combatPlayerScore"`
	DamageDealtToTurrets            int64 `json:"damageDealtToTurrets"`
	AltarsNeutralized               int   `json:"altarsNeutralized"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
	GoldSpent                       int   `json:"goldSpent"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
	ParticipantID                   int   `json:"participantId"`
	PentaKills                      int   `json:"pentaKills"`
	TotalHeal                       int64 `json:"totalHeal"`
	TotalMinionsKilled              int   `json:"totalMinionsKilled"`
	FirstBloodKill                  bool  `json:"firstBloodKill"`
	NodeCapture                     int   `json:"nodeCapture"`
	LargestMultiKill                int   `json:"largestMultiKill"`
	SightWardsBoughtInGame          int   `json:"sightWardsBoughtInGame"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
	TotalUnitsHealed                int   `json:"totalUnitsHealed"`
	InhibitorKills                  int   `json:"inhibitorKills"`
	TotalScoreRank                  int   `json:"totalScoreRank"`
	TotalDamageTaken                int64 `json:"totalDamageTaken"`
	KillingSprees                   int   `json:"killingSprees"`
	TimeCCingOthers                 int64 `json:"timeCCingOthers"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
}

// MatchRuneDTO contains rune data for a match.
type MatchRuneDTO struct {
	RuneID int `json:"runeId"`
	Rank   int `json:"rank"`
}

// ParticipantTimelineDTO contains participant timeline data.
type ParticipantTimelineDTO struct {
	Lane                        string             `json:"lane"`
	ParticipantID               int                `json:"participantId"`
	CSDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
	GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
	XPDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
	CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
	CPPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
	Role                        string             `json:"role"`
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
}

// MatchMasteryDTO contains mastery data for a match.
type MatchMasteryDTO struct {
	MasteryID int `json:"masteryId"`
	Rank      int `json:"rank"`
}

// MatchByID retrieves match data by ID.
func (m *MatchService) MatchByID(matchID int64) (*MatchDTO, *http.Response, error) {
	req, err := m.client.NewRequest("lol/match/v3/matches/" + strconv.FormatInt(matchID, 10))
	if err != nil {
		return nil, nil, err
	}

	match := &MatchDTO{}
	resp, err := m.client.Do(req, match)
	if err != nil {
		return nil, resp, err
	}

	return match, resp, nil
}

// MatchlistDTO contains match list data.
type MatchlistDTO struct {
	Matches    []MatchReferenceDTO `json:"matches"`
	TotalGames int                 `json:"totalGames"`
	StartIndex int                 `json:"startIndex"`
	EndIndex   int                 `json:"endIndex"`
}

// MatchReferenceDTO contains match reference data.
type MatchReferenceDTO struct {
	Lane       string `json:"lane"`
	GameID     int64  `json:"gameId"`
	Champion   int    `json:"champion"`
	PlatformID string `json:"platformId"`
	Season     int    `json:"season"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Timestamp  int64  `json:"timestamp"`
}

// MatchByAccountIDOptions specifies the optional parameters
// for the match by account ID service method.
type MatchByAccountIDOptions struct {
	// The end time to use for filtering matchlist specified as epoch milliseconds.
	// If beginTime is specified, but not endTime, then these parameters are ignored.
	// If endTime is specified, but not beginTime, then beginTime defaults to the start
	// of the account's match history. If both are specified, then endTime should be
	// greater than beginTime. The maximum time range allowed is one week, otherwise
	// a 400 error code is returned.
	EndTime int64 `url:"endTime"`

	// The begin index to use for filtering matchlist. If beginIndex is specified,
	// but not endIndex, then endIndex defaults to beginIndex+100. If endIndex is
	// specified, but not beginIndex, then beginIndex defaults to 0. If both are
	// specified, then endIndex must be greater than beginIndex. The maximum range
	// allowed is 100, otherwise a 400 error code is returned.
	BeginIndex int `url:"beginIndex"`

	// The begin time to use for filtering matchlist specified as epoch milliseconds.
	// If beginTime is specified, but not endTime, then these parameters are ignored.
	// If endTime is specified, but not beginTime, then beginTime defaults to the start
	// of the account's match history. If both are specified, then endTime should be
	// greater than beginTime. The maximum time range allowed is one week, otherwise a
	// 400 error code is returned.
	BeginTime int64 `url:"beginTime"`

	// Slice of champion IDs for filtering the matchlist.
	Champion []int `url:"champion"`

	// 	The end index to use for filtering matchlist. If beginIndex is specified, but
	// not endIndex, then endIndex defaults to beginIndex+100. If endIndex is specified,
	// but not beginIndex, then beginIndex defaults to 0. If both are specified, then
	// endIndex must be greater than beginIndex. The maximum range allowed is 100,
	// otherwise a 400 error code is returned.
	EndIndex int `url:"endIndex"`

	// Slice of queue IDs for filtering the matchlist.
	Queue []int `url:"queue"`

	// Slice of season IDs for filtering the matchlist.
	Season []int `url:"season"`
}

// MatchByAccountIDOption is a function which modifies the MatchByAccountIDOptions.
type MatchByAccountIDOption func(*MatchByAccountIDOptions)

// MatchesByAccountID retrieves matches by account ID.
func (m *MatchService) MatchesByAccountID(accountID int64, opts ...MatchByAccountIDOption) (*MatchlistDTO, *http.Response, error) {
	options := &MatchByAccountIDOptions{}
	for _, o := range opts {
		o(options)
	}

	u := "lol/match/v3/matchlists/by-account/" + strconv.FormatInt(accountID, 10)
	u, err := addOptions(u, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := m.client.NewRequest(u)
	if err != nil {
		return nil, nil, err
	}

	ml := &MatchlistDTO{}
	resp, err := m.client.Do(req, ml)
	if err != nil {
		return nil, resp, err
	}

	return ml, resp, nil
}

// RecentMatches retrieves the last 20 matches played on a given account ID.
func (m *MatchService) RecentMatches(accountID int64) (*MatchlistDTO, *http.Response, error) {
	req, err := m.client.NewRequest(fmt.Sprintf("lol/match/v3/matchlists/by-account/%d/recent/", accountID))
	if err != nil {
		return nil, nil, err
	}

	rm := &MatchlistDTO{}
	resp, err := m.client.Do(req, rm)
	if err != nil {
		return nil, resp, err
	}

	return rm, resp, nil
}

// MatchTimelineDTO contains match timeline data.
type MatchTimelineDTO struct {
	Frames        []MatchFrameDTO `json:"frames"`
	FrameInterval int64           `json:"frameInterval"`
}

// MatchFrameDTO contains match frame data.
type MatchFrameDTO struct {
	Timestamp         int64                            `json:"timestamp"`
	ParticipantFrames map[int]MatchParticipantFrameDTO `json:"participantFrames"`
	Events            []MatchEventDTO                  `json:"events"`
}

// MatchParticipantFrameDTO contains participant frame data for a match.
type MatchParticipantFrameDTO struct {
	TotalGold           int              `json:"totalGold"`
	TeamScore           int              `json:"teamScore"`
	ParticipantID       int              `json:"participantId"`
	Level               int              `json:"level"`
	CurrentGold         int              `json:"currentGold"`
	MinionsKilled       int              `json:"minionsKilled"`
	DominionScore       int              `json:"dominionScore"`
	Position            MatchPositionDTO `json:"position"`
	XP                  int              `json:"xp"`
	JungleMinionsKilled int              `json:"jungleMinionsKilled"`
}

// MatchPositionDTO contains information about the position of an object on the map.
type MatchPositionDTO struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// MatchEventDTO contains information about a match event.
//
// Type Legal Values:
// CHAMPION_KILL, WARD_PLACED, WARD_KILL, BUILDING_KILL, ELITE_MONSTER_KILL,
// ITEM_PURCHASED, ITEM_SOLD, ITEM_DESTROYED, ITEM_UNDO, SKILL_LEVEL_UP,
// ASCENDED_EVENT, CAPTURE_POINT, PORO_KING_SUMMON.
type MatchEventDTO struct {
	EventType               string           `json:"eventType"`
	TowerType               string           `json:"towerType"`
	TeamID                  int              `json:"teamId"`
	AscendedType            string           `json:"ascendedType"`
	KillerID                int              `json:"killerId"`
	LevelUpType             string           `json:"levelUpType"`
	PointCaptured           string           `json:"pointCaptured"`
	AssistingParticipantIds []int            `json:"assistingParticipantIds"`
	WardType                string           `json:"wardType"`
	MonsterType             string           `json:"monsterType"`
	Type                    string           `json:"type"`
	SkillSlot               int              `json:"skillSlot"`
	VictimID                int              `json:"victimId"`
	Timestamp               int64            `json:"timestamp"`
	AfterID                 int              `json:"afterId"`
	MonsterSubType          string           `json:"monsterSubType"`
	LaneType                string           `json:"laneType"`
	ItemID                  int              `json:"itemId"`
	ParticipantID           int              `json:"participantId"`
	BuildingType            string           `json:"buildingType"`
	CreatorID               int              `json:"creatorId"`
	Position                MatchPositionDTO `json:"position"`
	BeforeID                int              `json:"beforeId"`
}

// MatchTimelineByID retrieves a match timeline by match ID.
func (m *MatchService) MatchTimelineByID(matchID int64) (*MatchTimelineDTO, *http.Response, error) {
	req, err := m.client.NewRequest("lol/match/v3/timelines/by-match/" + strconv.FormatInt(matchID, 10))
	if err != nil {
		return nil, nil, err
	}

	mt := &MatchTimelineDTO{}
	resp, err := m.client.Do(req, mt)
	if err != nil {
		return nil, resp, err
	}

	return mt, resp, nil
}

// MatchIDsByTournamentCode retrieves match IDs for the given tournament code.
func (m *MatchService) MatchIDsByTournamentCode(tournamentCode string) ([]int64, *http.Response, error) {
	req, err := m.client.NewRequest(fmt.Sprintf("lol/match/v3/matches/by-tournament-code/%s/ids", tournamentCode))
	if err != nil {
		return nil, nil, err
	}

	mids := []int64{}
	resp, err := m.client.Do(req, &mids)
	if err != nil {
		return nil, resp, err
	}

	return mids, resp, nil
}

// MatchByIDAndTournamentCode retrieves a match by ID and tournament code.
func (m *MatchService) MatchByIDAndTournamentCode(matchID int64, tournamentCode string) (*MatchDTO, *http.Response, error) {
	u := fmt.Sprintf("lol/match/v3/matches/%d/by-tournament-code/%s", matchID, tournamentCode)
	req, err := m.client.NewRequest(u)
	if err != nil {
		return nil, nil, err
	}

	match := &MatchDTO{}
	resp, err := m.client.Do(req, match)
	if err != nil {
		return nil, resp, err
	}

	return match, resp, nil
}
