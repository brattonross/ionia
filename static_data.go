package ionia

import (
	"net/http"
	"strconv"
)

// StaticDataService represents the LOL-Static-Data-V3 API methods.
// https://developer.riotgames.com/api-methods/#lol-static-data-v3
type StaticDataService service

// StaticChampionListDTO contains champion list data.
type StaticChampionListDTO struct {
	Keys    map[string]string            `json:"keys"`
	Data    map[string]StaticChampionDTO `json:"data"`
	Version string                       `json:"version"`
	Type    string                       `json:"type"`
	Format  string                       `json:"format"`
}

// StaticChampionDTO contains champion data.
type StaticChampionDTO struct {
	Info        InfoDTO            `json:"info"`
	EnemyTips   []string           `json:"enemytips"`
	Stats       StatsDTO           `json:"stats"`
	Name        string             `json:"name"`
	Title       string             `json:"title"`
	Image       ImageDTO           `json:"image"`
	Tags        []string           `json:"tags"`
	Partype     string             `json:"partype"`
	Skins       []SkinDTO          `json:"skin"`
	Passive     PassiveDTO         `json:"passive"`
	Recommended RecommendedDTO     `json:"recommended"`
	AllyTips    []string           `json:"allytips"`
	Key         string             `json:"key"`
	Lore        string             `json:"lore"`
	ID          int                `json:"id"`
	Blurb       string             `json:"blurb"`
	Spells      []ChampionSpellDTO `json:"spells"`
}

// InfoDTO contains champion information.
type InfoDTO struct {
	Difficulty int `json:"difficulty"`
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
}

