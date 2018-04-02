package ionia

import (
	"net/http"
	"reflect"
	"testing"
)

func TestChallengerLeagueByQueue(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/league/v3/challengerleagues/by-queue/RANKED_SOLO_5x5", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(challengerLeagueJSON)
	})

	got, _, err := client.League.ChallengerLeagueByQueue("RANKED_SOLO_5x5")
	if err != nil {
		t.Errorf("League.ChallengerLeagueByQueue returned error: %v", err)
	}
	if want := wantChallengerLeague; !reflect.DeepEqual(got, want) {
		t.Errorf("League.ChallengerLeagueByQueue = %+v, want %+v", got, want)
	}
}

func TestLeagueByLeagueID(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/league/v3/leagues/40ac4f31-647b-3960-b95e-45a8ee1ef734", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(leagueJSON)
	})

	got, _, err := client.League.ByLeagueID("40ac4f31-647b-3960-b95e-45a8ee1ef734")
	if err != nil {
		t.Errorf("League.ByLeagueID returned error: %v", err)
	}
	if want := wantLeague; !reflect.DeepEqual(got, want) {
		t.Errorf("League.ByLeagueID = %+v, want %+v", got, want)
	}
}

func TestMasterLeagueByQueue(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/league/v3/masterleagues/by-queue/RANKED_SOLO_5x5", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(masterLeagueJSON)
	})

	got, _, err := client.League.MasterLeagueByQueue("RANKED_SOLO_5x5")
	if err != nil {
		t.Errorf("League.MasterLeagueByQueue returned error: %v", err)
	}
	if want := wantMasterLeague; !reflect.DeepEqual(got, want) {
		t.Errorf("League.MasterLeagueByQueue = %+v, want %+v", got, want)
	}
}

func TestPositionsBySummonerID(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/league/v3/positions/by-summoner/56236881", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(positionsJSON)
	})

	got, _, err := client.League.PositionsBySummonerID(56236881)
	if err != nil {
		t.Errorf("League.PositionsBySummonerID returned error: %v", err)
	}
	if want := wantPositions; !reflect.DeepEqual(got, want) {
		t.Errorf("League.PositionsBySummonerID = %+v, want %+v", got, want)
	}
}

