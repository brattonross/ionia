package ionia

import (
	"fmt"
	"net/http"
	"strconv"
)

// ChampionMasteryService represents the Champion-Mastery-V3 API methods.
// https://developer.riotgames.com/api-methods/#champion-mastery-v3
type ChampionMasteryService service

// ChampionMasteryDTO contains Champion Mastery information for player and champion combination.
type ChampionMasteryDTO struct {
	ChestGranted                 bool  `json:"chestGranted"`
	ChampionLevel                int   `json:"championLevel"`
	ChampionPoints               int   `json:"championPoints"`
	ChampionID                   int64 `json:"championId"`
	PlayerID                     int64 `json:"playerId"`
	ChampionPointsUntilNextLevel int64 `json:"championPointsUntilNextLevel"`
	TokensEarned                 int   `json:"tokensEarned"`
	ChampionPointsSinceLastLevel int64 `json:"championPointsSinceLastLevel"`
	LastPlayTime                 int64 `json:"lastPlayTime"`
}

// MasteryBySummonerID gets all champion mastery entries sorted by number of champion points in descending order.
func (c *ChampionMasteryService) MasteryBySummonerID(summonerID int64) ([]ChampionMasteryDTO, *http.Response, error) {
	req, err := c.client.NewRequest(http.MethodGet, "lol/champion-mastery/v3/champion-masteries/by-summoner/"+strconv.FormatInt(summonerID, 10), nil)
	if err != nil {
		return nil, nil, err
	}

	cm := []ChampionMasteryDTO{}
	resp, err := c.client.Do(req, &cm)
	if err != nil {
		return nil, resp, err
	}

	return cm, resp, nil
}

// BySummonerAndChampionID gets a champion mastery for the given summoner ID and champion ID combination.
func (c *ChampionMasteryService) BySummonerAndChampionID(summonerID, championID int64) (*ChampionMasteryDTO, *http.Response, error) {
	req, err := c.client.NewRequest(http.MethodGet, fmt.Sprintf("lol/champion-mastery/v3/champion-masteries/by-summoner/%d/by-champion/%d", summonerID, championID), nil)
	if err != nil {
		return nil, nil, err
	}

	cm := &ChampionMasteryDTO{}
	resp, err := c.client.Do(req, cm)
	if err != nil {
		return nil, resp, err
	}

	return cm, resp, nil
}

// ScoreBySummonerID gets a summoner's total champion mastery score by ID.
// The score is the sum of all individual champion mastery levels.
func (c *ChampionMasteryService) ScoreBySummonerID(summonerID int64) (int, *http.Response, error) {
	req, err := c.client.NewRequest(http.MethodGet, "lol/champion-mastery/v3/scores/by-summoner/"+strconv.FormatInt(summonerID, 10), nil)
	if err != nil {
		return 0, nil, err
	}

	var score int
	resp, err := c.client.Do(req, &score)
	if err != nil {
		return 0, resp, err
	}

	return score, resp, nil
}
