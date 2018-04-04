package ionia

import (
	"net/http"
	"reflect"
	"testing"
)

func TestStaticDataChampions(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/static-data/v3/champions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(staticChampionsJSON)
	})

	got, _, err := client.StaticData.Champions()

	if err != nil {
		t.Errorf("StaticData.AllChampions returned error: %v", err)
	}
	if want := wantStaticChampions; !reflect.DeepEqual(got, want) {
		t.Errorf("StaticData.AllChampions = %+v, want %+v", got, want)
	}
}

func TestStaticDataChampionByID(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/static-data/v3/champions/17", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(staticChampionJSON)
	})

	got, _, err := client.StaticData.ChampionByID(17)

	if err != nil {
		t.Errorf("StaticData.ChampionByID returned error: %v", err)
	}
	if want := wantStaticChampion; !reflect.DeepEqual(got, want) {
		t.Errorf("StaticData.ChampionByID = %+v, want %+v", got, want)
	}
}

func TestStaticDataItems(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/static-data/v3/items", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(staticItemsJSON)
	})

	got, _, err := client.StaticData.Items()

	if err != nil {
		t.Errorf("StaticData.Items returned error: %v", err)
	}
	if want := wantStaticItems; !reflect.DeepEqual(got, want) {
		t.Errorf("StaticData.Items = %+v, want %+v", got, want)
	}
}

func TestStaticDataItemByID(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/static-data/v3/items/123", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(staticItemJSON)
	})

	got, _, err := client.StaticData.ItemByID(123)

	if err != nil {
		t.Errorf("StaticData.ItemByID returned error: %v", err)
	}
	if want := wantStaticItem; !reflect.DeepEqual(got, want) {
		t.Errorf("StaticData.ItemByID = %+v, want %+v", got, want)
	}
}

func TestStaticDataLanguageStrings(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/static-data/v3/language-strings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(staticLanguageStringsJSON)
	})

	got, _, err := client.StaticData.LanguageStrings()

	if err != nil {
		t.Errorf("StaticData.LanguageStrings returned error: %v", err)
	}
	if want := wantStaticLanguageStrings; !reflect.DeepEqual(got, want) {
		t.Errorf("StaticData.LanguageStrings = %+v, want %+v", got, want)
	}
}

func TestStaticDataLanguages(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/static-data/v3/languages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(languagesJSON)
	})

	got, _, err := client.StaticData.Languages()

	if err != nil {
		t.Errorf("StaticData.Languages returned error: %v", err)
	}
	if want := wantLanguages; !reflect.DeepEqual(got, want) {
		t.Errorf("StaticData.Languages = %+v, want %+v", got, want)
	}
}

func TestStaticDataMaps(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/static-data/v3/maps", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mapsJSON)
	})

	got, _, err := client.StaticData.Maps()

	if err != nil {
		t.Errorf("StaticData.Maps returned error: %v", err)
	}
	if want := wantMaps; !reflect.DeepEqual(got, want) {
		t.Errorf("StaticData.Maps = %+v, want %+v", got, want)
	}
}

func TestStaticDataMasteries(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/lol/static-data/v3/masteries", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(masteriesJSON)
	})

	got, _, err := client.StaticData.Masteries()

	if err != nil {
		t.Errorf("StaticData.Masteries returned error: %v", err)
	}
	if want := wantMasteries; !reflect.DeepEqual(got, want) {
		t.Errorf("StaticData.Masteries = %+v, want %+v", got, want)
	}
}