var (
	challengerLeagueJSON = []byte(`{
		"tier": "CHALLENGER",
    	"queue": "RANKED_SOLO_5x5",
    	"leagueId": "40ac4f31-647b-3960-b95e-45a8ee1ef734",
    	"name": "Leona's Spellswords",
    	"entries": [
			{
				"hotStreak": true,
				"wins": 113,
				"veteran": false,
				"losses": 58,
				"rank": "I",
				"playerOrTeamName": "Kaos Malthael",
				"inactive": false,
				"playerOrTeamId": "56236881",
				"freshBlood": true,
				"leaguePoints": 495
			},
			{
				"hotStreak": true,
				"wins": 225,
				"veteran": false,
				"losses": 186,
				"rank": "I",
				"playerOrTeamName": "Raven2",
				"inactive": false,
				"playerOrTeamId": "57396939",
				"freshBlood": false,
				"leaguePoints": 594
			}
		]
	}`)

	wantChallengerLeague = &LeagueListDTO{
		Tier:     "CHALLENGER",
		Queue:    "RANKED_SOLO_5x5",
		LeagueID: "40ac4f31-647b-3960-b95e-45a8ee1ef734",
		Name:     "Leona's Spellswords",
		Entries: []LeagueItemDTO{
			{
				HotStreak:        true,
				Wins:             113,
				Veteran:          false,
				Losses:           58,
				Rank:             "I",
				PlayerOrTeamName: "Kaos Malthael",
				Inactive:         false,
				PlayerOrTeamID:   "56236881",
				FreshBlood:       true,
				LeaguePoints:     495,
			},
			{
				HotStreak:        true,
				Wins:             225,
				Veteran:          false,
				Losses:           186,
				Rank:             "I",
				PlayerOrTeamName: "Raven2",
				Inactive:         false,
				PlayerOrTeamID:   "57396939",
				FreshBlood:       false,
				LeaguePoints:     594,
			},
		},
	}

	leagueJSON = []byte(`{
		"tier": "CHALLENGER",
		"queue": "RANKED_SOLO_5x5",
		"leagueId": "40ac4f31-647b-3960-b95e-45a8ee1ef734",
		"name": "Leona's Spellswords",
		"entries": [
			{
				"hotStreak": false,
				"wins": 120,
				"veteran": false,
				"losses": 90,
				"rank": "I",
				"playerOrTeamName": "Jiissee",
				"inactive": false,
				"playerOrTeamId": "20779248",
				"freshBlood": false,
				"leaguePoints": 491
			},
			{
				"hotStreak": false,
				"wins": 173,
				"veteran": false,
				"losses": 142,
				"rank": "I",
				"playerOrTeamName": "S04 Pride",
				"inactive": false,
				"playerOrTeamId": "70858329",
				"freshBlood": true,
				"leaguePoints": 604
			}
		]
	}`)

	wantLeague = &LeagueListDTO{
		Tier:     "CHALLENGER",
		Queue:    "RANKED_SOLO_5x5",
		LeagueID: "40ac4f31-647b-3960-b95e-45a8ee1ef734",
		Name:     "Leona's Spellswords",
		Entries: []LeagueItemDTO{
			{
				HotStreak:        false,
				Wins:             120,
				Veteran:          false,
				Losses:           90,
				Rank:             "I",
				PlayerOrTeamName: "Jiissee",
				Inactive:         false,
				PlayerOrTeamID:   "20779248",
				FreshBlood:       false,
				LeaguePoints:     491,
			},
			{
				HotStreak:        false,
				Wins:             173,
				Veteran:          false,
				Losses:           142,
				Rank:             "I",
				PlayerOrTeamName: "S04 Pride",
				Inactive:         false,
				PlayerOrTeamID:   "70858329",
				FreshBlood:       true,
				LeaguePoints:     604,
			},
		},
	}

	masterLeagueJSON = []byte(`{
		"tier": "MASTER",
		"queue": "RANKED_SOLO_5x5",
		"leagueId": "49206c68-7ef1-3012-bb19-22286bb0798a",
		"name": "Urgot's Commandos",
		"entries": [
			{
				"hotStreak": false,
				"wins": 192,
				"veteran": true,
				"losses": 170,
				"rank": "I",
				"playerOrTeamName": "G2V Hero",
				"inactive": false,
				"playerOrTeamId": "32081620",
				"freshBlood": false,
				"leaguePoints": 304
			},
			{
				"hotStreak": false,
				"wins": 136,
				"veteran": false,
				"losses": 119,
				"rank": "I",
				"playerOrTeamName": "Touch",
				"inactive": false,
				"playerOrTeamId": "29576970",
				"freshBlood": false,
				"leaguePoints": 0
			}
		]
	}`)

	wantMasterLeague = &LeagueListDTO{
		Tier:     "MASTER",
		Queue:    "RANKED_SOLO_5x5",
		LeagueID: "49206c68-7ef1-3012-bb19-22286bb0798a",
		Name:     "Urgot's Commandos",
		Entries: []LeagueItemDTO{
			{
				HotStreak:        false,
				Wins:             192,
				Veteran:          true,
				Losses:           170,
				Rank:             "I",
				PlayerOrTeamName: "G2V Hero",
				Inactive:         false,
				PlayerOrTeamID:   "32081620",
				FreshBlood:       false,
				LeaguePoints:     304,
			},
			{
				HotStreak:        false,
				Wins:             136,
				Veteran:          false,
				Losses:           119,
				Rank:             "I",
				PlayerOrTeamName: "Touch",
				Inactive:         false,
				PlayerOrTeamID:   "29576970",
				FreshBlood:       false,
				LeaguePoints:     0,
			},
		},
	}

	positionsJSON = []byte(`[
		{
			"queueType": "RANKED_FLEX_SR",
			"hotStreak": false,
			"wins": 106,
			"veteran": false,
			"losses": 49,
			"playerOrTeamId": "56236881",
			"leagueName": "Dr. Mundo's Scouts",
			"playerOrTeamName": "Kaos Malthael",
			"inactive": false,
			"rank": "I",
			"freshBlood": false,
			"leagueId": "2346e99c-d5ac-3a9b-b3f4-b9b683a9b524",
			"tier": "CHALLENGER",
			"leaguePoints": 431
		},
		{
			"queueType": "RANKED_SOLO_5x5",
			"hotStreak": true,
			"wins": 113,
			"veteran": false,
			"losses": 58,
			"playerOrTeamId": "56236881",
			"leagueName": "Leona's Spellswords",
			"playerOrTeamName": "Kaos Malthael",
			"inactive": false,
			"rank": "I",
			"freshBlood": true,
			"leagueId": "40ac4f31-647b-3960-b95e-45a8ee1ef734",
			"tier": "CHALLENGER",
			"leaguePoints": 495
		}
	]`)

	wantPositions = []LeaguePositionDTO{
		{
			QueueType:        "RANKED_FLEX_SR",
			HotStreak:        false,
			Wins:             106,
			Veteran:          false,
			Losses:           49,
			PlayerOrTeamID:   "56236881",
			LeagueName:       "Dr. Mundo's Scouts",
			PlayerOrTeamName: "Kaos Malthael",
			Inactive:         false,
			Rank:             "I",
			FreshBlood:       false,
			LeagueID:         "2346e99c-d5ac-3a9b-b3f4-b9b683a9b524",
			Tier:             "CHALLENGER",
			LeaguePoints:     431,
		},
		{
			QueueType:        "RANKED_SOLO_5x5",
			HotStreak:        true,
			Wins:             113,
			Veteran:          false,
			Losses:           58,
			PlayerOrTeamID:   "56236881",
			LeagueName:       "Leona's Spellswords",
			PlayerOrTeamName: "Kaos Malthael",
			Inactive:         false,
			Rank:             "I",
			FreshBlood:       true,
			LeagueID:         "40ac4f31-647b-3960-b95e-45a8ee1ef734",
			Tier:             "CHALLENGER",
			LeaguePoints:     495,
		},
	}
)
