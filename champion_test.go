package ionia

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAllChampions(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/platform/v3/champions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(championsJSON)
	})

	got, _, err := client.Champion.All()
	if err != nil {
		t.Errorf("Champion.All returned error: %v", err)
	}
	if want := wantChampions; !reflect.DeepEqual(got, want) {
		t.Errorf("Champion.All = %+v, want %+v", got, want)
	}
}

func TestChampionByID(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/platform/v3/champions/123", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(championJSON)
	})

	got, _, err := client.Champion.ByID(123)
	if err != nil {
		t.Errorf("Champion.ByID returned error: %v", err)
	}
	if want := wantChampion; !reflect.DeepEqual(got, want) {
		t.Errorf("Champion.ByID = %+v, want %+v", got, want)
	}
}

var (
	championsJSON = []byte(`{
		"champions": [
			{
				"rankedPlayEnabled": true,
				"botEnabled": true,
				"botMmEnabled": true,
				"active": true,
				"freeToPlay": false,
				"id": 123
			},
			{
				"rankedPlayEnabled": false,
				"botEnabled": false,
				"botMmEnabled": false,
				"active": false,
				"freeToPlay": false,
				"id": 200
			}
		]
	}`)

	wantChampions = &ChampionListDTO{
		Champions: []ChampionDTO{
			{
				RankedPlayEnabled: true,
				BotEnabled:        true,
				BotMmEnabled:      true,
				Active:            true,
				FreeToPlay:        false,
				ID:                123,
			},
			{
				RankedPlayEnabled: false,
				BotEnabled:        false,
				BotMmEnabled:      false,
				Active:            false,
				FreeToPlay:        false,
				ID:                200,
			},
		},
	}

	championJSON = []byte(`{
		"rankedPlayEnabled": true,
		"botEnabled": true,
		"botMmEnabled": true,
		"active": true,
		"freeToPlay": false,
		"id": 123
	}`)

	wantChampion = &ChampionDTO{
		RankedPlayEnabled: true,
		BotEnabled:        true,
		BotMmEnabled:      true,
		Active:            true,
		FreeToPlay:        false,
		ID:                123,
	}
)