var (
	staticChampionsJSON = []byte(`{
		"type": "champion",
		"version": "8.6.1",
		"data": {
			"MonkeyKing": {
				"title": "the Monkey King",
				"id": 62,
				"key": "MonkeyKing",
				"name": "Wukong"
			},
			"Jax": {
				"title": "Grandmaster at Arms",
				"id": 24,
				"key": "Jax",
				"name": "Jax"
			}
		}
	}`)

	wantStaticChampions = &StaticChampionListDTO{
		Type:    "champion",
		Version: "8.6.1",
		Data: map[string]StaticChampionDTO{
			"MonkeyKing": StaticChampionDTO{
				Title: "the Monkey King",
				ID:    62,
				Key:   "MonkeyKing",
				Name:  "Wukong",
			},
			"Jax": StaticChampionDTO{
				Title: "Grandmaster at Arms",
				ID:    24,
				Key:   "Jax",
				Name:  "Jax",
			},
		},
	}

	staticChampionsQueryJSON = []byte(`{
		"type": "champion",
		"version": "8.6.0",
		"data": {
			"MonkeyKing": {
				"name": "Wukong",
				"title": "the Monkey King",
				"allytips": [
					"Decoy and Nimbus Strike work well together to quickly strike your enemy and get out before they can retaliate.",
					"Try using Decoy near brush to make an enemy overreact to your movement."
				],
				"key": "MonkeyKing",
				"id": 62,
				"blurb": "Wukong is a vastayan trickster who uses his strength, agility, and intelligence to confuse his opponents and gain the upper hand. After finding a lifelong friend in the warrior known as Master Yi, Wukong became the last student of the ancient martial..."
			},
			"Jax": {
				"name": "Jax",
				"title": "Grandmaster at Arms",
				"allytips": [
					"Jax can Leap Strike to friendly units, including wards. You can use them to plan your escape.",
					"Jax benefits greatly from items that have both Ability Power and Attack Damage such as Guinsoo's Rageblade and Hextech Gunblade."
				],
				"key": "Jax",
				"id": 24,
				"blurb": "Unmatched in both his skill with unique armaments and his biting sarcasm, Jax is the last known weapons master of Icathia. After his homeland was laid low by its own hubris in unleashing the Void, Jax and his kind vowed to protect what little remained..."
			}
		}
	}`)

	wantStaticChampionsQuery = &StaticChampionListDTO{
		Type:    "champion",
		Version: "8.6.0",
		Data: map[string]StaticChampionDTO{
			"MonkeyKing": StaticChampionDTO{
				Name:  "Wukong",
				Title: "the Monkey King",
				AllyTips: []string{
					"Decoy and Nimbus Strike work well together to quickly strike your enemy and get out before they can retaliate.",
					"Try using Decoy near brush to make an enemy overreact to your movement.",
				},
				Key:   "MonkeyKing",
				ID:    62,
				Blurb: "Wukong is a vastayan trickster who uses his strength, agility, and intelligence to confuse his opponents and gain the upper hand. After finding a lifelong friend in the warrior known as Master Yi, Wukong became the last student of the ancient martial...",
			},
			"Jax": StaticChampionDTO{
				Name:  "Jax",
				Title: "Grandmaster at Arms",
				AllyTips: []string{
					"Jax can Leap Strike to friendly units, including wards. You can use them to plan your escape.",
					"Jax benefits greatly from items that have both Ability Power and Attack Damage such as Guinsoo's Rageblade and Hextech Gunblade.",
				},
				Key:   "Jax",
				ID:    24,
				Blurb: "Unmatched in both his skill with unique armaments and his biting sarcasm, Jax is the last known weapons master of Icathia. After his homeland was laid low by its own hubris in unleashing the Void, Jax and his kind vowed to protect what little remained...",
			},
		},
	}

	staticChampionJSON = []byte(`{
		"title": "the Swift Scout",
		"name": "Teemo",
		"key": "Teemo",
		"id": 17
	}`)

	wantStaticChampion = &StaticChampionDTO{
		Title: "the Swift Scout",
		Name:  "Teemo",
		Key:   "Teemo",
		ID:    17,
	}

	staticItemsJSON = []byte(`{
		"type": "item",
		"version": "8.7.1",
		"data": {
			"1001": {
				"plaintext": "Slightly increases Movement Speed",
				"description": "<groupLimit>Limited to 1 pair of boots.</groupLimit><br><br><unique>UNIQUE Passive - Enhanced Movement:</unique> +25 Movement Speed",
				"id": 1001,
				"name": "Boots of Speed"
			},
			"1004": {
				"plaintext": "Slightly increases Mana Regen",
				"description": "<stats><mana>+25% Base Mana Regen </mana></stats>",
				"id": 1004,
				"name": "Faerie Charm"
			},
			"1006": {
				"plaintext": "Slightly increases Health Regen",
				"description": "<stats>+50% Base Health Regen </stats>",
				"id": 1006,
				"name": "Rejuvenation Bead"
			}
		}
	}`)

	wantStaticItems = &ItemListDTO{
		Type:    "item",
		Version: "8.7.1",
		Data: map[string]ItemDTO{
			"1001": ItemDTO{
				PlainText:   "Slightly increases Movement Speed",
				Description: "<groupLimit>Limited to 1 pair of boots.</groupLimit><br><br><unique>UNIQUE Passive - Enhanced Movement:</unique> +25 Movement Speed",
				ID:          1001,
				Name:        "Boots of Speed",
			},
			"1004": ItemDTO{
				PlainText:   "Slightly increases Mana Regen",
				Description: "<stats><mana>+25% Base Mana Regen </mana></stats>",
				ID:          1004,
				Name:        "Faerie Charm",
			},
			"1006": ItemDTO{
				PlainText:   "Slightly increases Health Regen",
				Description: "<stats>+50% Base Health Regen </stats>",
				ID:          1006,
				Name:        "Rejuvenation Bead",
			},
		},
	}

	staticItemJSON = []byte(`{
		"plainText": "Slightly increases Movement Speed",
		"description": "<groupLimit>Limited to 1 pair of boots.</groupLimit><br><br><unique>UNIQUE Passive - Enhanced Movement:</unique> +25 Movement Speed",
		"id": 1001,
		"name": "Boots of Speed"
	}`)

	wantStaticItem = &ItemDTO{
		PlainText:   "Slightly increases Movement Speed",
		Description: "<groupLimit>Limited to 1 pair of boots.</groupLimit><br><br><unique>UNIQUE Passive - Enhanced Movement:</unique> +25 Movement Speed",
		ID:          1001,
		Name:        "Boots of Speed",
	}

	staticLanguageStringsJSON = []byte(`{
		"data": {
			"rPercentMagicPenetrationModPerLevel": "Magic Pen. % at level 18",
			"rFlatEnergyModPerLevel": "Energy at level 18",
			"colloq_ManaRegen": ";mpregen;mp5"
		}
	}`)

	wantStaticLanguageStrings = &LanguageStringsDTO{
		Data: map[string]string{
			"rPercentMagicPenetrationModPerLevel": "Magic Pen. % at level 18",
			"rFlatEnergyModPerLevel":              "Energy at level 18",
			"colloq_ManaRegen":                    ";mpregen;mp5",
		},
	}

	languagesJSON = []byte(`[
		"en_US",
    	"cs_CZ",
    	"de_DE",
    	"el_GR",
    	"en_AU",
    	"en_GB",
    	"en_PH",
    	"en_SG",
    	"es_AR",
    	"es_ES",
    	"es_MX",
    	"fr_FR",
    	"hu_HU",
    	"id_ID",
    	"it_IT",
    	"ja_JP",
    	"ko_KR",
    	"ms_MY",
    	"pl_PL",
    	"pt_BR",
    	"ro_RO",
    	"ru_RU",
    	"th_TH",
    	"tr_TR",
    	"vn_VN",
    	"zh_CN",
    	"zh_MY",
    	"zh_TW"
	]`)

	wantLanguages = []string{
		"en_US",
		"cs_CZ",
		"de_DE",
		"el_GR",
		"en_AU",
		"en_GB",
		"en_PH",
		"en_SG",
		"es_AR",
		"es_ES",
		"es_MX",
		"fr_FR",
		"hu_HU",
		"id_ID",
		"it_IT",
		"ja_JP",
		"ko_KR",
		"ms_MY",
		"pl_PL",
		"pt_BR",
		"ro_RO",
		"ru_RU",
		"th_TH",
		"tr_TR",
		"vn_VN",
		"zh_CN",
		"zh_MY",
		"zh_TW",
	}

	mapsJSON = []byte(`{
		"data": {
			"10": {
				"mapName": "The Twisted Treeline",
				"image": {
					"full": "map10.png",
					"group": "map",
					"sprite": "map0.png",
					"h": 48,
					"w": 48,
					"y": 0,
					"x": 0
				},
				"mapId": 10
			},
			"11": {
				"mapName": "Summoner's Rift",
				"image": {
					"full": "map11.png",
					"group": "map",
					"sprite": "map0.png",
					"h": 48,
					"w": 48,
					"y": 0,
					"x": 48
				},
				"mapId": 11
			}
		}
	}`)

	wantMaps = &MapDataDTO{
		Data: map[string]MapDetailsDTO{
			"10": MapDetailsDTO{
				MapName: "The Twisted Treeline",
				Image: ImageDTO{
					Full:   "map10.png",
					Group:  "map",
					Sprite: "map0.png",
					H:      48,
					W:      48,
					Y:      0,
					X:      0,
				},
				MapID: 10,
			},
			"11": MapDetailsDTO{
				MapName: "Summoner's Rift",
				Image: ImageDTO{
					Full:   "map11.png",
					Group:  "map",
					Sprite: "map0.png",
					H:      48,
					W:      48,
					Y:      0,
					X:      48,
				},
				MapID: 11,
			},
		},
	}

	masteriesJSON = []byte(`{
		"type": "mastery",
		"version": "7.23.1",
		"data": {
			"6111": {
				"description": [
					"+0.8% Attack Speed",
					"+1.6% Attack Speed",
					"+2.4% Attack Speed",
					"+3.2% Attack Speed",
					"+4% Attack Speed"
				],
				"id": 6111,
				"name": "Fury"
			},
			"6114": {
				"description": [
					"+0.4% increased Ability damage",
					"+0.8% increased Ability damage",
					"+1.2% increased Ability damage",
					"+1.6% increased Ability damage",
					"+2.0% increased Ability damage"
				],
				"id": 6114,
				"name": "Sorcery"
			}
		}
	}`)

	wantMasteries = &MasteryListDTO{
		Type:    "mastery",
		Version: "7.23.1",
		Data: map[string]MasteryDTO{
			"6111": MasteryDTO{
				Description: []string{
					"+0.8% Attack Speed",
					"+1.6% Attack Speed",
					"+2.4% Attack Speed",
					"+3.2% Attack Speed",
					"+4% Attack Speed",
				},
				ID:   6111,
				Name: "Fury",
			},
			"6114": MasteryDTO{
				Description: []string{
					"+0.4% increased Ability damage",
					"+0.8% increased Ability damage",
					"+1.2% increased Ability damage",
					"+1.6% increased Ability damage",
					"+2.0% increased Ability damage",
				},
				ID:   6114,
				Name: "Sorcery",
			},
		},
	}
)