// StatsDTO contains champion stats data.
type StatsDTO struct {
	ArmorPerLevel        float64 `json:"armorperlevel"`
	HPPerLevel           float64 `json:"hpperlevel"`
	AttackDamage         float64 `json:"attackdamage"`
	MPPerLevel           float64 `json:"mpperlevel"`
	AttackSpeedOffset    float64 `json:"attackspeedoffset"`
	Armor                float64 `json:"armor"`
	HP                   float64 `json:"hp"`
	HPRegenPerLevel      float64 `json:"hpregenperlevel"`
	SpellBlock           float64 `json:"spellblock"`
	AttackRange          float64 `json:"attackrange"`
	MoveSpeed            float64 `json:"movespeed"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
	MPRegenPerLevel      float64 `json:"mpregenperlevel"`
	MP                   float64 `json:"mp"`
	SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
	Crit                 float64 `json:"crit"`
	MPRegen              float64 `json:"mpregen"`
	AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
	HPRegen              float64 `json:"hpregen"`
	CritPerLevel         float64 `json:"critperlevel"`
}

// SkinDTO contains champion skin data.
type SkinDTO struct {
	Num  int    `json:"num"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// PassiveDTO contains champion passive data.
type PassiveDTO struct {
	Image                ImageDTO `json:"image"`
	SanitizedDescription string   `json:"sanitizedDescription"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
}

// RecommendedDTO contains champion recommended data.
type RecommendedDTO struct {
	Map      string     `json:"map"`
	Blocks   []BlockDTO `json:"blocks"`
	Champion string     `json:"champion"`
	Title    string     `json:"title"`
	Priority bool       `json:"priority"`
	Mode     string     `json:"mode"`
	Type     string     `json:"type"`
}

// BlockDTO contains champion recommended block data.
type BlockDTO struct {
	Items   []BlockItemDTO `json:"items"`
	RecMath bool           `json:"recMath"`
	Type    string         `json:"type"`
}

// BlockItemDTO contains champion recommended block item data.
type BlockItemDTO struct {
	Count int `json:"count"`
	ID    int `json:"id"`
}

// ChampionSpellDTO contains champion spell data.
type ChampionSpellDTO struct {
	CooldownBurn         string         `json:"cooldownBurn"`
	Resource             string         `json:"resource"`
	LevelTip             LevelTipDTO    `json:"levelTip"`
	Vars                 []SpellVarsDTO `json:"vars"`
	CostType             string         `json:"costType"`
	Image                ImageDTO       `json:"image"`
	SanitizedDescription string         `json:"sanitizedDescription"`
	SanitizedTooltip     string         `json:"sanitizedTooltip"`
	Effect               [][]float64    `json:"effect"`
	Tooltip              string         `json:"tooltip"`
	MaxRank              int            `json:"maxrank"`
	CostBurn             string         `json:"costBurn"`
	RangeBurn            string         `json:"rangeBurn"`
	Range                interface{}    `json:"range"` // Either []int or string ("self")
	Cooldown             []float64      `json:"cooldown"`
	Cost                 []int          `json:"cost"`
	Key                  string         `json:"key"`
	Description          string         `json:"description"`
	EffectBurn           []string       `json:"effectBurn"`
	AltImages            []ImageDTO     `json:"altimages"`
	Name                 string         `json:"name"`
}

// ImageDTO contains image data.
type ImageDTO struct {
	Full   string `json:"full"`
	Group  string `json:"group"`
	Sprite string `json:"sprite"`
	H      int    `json:"h"`
	W      int    `json:"w"`
	Y      int    `json:"y"`
	X      int    `json:"x"`
}

// LevelTipDTO contains champion level tip data.
type LevelTipDTO struct {
	Effect []string `json:"effect"`
	Label  []string `json:"label"`
}

// SpellVarsDTO contains spell vars data.
type SpellVarsDTO struct {
	RanksWith string    `json:"ranksWith"`
	Dyn       string    `json:"dyn"`
	Link      string    `json:"link"`
	Coeff     []float64 `json:"coeff"`
	Key       string    `json:"key"`
}

// StaticDataChampionsOptions specifies the optional parameters to the Static Data champions service methods.
type StaticDataChampionsOptions struct {
	// Locale of the data to be returned.
	// Default: Locale of the client's region.
	Locale string `url:"locale,omitempty"`

	// Patch version for the data to be returned.
	// A list of valid versions can be obtained from the /versions endpoint.
	// Default: Latest version for the client's region.
	Version string `url:"version,omitempty"`

	// Tags to return additional data.
	// To return all additional data, use the tag 'all'.
	// Default: type, version, data, id, key, name, and title are returned.
	ChampionListData []string `url:"champListData,omitempty"`

	// Tags to return additional data.
	// To return all additional data, use the tag 'all'.
	// Default: type, version, data, id, key, name, and title are returned.
	Tags []string `url:"tags,omitempty"`

	// If true, the returned data will use the champions' IDs as keys.
	// Otherwise, the champions' keys will be used instead.
	// Default: false.
	DataByID bool `url:"dataById,omitempty"`
}

// StaticDataChampionsOption is a function which modifies the StaticDataChampionsOptions.
type StaticDataChampionsOption func(*StaticDataChampionsOptions)

// Champions retrieves a list of champion information.
func (s *StaticDataService) Champions(opts ...StaticDataChampionsOption) (*StaticChampionListDTO, *http.Response, error) {
	options := &StaticDataChampionsOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/champions", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, nil, err
	}

	cl := &StaticChampionListDTO{}
	resp, err := s.client.Do(req, cl)
	if err != nil {
		return nil, resp, err
	}

	return cl, resp, nil
}

// ChampionByID gets champion information by champion ID.
func (s *StaticDataService) ChampionByID(championID int64, opts ...StaticDataChampionsOption) (*StaticChampionDTO, *http.Response, error) {
	options := &StaticDataChampionsOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/champions/"+strconv.FormatInt(championID, 10), options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, nil, err
	}

	c := &StaticChampionDTO{}
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// ItemListDTO contains item list data.
type ItemListDTO struct {
	Data    map[string]ItemDTO `json:"data"`
	Version string             `json:"version"`
	Tree    []ItemTreeDTO      `json:"tree"`
	Groups  []GroupDTO         `json:"groups"`
	Type    string             `json:"type"`
}

// ItemTreeDTO contains item tree data.
type ItemTreeDTO struct {
	Header string   `json:"header"`
	Tags   []string `json:"tags"`
}

// ItemDTO contains item data.
type ItemDTO struct {
	Gold                 GoldDTO               `json:"gold"`
	PlainText            string                `json:"plaintext"`
	HideFromAll          bool                  `json:"hideFromAll"`
	InStore              bool                  `json:"inStore"`
	Into                 []string              `json:"into"`
	ID                   int                   `json:"id"`
	Stats                InventoryDataStatsDTO `json:"stats"`
	Colloq               string                `json:"colloq"`
	Maps                 map[string]bool       `json:"maps"`
	SpecialRecipe        int                   `json:"specialRecipe"`
	Image                ImageDTO              `json:"image"`
	Description          string                `json:"description"`
	Tags                 []string              `json:"tags"`
	Effect               map[string]string     `json:"effect"`
	RequiredChampion     string                `json:"requiredChampion"`
	From                 []string              `json:"from"`
	Group                string                `json:"group"`
	ConsumeOnFull        bool                  `json:"consumeOnFull"`
	Name                 string                `json:"name"`
	Consumed             bool                  `json:"consumed"`
	SanitizedDescription string                `json:"sanitizedDescription"`
	Depth                int                   `json:"depth"`
	Stacks               int                   `json:"stacks"`
}

// GoldDTO contains item gold data.
type GoldDTO struct {
	Sell        int  `json:"sell"`
	Total       int  `json:"total"`
	Base        int  `json:"base"`
	Purchasable bool `json:"purchasable"`
}

// InventoryDataStatsDTO contains stats for inventory (e.g. runes and items).
type InventoryDataStatsDTO struct {
	PercentCritDamageMod     float64 `json:"PercentCritDamageMod"`
	PercentSpellBlockMod     float64 `json:"PercentSpellBlockMod"`
	PercentHPRegenMod        float64 `json:"PercentHPRegenMod"`
	PercentMovementSpeedMod  float64 `json:"PercentMovementSpeedMod"`
	FlatSpellBlockMod        float64 `json:"FlatSpellBlockMod"`
	FlatCritDamageMod        float64 `json:"FlatCritDamageMod"`
	FlatEnergyPoolMod        float64 `json:"FlatEnergyPoolMod"`
	PercentLifeStealMod      float64 `json:"PercentLifeStealMod"`
	FlatMPPoolMod            float64 `json:"FlatMPPoolMod"`
	FlatMovementSpeedMod     float64 `json:"FlatMovementSpeedMod"`
	PercentAttackSpeedMod    float64 `json:"PercentAttackSpeedMod"`
	FlatBlockMod             float64 `json:"FlatBlockMod"`
	PercentBlockMod          float64 `json:"PercentBlockMod"`
	FlatEnergyRegenMod       float64 `json:"FlatEnergyRegenMod"`
	PercentSpellVampMod      float64 `json:"PercentSpellVampMod"`
	FlatMPRegenMod           float64 `json:"FlatMPRegenMod"`
	PercentDodgeMod          float64 `json:"PercentDodgeMod"`
	FlatAttackSpeedMod       float64 `json:"FlatAttackSpeedMod"`
	FlatArmorMod             float64 `json:"FlatArmorMod"`
	FlatHPRegenMod           float64 `json:"FlatHPRegenMod"`
	PercentMagicDamageMod    float64 `json:"PercentMagicDamageMod"`
	PercentMPPoolMod         float64 `json:"PercentMPPoolMod"`
	FlatMagicDamageMod       float64 `json:"FlatMagicDamageMod"`
	PercentMPRegenMod        float64 `json:"PercentMPRegenMod"`
	PercentPhysicalDamageMod float64 `json:"PercentPhysicalDamageMod"`
	FlatPhysicalDamageMod    float64 `json:"FlatPhysicalDamageMod"`
	PercentHPPoolMod         float64 `json:"PercentHPPoolMod"`
	PercentArmorMod          float64 `json:"PercentArmorMod"`
	PercentCritChanceMod     float64 `json:"PercentCritChanceMod"`
	PercentEXPBonus          float64 `json:"PercentEXPBonus"`
	FlatHPPoolMod            float64 `json:"FlatHPPoolMod"`
	FlatCritChanceMod        float64 `json:"FlatCritChanceMod"`
	FlatEXPBonus             float64 `json:"FlatEXPBonus"`
}

// GroupDTO contains item group data.
type GroupDTO struct {
	MaxGroupOwnable string `json:"MaxGroupOwnable"`
	Key             string `json:"key"`
}

// StaticDataItemsOptions specifies the optional parameters to the Static Data items service methods.
type StaticDataItemsOptions struct {
	// Locale of the data to be returned.
	// Default: Locale of the client's region.
	Locale string `url:"locale,omitempty"`

	// Patch version for the data to be returned.
	// A list of valid versions can be obtained from the /versions endpoint.
	// Default: Latest version for the client's region.
	Version string `url:"version,omitempty"`

	// Tags to return additional data.
	// To return all additional data, use the tag 'all'.
	// Default: type, version, data, id, name, description, plaintext, and group are returned.
	ItemListData []string `url:"itemListData,omitempty"`

	// Tags to return additional data.
	// To return all additional data, use the tag 'all'.
	// Default: type, version, data, id, name, description, plaintext, and group are returned.
	Tags []string `url:"tags,omitempty"`
}

// StaticDataItemsOption is a function which modifies the StaticDataItemsOptions.
type StaticDataItemsOption func(*StaticDataItemsOptions)

// Items retrieves a list of items.
func (s *StaticDataService) Items(opts ...StaticDataItemsOption) (*ItemListDTO, *http.Response, error) {
	options := &StaticDataItemsOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/items", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, nil, err
	}

	il := &ItemListDTO{}
	resp, err := s.client.Do(req, il)
	if err != nil {
		return nil, resp, err
	}

	return il, resp, nil
}

// ItemByID retrieves an item by ID.
func (s *StaticDataService) ItemByID(itemID int64, opts ...StaticDataItemsOption) (*ItemDTO, *http.Response, error) {
	options := &StaticDataItemsOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/items/"+strconv.FormatInt(itemID, 10), options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, nil, err
	}

	i := &ItemDTO{}
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// LanguageStringsDTO contains language strings data.
type LanguageStringsDTO struct {
	Data    map[string]string `json:"data"`
	Version string            `json:"version"`
	Type    string            `json:"type"`
}

// StaticDataLanguageStringsOptions specifies the optional parameters to the Static Data language strings service method.
type StaticDataLanguageStringsOptions struct {
	Locale  string `url:"locale"`
	Version string `url:"version"`
}

// StaticDataLanguageStringsOption is a function which modifies the StaticDataLanguageStringsOptions.
type StaticDataLanguageStringsOption func(*StaticDataLanguageStringsOptions)

// LanguageStrings retrieves language strings data.
func (s *StaticDataService) LanguageStrings(opts ...StaticDataLanguageStringsOption) (*LanguageStringsDTO, *http.Response, error) {
	options := &StaticDataLanguageStringsOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/language-strings", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, nil, err
	}

	ls := &LanguageStringsDTO{}
	resp, err := s.client.Do(req, ls)
	if err != nil {
		return nil, resp, err
	}

	return ls, resp, nil
}

// Languages retrieves supported languages data.
func (s *StaticDataService) Languages() ([]string, *http.Response, error) {
	req, err := s.client.NewRequest("lol/static-data/v3/languages")
	if err != nil {
		return nil, nil, err
	}

	langs := []string{}
	resp, err := s.client.Do(req, &langs)
	if err != nil {
		return nil, resp, err
	}

	return langs, resp, nil
}

// MapDataDTO contains map data.
type MapDataDTO struct {
	Data    map[string]MapDetailsDTO `json:"data"`
	Version string                   `json:"version"`
	Type    string                   `json:"type"`
}

// MapDetailsDTO contains map details data.
type MapDetailsDTO struct {
	MapName               string   `json:"mapName"`
	Image                 ImageDTO `json:"image"`
	MapID                 int64    `json:"mapId"`
	UnpurchasableItemList []int64  `json:"unpurchasableItemList"`
}

// StaticDataMapsOptions specifies the optional parameters to the Static Data maps service method.
type StaticDataMapsOptions struct {
	Locale  string `url:"locale"`
	Version string `url:"version"`
}

// StaticDataMapsOption is a function which modifies the StaticDataMapsOptions.
type StaticDataMapsOption func(*StaticDataMapsOptions)

// Maps retrieves map data.
func (s *StaticDataService) Maps(opts ...StaticDataMapsOption) (*MapDataDTO, *http.Response, error) {
	options := &StaticDataMapsOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/maps", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, nil, err
	}

	m := &MapDataDTO{}
	resp, err := s.client.Do(req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}

// MasteryListDTO contains mastery list data.
type MasteryListDTO struct {
	Data    map[string]MasteryDTO `json:"data"`
	Version string                `json:"version"`
	Tree    MasteryTreeDTO        `json:"tree"`
	Type    string                `json:"type"`
}

// MasteryTreeDTO contains mastery tree data.
type MasteryTreeDTO struct {
	Resolve  []MasteryTreeListDTO `json:"Resolve"`
	Defense  []MasteryTreeListDTO `json:"Defense"`
	Utility  []MasteryTreeListDTO `json:"Utility"`
	Offense  []MasteryTreeListDTO `json:"Offense"`
	Ferocity []MasteryTreeListDTO `json:"Ferocity"`
	Cunning  []MasteryTreeListDTO `json:"Cunning"`
}

// MasteryTreeListDTO contains mastery tree list data.
type MasteryTreeListDTO struct {
	MasteryTreeItems []MasteryTreeItemDTO `json:"masteryTreeItems"`
}

// MasteryTreeItemDTO contains mastery tree item data.
type MasteryTreeItemDTO struct {
	MasteryID int    `json:"masteryId"`
	Prereq    string `json:"prereq"`
}

// MasteryDTO contains mastery data.
type MasteryDTO struct {
	Prereq               string   `json:"prereq"`
	MasteryTree          string   `json:"masteryTree"`
	Name                 string   `json:"name"`
	Ranks                int      `json:"ranks"`
	Image                ImageDTO `json:"image"`
	SanitizedDescription []string `json:"sanitizedDescription"`
	ID                   int      `json:"id"`
	Description          []string `json:"description"`
}

// StaticDataMasteriesOptions specifies the optional parameters to the Static Data masteries service method.
type StaticDataMasteriesOptions struct {
	Locale          string   `url:"locale"`
	Version         string   `url:"version"`
	Tags            []string `url:"tags"`
	MasteryListData []string `url:"masteryListData"`
}

// StaticDataMasteriesOption is a function which modifies the StaticDataMasteriesOptions.
type StaticDataMasteriesOption func(*StaticDataMasteriesOptions)

// Masteries retrieves the list of masteries.
func (s *StaticDataService) Masteries(opts ...StaticDataMasteriesOption) (*MasteryListDTO, *http.Response, error) {
	options := &StaticDataMasteriesOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/masteries", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, nil, err
	}

	m := &MasteryListDTO{}
	resp, err := s.client.Do(req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}
