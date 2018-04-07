package ionia

import (
	"net/http"
	"strconv"
)

// ThirdPartyCodeService represents the Match-V3 API methods.
// https://developer.riotgames.com/api-methods/#third-party-code-v3
type ThirdPartyCodeService service

// BySummonerID retrieves third party code for a given summoner ID.
func (t *ThirdPartyCodeService) BySummonerID(summonerID int64) (*string, *http.Response, error) {
	req, err := t.client.NewRequest("lol/platform/v3/third-party-code/by-summoner/" + strconv.FormatInt(summonerID, 10))
	if err != nil {
		return nil, nil, err
	}

	var s *string
	resp, err := t.client.Do(req, s)
	if err != nil {
		return nil, resp, err
	}

	return s, resp, nil
}
