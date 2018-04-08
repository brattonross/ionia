package ionia

import (
	"net/http"
	"strconv"
)

// SpectatorService represents the Spectator-V3 API methods.
// https://developer.riotgames.com/api-methods/#spectator-v3
type SpectatorService service

// CurrentGameInfo contains current game info.
type CurrentGameInfo struct {
	GameID            int64                    `json:"gameId"`
	GameStartTime     int64                    `json:"gameStartTime"`
	PlatformID        string                   `json:"platformId"`
	GameMode          string                   `json:"gameMode"`
	MapID             int64                    `json:"mapId"`
	GameType          string                   `json:"gameType"`
	BannedChampions   []BannedChampion         `json:"bannedChampions"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
	GameLength        int64                    `json:"gameLength"`
	GameQueueConfigID int64                    `json:"gameQueueConfigId"`
}

// BannedChampion contains information about a banned champion.
type BannedChampion struct {
	PickTurn   int   `json:"pickTurn"`
	ChampionID int64 `json:"championId"`
	TeamID     int64 `json:"teamId"`
}

// Observer contains observer information.
type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

// CurrentGameParticipant contains current game participant information.
type CurrentGameParticipant struct {
	ProfileIconID            int64                     `json:"profileIconId"`
	ChampionID               int64                     `json:"championId"`
	SummonerName             string                    `json:"summonerName"`
	GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
	Bot                      bool                      `json:"bot"`
	Perks                    Perks                     `json:"perks"`
	Spell2ID                 int64                     `json:"spell2Id"`
	TeamID                   int64                     `json:"teamId"`
	Spell1ID                 int64                     `json:"spell1Id"`
	SummonerID               int64                     `json:"summonerId"`
}

// GameCustomizationObject contains game customization information.
type GameCustomizationObject struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}

// Perks contains perks information.
type Perks struct {
	PerkStyle    int64   `json:"perkStyle"`
	PerkIDs      []int64 `json:"perkids"`
	PerkSubStyle int64   `json:"perkSubStyle"`
}

// CurrentGame retrieves the current game information for the given summoner ID.
func (s *SpectatorService) CurrentGame(summonerID int64) (*CurrentGameInfo, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "lol/spectator/v3/active-games/by-summoner/"+strconv.FormatInt(summonerID, 10), nil)
	if err != nil {
		return nil, nil, err
	}

	c := &CurrentGameInfo{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// FeaturedGames contains featured games information.
type FeaturedGames struct {
	ClientRefreshInterval int64              `json:"clientRefreshInterval"`
	GameList              []FeaturedGameInfo `json:"gameList"`
}

// FeaturedGameInfo contains featured game information.
type FeaturedGameInfo struct {
	GameID            int64            `json:"gameId"`
	GameStartTime     int64            `json:"gameStartTime"`
	PlatformID        string           `json:"platformId"`
	GameMode          string           `json:"gameMode"`
	MapID             int64            `json:"mapId"`
	GameType          string           `json:"gameType"`
	BannedChampions   []BannedChampion `json:"bannedChampions"`
	Observers         Observer         `json:"observers"`
	Participants      []Participant    `json:"participants"`
	GameLength        int64            `json:"gameLength"`
	GameQueueConfigID int64            `json:"gameQueueConfigId"`
}

// Participant contains participant information.
type Participant struct {
	ProfileIconID int64  `json:"profileIconId"`
	ChampionID    int64  `json:"championId"`
	SummonerName  string `json:"summonerName"`
	Bot           bool   `json:"bot"`
	Spell2ID      int64  `json:"spell2Id"`
	TeamID        int64  `json:"teamId"`
	Spell1ID      int64  `json:"spell1Id"`
}

// FeaturedGames retrieves a list of featured games.
func (s *SpectatorService) FeaturedGames() (*FeaturedGames, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "lol/spectator/v3/featured-games", nil)
	if err != nil {
		return nil, nil, err
	}

	f := &FeaturedGames{}
	resp, err := s.client.Do(req, f)
	if err != nil {
		return nil, resp, err
	}

	return f, resp, nil
}
