package ionia

import (
	"net/http"
	"reflect"
	"testing"
)

func TestChampionMasteryBySummonerID(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/champion-mastery/v3/champion-masteries/by-summoner/123", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(championsMasteryJSON)
	})

	got, _, err := client.ChampionMastery.MasteryBySummonerID(123)
	if err != nil {
		t.Errorf("ChampionMastery.MasteryBySummonerID returned error: %v", err)
	}
	if want := wantChampionsMastery; !reflect.DeepEqual(got, want) {
		t.Errorf("ChampionMastery.MasteryBySummonerID = %+v, want %+v", got, want)
	}
}

func TestChampionMasteryBySummonerAndChampionID(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/champion-mastery/v3/champion-masteries/by-summoner/123/by-champion/17", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(championMasteryJSON)
	})

	got, _, err := client.ChampionMastery.BySummonerAndChampionID(123, 17)
	if err != nil {
		t.Errorf("ChampionMastery.BySummonerAndChampionID returned error: %v", err)
	}
	if want := wantChampionMastery; !reflect.DeepEqual(got, want) {
		t.Errorf("ChampionMastery.BySummonerAndChampionID = %+v, want %+v", got, want)
	}
}

func TestChampionMasteryScoreBySummonerID(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/champion-mastery/v3/scores/by-summoner/123", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`327`))
	})

	got, _, err := client.ChampionMastery.ScoreBySummonerID(123)
	if err != nil {
		t.Errorf("ChampionMastery.ScoreBySummonerID return error: %v", err)
	}
	if got != 327 {
		t.Errorf("ChampionMastery.ScoreBySummonerID = %d, want 327", got)
	}
}

var (
	championsMasteryJSON = []byte(`[
		{
			"championLevel": 7,
			"chestGranted": false,
			"championPoints": 240114,
			"championPointsSinceLastLevel": 218514,
			"playerId": 27964133,
			"championPointsUntilNextLevel": 0,
			"tokensEarned": 0,
			"championId": 17,
			"lastPlayTime": 1519597116000
		},
		{
			"championLevel": 7,
        	"chestGranted": true,
        	"championPoints": 49784,
        	"championPointsSinceLastLevel": 28184,
        	"playerId": 27964133,
        	"championPointsUntilNextLevel": 0,
        	"tokensEarned": 0,
        	"championId": 82,
        	"lastPlayTime": 1519580558000
		},
		{
			"championLevel": 7,
        	"chestGranted": true,
        	"championPoints": 47663,
        	"championPointsSinceLastLevel": 26063,
        	"playerId": 27964133,
        	"championPointsUntilNextLevel": 0,
        	"tokensEarned": 0,
        	"championId": 43,
        	"lastPlayTime": 1519857361000
		}
	]`)

	wantChampionsMastery = []ChampionMasteryDTO{
		{
			ChampionLevel:                7,
			ChestGranted:                 false,
			ChampionPoints:               240114,
			ChampionPointsSinceLastLevel: 218514,
			PlayerID:                     27964133,
			ChampionPointsUntilNextLevel: 0,
			TokensEarned:                 0,
			ChampionID:                   17,
			LastPlayTime:                 1519597116000,
		},
		{
			ChampionLevel:                7,
			ChestGranted:                 true,
			ChampionPoints:               49784,
			ChampionPointsSinceLastLevel: 28184,
			PlayerID:                     27964133,
			ChampionPointsUntilNextLevel: 0,
			TokensEarned:                 0,
			ChampionID:                   82,
			LastPlayTime:                 1519580558000,
		},
		{
			ChampionLevel:                7,
			ChestGranted:                 true,
			ChampionPoints:               47663,
			ChampionPointsSinceLastLevel: 26063,
			PlayerID:                     27964133,
			ChampionPointsUntilNextLevel: 0,
			TokensEarned:                 0,
			ChampionID:                   43,
			LastPlayTime:                 1519857361000,
		},
	}

	championMasteryJSON = []byte(`{
		"championLevel": 7,
		"chestGranted": false,
		"championPoints": 240114,
		"championPointsSinceLastLevel": 218514,
		"playerId": 27964133,
		"championPointsUntilNextLevel": 0,
		"tokensEarned": 0,
		"championId": 17,
		"lastPlayTime": 1519597116000
	}`)

	wantChampionMastery = &ChampionMasteryDTO{
		ChampionLevel:                7,
		ChestGranted:                 false,
		ChampionPoints:               240114,
		ChampionPointsSinceLastLevel: 218514,
		PlayerID:                     27964133,
		ChampionPointsUntilNextLevel: 0,
		TokensEarned:                 0,
		ChampionID:                   17,
		LastPlayTime:                 1519597116000,
	}
)
