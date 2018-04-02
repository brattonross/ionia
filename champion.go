package ionia

import (
	"net/http"
	"strconv"
)

// ChampionService represents the Champion-V3 API methods.
// https://developer.riotgames.com/api-methods/#champion-v3
type ChampionService service

// ChampionListDTO contains a collection of champion information.
type ChampionListDTO struct {
	Champions []ChampionDTO `json:"champions"`
}

// ChampionDTO contains champion information.
type ChampionDTO struct {
	RankedPlayEnabled bool  `json:"rankedPlayEnabled"`
	BotEnabled        bool  `json:"botEnabled"`
	BotMmEnabled      bool  `json:"botMmEnabled"`
	Active            bool  `json:"active"`
	FreeToPlay        bool  `json:"freeToPlay"`
	ID                int64 `json:"id"`
}

// All lists all of the currently available champions.
func (c *ChampionService) All() (*ChampionListDTO, *http.Response, error) {
	req, err := c.client.NewRequest("lol/platform/v3/champions")
	if err != nil {
		return nil, nil, err
	}

	cl := &ChampionListDTO{}
	resp, err := c.client.Do(req, cl)
	if err != nil {
		return nil, resp, err
	}

	return cl, resp, nil
}

// ByID retrieves champion information by ID.
func (c *ChampionService) ByID(id int) (*ChampionDTO, *http.Response, error) {
	req, err := c.client.NewRequest("lol/platform/v3/champions/" + strconv.Itoa(id))
	if err != nil {
		return nil, nil, err
	}

	champion := &ChampionDTO{}
	resp, err := c.client.Do(req, champion)
	if err != nil {
		return nil, resp, err
	}

	return champion, resp, nil
}
