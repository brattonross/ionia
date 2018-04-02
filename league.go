package ionia

import (
	"net/http"
	"strconv"
)

// LeagueService represents the League-V3 API methods.
// https://developer.riotgames.com/api-methods/#league-v3
type LeagueService service

// LeagueListDTO contains information about a ranked league.
type LeagueListDTO struct {
	LeagueID string          `json:"leagueId"`
	Tier     string          `json:"tier"`
	Entries  []LeagueItemDTO `json:"entries"`
	Queue    string          `json:"queue"`
	Name     string          `json:"name"`
}

// LeagueItemDTO represents a player or ranked team in a league.
type LeagueItemDTO struct {
	Rank             string        `json:"rank"`
	HotStreak        bool          `json:"hotStreak"`
	MiniSeries       MiniSeriesDTO `json:"miniSeries"`
	Wins             int           `json:"wins"`
	Veteran          bool          `json:"veteran"`
	Losses           int           `json:"losses"`
	FreshBlood       bool          `json:"freshBlood"`
	PlayerOrTeamName string        `json:"playerOrTeamName"`
	Inactive         bool          `json:"inactive"`
	PlayerOrTeamID   string        `json:"playerOrTeamId"`
	LeaguePoints     int           `json:"leaguePoints"`
}

// MiniSeriesDTO contains information about a league item's current mini series.
type MiniSeriesDTO struct {
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Progress string `json:"progress"`
}

// ChallengerLeagueByQueue gets the challenger league information for the given queue.
func (l *LeagueService) ChallengerLeagueByQueue(queue string) (*LeagueListDTO, *http.Response, error) {
	req, err := l.client.NewRequest("lol/league/v3/challengerleagues/by-queue/" + queue)
	if err != nil {
		return nil, nil, err
	}

	ll := &LeagueListDTO{}
	resp, err := l.client.Do(req, ll)
	if err != nil {
		return nil, resp, err
	}

	return ll, resp, nil
}

// ByLeagueID gets information for the league with the given ID, including inactive entries.
//
// WARNING: Consistently looking up leagues that don't exist will result in a blacklist.
func (l *LeagueService) ByLeagueID(leagueID string) (*LeagueListDTO, *http.Response, error) {
	req, err := l.client.NewRequest("lol/league/v3/leagues/" + leagueID)
	if err != nil {
		return nil, nil, err
	}

	ll := &LeagueListDTO{}
	resp, err := l.client.Do(req, ll)
	if err != nil {
		return nil, resp, err
	}

	return ll, resp, nil
}

// MasterLeagueByQueue gets the master league information for the given queue.
func (l *LeagueService) MasterLeagueByQueue(queue string) (*LeagueListDTO, *http.Response, error) {
	req, err := l.client.NewRequest("lol/league/v3/masterleagues/by-queue/" + queue)
	if err != nil {
		return nil, nil, err
	}

	ll := &LeagueListDTO{}
	resp, err := l.client.Do(req, ll)
	if err != nil {
		return nil, resp, err
	}

	return ll, resp, nil
}

// LeaguePositionDTO represents the position of a summoner in a league.
type LeaguePositionDTO struct {
	Rank             string        `json:"rank"`
	QueueType        string        `json:"queueType"`
	HotStreak        bool          `json:"hotStreak"`
	MiniSeries       MiniSeriesDTO `json:"miniSeries"`
	Wins             int           `json:"wins"`
	Veteran          bool          `json:"veteran"`
	Losses           int           `json:"losses"`
	FreshBlood       bool          `json:"freshBlood"`
	LeagueID         string        `json:"leagueId"`
	PlayerOrTeamName string        `json:"playerOrTeamName"`
	Inactive         bool          `json:"inactive"`
	PlayerOrTeamID   string        `json:"playerOrTeamId"`
	LeagueName       string        `json:"leagueName"`
	Tier             string        `json:"tier"`
	LeaguePoints     int           `json:"leaguePoints"`
}

// PositionsBySummonerID gets league positions in all queues for the given summoner ID.
func (l *LeagueService) PositionsBySummonerID(summonerID int64) ([]LeaguePositionDTO, *http.Response, error) {
	req, err := l.client.NewRequest("lol/league/v3/positions/by-summoner/" + strconv.FormatInt(summonerID, 10))
	if err != nil {
		return nil, nil, err
	}

	lp := []LeaguePositionDTO{}
	resp, err := l.client.Do(req, &lp)
	if err != nil {
		return nil, resp, err
	}

	return lp, resp, nil
}
