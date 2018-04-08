package ionia

import "net/http"

// TournamentService represents the Tournament-Stub-V3 API methods.
// https://developer.riotgames.com/api-methods/#tournament-v3
type TournamentService service

// TournamentCodesOption is a function that modifies the TournamentCodesOptions.
type TournamentCodesOption func(*TournamentCodesOptions)

// TournamentCode represents the required parameters for the Tournament codes service method.
type TournamentCode struct {
	// Optional list of participants in order to validate the players eligible to join the lobby.
	// NOTE: Riot does not currently enforce participants at the team level, but rather the
	// aggregate of teamOne and teamTwo. Riot may add the ability to enforce at the team level
	// in the future.
	AllowedSummonerIDs []int `json:"allowedSummonerIds"`

	// The map type of the game. (Legal values: SUMMONERS_RIFT, TWISTED_TREELINE, HOWLING_ABYSS)
	MapType string `json:"mapType"`

	// Optional string that may contain any data in any format, if specified at all.
	// Used to denote any custom information about the game.
	Metadata string `json:"metadata"`

	// The pick type of the game.
	// (Legal values: BLIND_PICK, DRAFT_MODE, ALL_RANDOM, TOURNAMENT_DRAFT)
	PickType string `json:"pickType"`

	// The spectator type of the game. (Legal values: NONE, LOBBYONLY, ALL)
	SpectatorType string `json:"spectatorType"`

	// The team size of the game. Valid values are 1-5.
	TeamSize int `json:"teamSize"`
}

// TournamentCodesOptions specifies the optional parameters for the Tournament Stub codes service method.
type TournamentCodesOptions struct {
	// The number of codes to create (max 1000).
	Count int `url:"count"`

	// The tournament ID.
	TournamentID int64 `url:"tournamentId"`
}

// Codes creates a tournament code for the given tournament.
func (t *TournamentService) Codes(tc *TournamentCode, opts ...TournamentCodesOption) ([]string, *http.Response, error) {
	options := &TournamentCodesOptions{}
	for _, o := range opts {
		o(options)
	}

	u := "lol/tournament/v3/codes"
	u, err := addOptions(u, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := t.client.NewRequest(http.MethodPost, u, tc)
	if err != nil {
		return nil, nil, err
	}

	var c []string
	resp, err := t.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// TournamentCodeUpdate specifies the optional body parameters for the Tournament code update service method.
type TournamentCodeUpdate struct {
	AllowedSummonerIDs []int  `json:"allowedSummonerIds"`
	MapType            string `json:"mapType"`
	PickType           string `json:"pickType"`
	SpectatorType      string `json:"spectatorType"`
}

// UpdateTournament updates the pick type, map, spectator type, or allowed summoners for the given code.
func (t *TournamentService) UpdateTournament(tournamentCode string, tcu *TournamentCodeUpdate) (*http.Response, error) {
	req, err := t.client.NewRequest(http.MethodPut, "lol/tournament/v3/codes/"+tournamentCode, tcu)
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// TournamentCodeDTO contains tournament code data.
type TournamentCodeDTO struct {
	Map          string  `json:"map"`
	Code         string  `json:"code"`
	Spectators   string  `json:"spectators"`
	Region       string  `json:"region"`
	ProviderID   int     `json:"providerId"`
	TeamSize     int     `json:"teamSize"`
	Participants []int64 `json:"participants"`
	PickType     string  `json:"pickType"`
	TournamentID int     `json:"tournamentId"`
	LobbyName    string  `json:"lobbyName"`
	Password     string  `json:"password"`
	ID           int     `json:"id"`
	Metadata     string  `json:"metaData"`
}

// TournamentCode retrieves the tournament code information for the given tournament code.
func (t *TournamentService) TournamentCode(tournamentCode string) (*TournamentCodeDTO, *http.Response, error) {
	req, err := t.client.NewRequest(http.MethodGet, "lol/tournament/v3/codes/"+tournamentCode, nil)
	if err != nil {
		return nil, nil, err
	}

	tc := &TournamentCodeDTO{}
	resp, err := t.client.Do(req, tc)
	if err != nil {
		return nil, resp, err
	}

	return tc, resp, nil
}

// LobbyEventDTOWrapper contains a slice of LobbyEventDTO.
type LobbyEventDTOWrapper struct {
	EventList []LobbyEventDTO `json:"eventList"`
}

// LobbyEventDTO contains lobby event data.
type LobbyEventDTO struct {
	EventType  string `json:"eventType"`
	SummonerID string `json:"summonerId"`
	Timestamp  string `json:"timestamp"`
}

// LobbyEvents retrieves a list of lobby events by tournament code.
func (t *TournamentService) LobbyEvents(tournamentCode string) (*LobbyEventDTOWrapper, *http.Response, error) {
	req, err := t.client.NewRequest(http.MethodGet, "lol/tournament/v3/lobby-events/by-code/"+tournamentCode, nil)
	if err != nil {
		return nil, nil, err
	}

	l := &LobbyEventDTOWrapper{}
	resp, err := t.client.Do(req, l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, nil
}

// ProviderRegistration specifies the required parameters for the Tournament Stub providers service method.
type ProviderRegistration struct {
	// The provider's callback URL to which tournament game results in this region should
	// be posted. The URL must be well-formed, use the http or https protocol, and use
	// the default port for the protocol (http URLs must use port 80, https URLs must
	// use port 443).
	URL string `json:"url"`

	// The region in which the provider will be running tournaments.
	// (Legal values: BR, EUNE, EUW, JP, LAN, LAS, NA, OCE, PBE, RU, TR)
	Region string `json:"region"`
}

// Provider creates a tournament provider and returns its ID.
func (t *TournamentService) Provider(pr *ProviderRegistration) (*int, *http.Response, error) {
	req, err := t.client.NewRequest(http.MethodPost, "lol/tournament/v3/providers", pr)
	if err != nil {
		return nil, nil, err
	}

	var id *int
	resp, err := t.client.Do(req, id)
	if err != nil {
		return nil, resp, err
	}

	return id, resp, nil
}

// TournamentRegistration specifies the required parameters for the Tournament Stub tournaments service method.
type TournamentRegistration struct {
	// The optional name of the tournament.
	Name string `json:"name"`

	// The provider ID to specify the regional registered provider data to associate this tournament.
	ProviderID int `json:"providerId"`
}

// Tournament creates a mock tournament and returns its ID.
func (t *TournamentService) Tournament(tr *TournamentRegistration) (*int, *http.Response, error) {
	req, err := t.client.NewRequest(http.MethodPost, "lol/tournament/v3/tournaments", tr)
	if err != nil {
		return nil, nil, err
	}

	var id *int
	resp, err := t.client.Do(req, id)
	if err != nil {
		return nil, resp, err
	}

	return id, resp, nil
}
