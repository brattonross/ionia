package ionia

import "net/http"

// TournamentStubService represents the Tournament-Stub-V3 API methods.
// https://developer.riotgames.com/api-methods/#tournament-stub-v3
type TournamentStubService service

// Codes creates a mock tournament code for the given tournament.
func (t *TournamentStubService) Codes(tc *TournamentCode, opts ...TournamentCodesOption) ([]string, *http.Response, error) {
	options := &TournamentCodesOptions{}
	for _, o := range opts {
		o(options)
	}

	u := "lol/tournament-stub/v3/codes"
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

// LobbyEventsByCode retrieves a list of lobby events for the given tournament code.
func (t *TournamentStubService) LobbyEventsByCode(tournamentCode string) (*LobbyEventDTOWrapper, *http.Response, error) {
	req, err := t.client.NewRequest(http.MethodGet, "lol/tournament-stub/v3/lobby-events/by-code/"+tournamentCode, nil)
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

// Provider creates a mock tournament provider and returns its ID.
func (t *TournamentStubService) Provider(pr *ProviderRegistration) (*int, *http.Response, error) {
	req, err := t.client.NewRequest(http.MethodPost, "lol/tournament-stub/v3/providers", pr)
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

// Tournament creates a mock tournament and returns its ID.
func (t *TournamentStubService) Tournament(tr *TournamentRegistration) (*int, *http.Response, error) {
	req, err := t.client.NewRequest(http.MethodPost, "lol/tournament-stub/v3/tournaments", tr)
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
