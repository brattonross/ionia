package ionia

import (
	"net/http"
	"strconv"
)

// SummonerService represents the Match-V3 API methods.
// https://developer.riotgames.com/api-methods/#summoner-v3/GET_getByAccountId
type SummonerService service

// SummonerDTO contains summoner information.
type SummonerDTO struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int64  `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
	ID            int64  `json:"id"`
	AccountID     int64  `json:"accountId"`
}

// ByAccountID retrieves a summoner by account ID.
func (s *SummonerService) ByAccountID(accountID int64) (*SummonerDTO, *http.Response, error) {
	req, err := s.client.NewRequest("lol/summoner/v3/summoners/by-account/" + strconv.FormatInt(accountID, 10))
	if err != nil {
		return nil, nil, err
	}

	sm := &SummonerDTO{}
	resp, err := s.client.Do(req, sm)
	if err != nil {
		return nil, resp, err
	}

	return sm, resp, nil
}

// BySummonerName retrieves a summoner by summoner name.
func (s *SummonerService) BySummonerName(summonerName string) (*SummonerDTO, *http.Response, error) {
	req, err := s.client.NewRequest("lol/summoner/v3/summoners/by-name/" + summonerName)
	if err != nil {
		return nil, nil, err
	}

	sm := &SummonerDTO{}
	resp, err := s.client.Do(req, sm)
	if err != nil {
		return nil, resp, err
	}

	return sm, resp, nil
}

// BySummonerID retrives a summoner by summoner ID.
func (s *SummonerService) BySummonerID(summonerID int64) (*SummonerDTO, *http.Response, error) {
	req, err := s.client.NewRequest("lol/summoner/v3/summoners/" + strconv.FormatInt(summonerID, 10))
	if err != nil {
		return nil, nil, err
	}

	sm := &SummonerDTO{}
	resp, err := s.client.Do(req, sm)
	if err != nil {
		return nil, resp, err
	}

	return sm, resp, nil
}
