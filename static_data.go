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

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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
	req, err := s.client.NewRequest(http.MethodGet, "lol/static-data/v3/languages", nil)
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

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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

// StaticDataMasteryOptions specifies the optional parameters to the Static Data mastery service method.
type StaticDataMasteryOptions struct {
	MasteryData []string `url:"masteryData"`
	Locale      string   `url:"locale"`
	Version     string   `url:"version"`
	Tags        []string `url:"tags"`
}

// StaticDataMasteryOption is a function which modifies the StaticDataMasteryOptions.
type StaticDataMasteryOption func(*StaticDataMasteryOptions)

// MasteryByID retrieves a mastery by ID.
func (s *StaticDataService) MasteryByID(masteryID int64, opts ...StaticDataMasteryOption) (*MasteryDTO, *http.Response, error) {
	options := &StaticDataMasteryOptions{}
	for _, o := range opts {
		o(options)
	}

	u := "lol/static-data/v3/masteries/" + strconv.FormatInt(masteryID, 10)
	u, err := addOptions(u, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	m := &MasteryDTO{}
	resp, err := s.client.Do(req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}

// ProfileIconDataDTO contains profile icon data.
type ProfileIconDataDTO struct {
	Data    map[string]ProfileIconDetailsDTO `json:"data"`
	Version string                           `json:"version"`
	Type    string                           `json:"type"`
}

// ProfileIconDetailsDTO contains profile icon details data.
type ProfileIconDetailsDTO struct {
	Image ImageDTO `json:"image"`
	ID    int64    `json:"id"`
}

// StaticDataProfileIconsOptions specifies the optional parameters to the Static Data profile icons service method.
type StaticDataProfileIconsOptions struct {
	Locale  string `url:"locale"`
	Version string `url:"version"`
}

// StaticDataProfileIconsOption is a function which modifies the StaticDataProfileIconsOptions.
type StaticDataProfileIconsOption func(*StaticDataProfileIconsOptions)

// ProfileIcons retrieves profile icons.
func (s *StaticDataService) ProfileIcons(opts ...StaticDataProfileIconsOption) (*ProfileIconDataDTO, *http.Response, error) {
	options := &StaticDataProfileIconsOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/profile-icons", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	p := &ProfileIconDataDTO{}
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// RealmDTO contains realm data.
type RealmDTO struct {
	Lg             string            `json:"lg"`
	Dd             string            `json:"dd"`
	L              string            `json:"l"`
	N              map[string]string `json:"n"`
	ProfileIconMax int               `json:"profileiconmax"`
	Store          string            `json:"store"`
	V              string            `json:"v"`
	CDN            string            `json:"cdn"`
	CSS            string            `json:"css"`
}

// Realms retrieves realms data.
func (s *StaticDataService) Realms() (*RealmDTO, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "lol/static-data/v3/realms", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &RealmDTO{}
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// ReforgedRunePathDTO contains reforged rune path data.
type ReforgedRunePathDTO struct {
	Slots []ReforgedRuneSlotDTO `json:"slots"`
	Icon  string                `json:"icon"`
	ID    int                   `json:"id"`
	Key   string                `json:"key"`
	Name  string                `json:"name"`
}

// ReforgedRuneSlotDTO contains reforged rune slot data.
type ReforgedRuneSlotDTO struct {
	Runes []ReforgedRuneDTO `json:"runes"`
}

// ReforgedRuneDTO contains reforged rune data.
type ReforgedRuneDTO struct {
	RunePathName string `json:"runePathName"`
	RunePathID   int    `json:"runePathId"`
	Name         string `json:"name"`
	ID           int    `json:"id"`
	Key          string `json:"key"`
	ShortDesc    string `json:"shortDesc"`
	LongDesc     string `json:"longDesc"`
	Icon         string `json:"icon"`
}

// StaticDataReforgedRuneOptions specifies the optional parameters to the Static Data reforged runes service method.
type StaticDataReforgedRuneOptions struct {
	Version string `url:"version"`
	Locale  string `url:"locale"`
}

// StaticDataReforgedRuneOption is a function which modifies the StaticDataReforgedRuneOptions.
type StaticDataReforgedRuneOption func(*StaticDataReforgedRuneOptions)

// ReforgedRunePaths retrieves reforged rune paths.
func (s *StaticDataService) ReforgedRunePaths(opts ...StaticDataReforgedRuneOption) ([]ReforgedRunePathDTO, *http.Response, error) {
	options := &StaticDataReforgedRuneOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/reforged-rune-paths", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	p := []ReforgedRunePathDTO{}
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ReforgedRunePathByID retrieves a reforged rune path by ID.
func (s *StaticDataService) ReforgedRunePathByID(pathID int, opts ...StaticDataReforgedRuneOption) (*ReforgedRunePathDTO, *http.Response, error) {
	options := &StaticDataReforgedRuneOptions{}
	for _, o := range opts {
		o(options)
	}

	u := "lol/static-data/v3/reforged-rune-paths/" + strconv.Itoa(pathID)
	u, err := addOptions(u, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	p := &ReforgedRunePathDTO{}
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// ReforgedRunes retrieves reforged runes.
func (s *StaticDataService) ReforgedRunes(opts ...StaticDataReforgedRuneOption) ([]ReforgedRuneDTO, *http.Response, error) {
	options := &StaticDataReforgedRuneOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/reforged-runes", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := []ReforgedRuneDTO{}
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// ReforgedRuneByID retrieves a reforged rune by ID.
func (s *StaticDataService) ReforgedRuneByID(runeID int, opts ...StaticDataReforgedRuneOption) (*ReforgedRuneDTO, *http.Response, error) {
	options := &StaticDataReforgedRuneOptions{}
	for _, o := range opts {
		o(options)
	}

	u := "lol/static-data/v3/reforged-runes/" + strconv.Itoa(runeID)
	u, err := addOptions(u, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := &ReforgedRuneDTO{}
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// RuneListDTO contains rune list data.
type RuneListDTO struct {
	Data    map[string]RuneDTO `json:"data"`
	Version string             `json:"version"`
	Type    string             `json:"type"`
}

// RuneDTO contains rune data.
type RuneDTO struct {
	Stats                RuneStatsDTO `json:"stats"`
	Name                 string       `json:"name"`
	Tags                 []string     `json:"tags"`
	Image                ImageDTO     `json:"image"`
	SanitizedDescription string       `json:"sanitizedDescription"`
	Rune                 MetaDataDTO  `json:"rune"`
	ID                   int          `json:"id"`
	Description          string       `json:"description"`
}

// RuneStatsDTO contains stats for runes.
type RuneStatsDTO struct {
	PercentTimeDeadModPerLevel         int64
	PercentArmorPenetrationModPerLevel int64
	PercentCritDamageMod               int64
	PercentSpellBlockMod               int64
	PercentHPRegenMod                  int64
	PercentMovementSpeedMod            int64
	FlatSpellBlockMod                  int64
	FlatEnergyRegenModPerLevel         int64
	FlatEnergyPoolMod                  int64
	FlatMagicPenetrationModPerLevel    int64
	PercentLifeStealMod                int64
	FlatMPPoolMod                      int64
	PercentCooldownMod                 int64
	PercentMagicPenetrationMod         int64
	FlatArmorPenetrationModPerLevel    int64
	FlatMovementSpeedMod               int64
	FlatTimeDeadModPerLevel            int64
	FlatArmorModPerLevel               int64
	PercentAttackSpeedMod              int64
	FlatDodgeModPerLevel               int64
	PercentMagicDamageMod              int64
	PercentBlockMod                    int64
	FlatDodgeMod                       int64
	FlatEnergyRegenMod                 int64
	FlatHPModPerLevel                  int64
	PercentAttackSpeedModPerLevel      int64
	PercentSpellVampMod                int64
	FlatMPRegenMod                     int64
	PercentHPPoolMod                   int64
	PercentDodgeMod                    int64
	FlatAttackSpeedMod                 int64
	FlatArmorMod                       int64
	FlatMagicDamageModPerLevel         int64
	FlatHPRegenMod                     int64
	PercentPhysicalDamageMod           int64
	FlatCritChanceModPerLevel          int64
	FlatSpellBlockModPerLevel          int64
	PercentTimeDeadMod                 int64
	FlatBlockMod                       int64
	PercentMPPoolMod                   int64
	FlatMagicDamageMod                 int64
	PercentMPRegenMod                  int64
	PercentMovementSpeedModPerLevel    int64
	PercentCooldownModPerLevel         int64
	FlatMPModPerLevel                  int64
	FlatEnergyModPerLevel              int64
	FlatPhysicalDamageMod              int64
	FlatHPRegenModPerLevel             int64
	FlatCritDamageMod                  int64
	PercentArmorMod                    int64
	FlatMagicPenetrationMod            int64
	PercentCritChanceMod               int64
	FlatPhysicalDamageModPerLevel      int64
	PercentArmorPenetrationMod         int64
	PercentEXPBonus                    int64
	FlatMPRegenModPerLevel             int64
	PercentMagicPenetrationModPerLevel int64
	FlatTimeDeadMod                    int64
	FlatMovementSpeedModPerLevel       int64
	FlatGoldPer10Mod                   int64
	FlatArmorPenetrationMod            int64
	FlatCritDamageModPerLevel          int64
	FlatHPPoolMod                      int64
	FlatCritChanceMod                  int64
	FlatEXPBonus                       int64
}

// MetaDataDTO contains meta data.
type MetaDataDTO struct {
	Tier   string `json:"tier"`
	Type   string `json:"type"`
	IsRune bool   `json:"isRune"`
}

// StaticDataRuneListOptions specifies the optional parameters for the Static Data rune service method.
type StaticDataRuneListOptions struct {
	Locale       string   `url:"locale"`
	Version      string   `url:"version"`
	RuneListData []string `url:"runeListData"`
	Tags         []string `url:"tags"`
}

// StaticDataRuneListOption is a function which modifies the StaticDataRuneListOptions.
type StaticDataRuneListOption func(*StaticDataRuneListOptions)

// Runes retrieves the list of runes.
func (s *StaticDataService) Runes(opts ...StaticDataRuneListOption) (*RuneListDTO, *http.Response, error) {
	options := &StaticDataRuneListOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/runes", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := &RuneListDTO{}
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// StaticDataRuneOptions specifies the optional parameters for the Static Data rune service method.
type StaticDataRuneOptions struct {
	Locale   string   `url:"locale"`
	Version  string   `url:"version"`
	RuneData []string `url:"runeData"`
	Tags     []string `url:"tags"`
}

// StaticDataRuneOption is a function which modifies the StaticDataRuneOptions.
type StaticDataRuneOption func(*StaticDataRuneOptions)

// RuneByID retrieves a rune by ID.
func (s *StaticDataService) RuneByID(runeID int64, opts ...StaticDataRuneOption) (*RuneDTO, *http.Response, error) {
	options := &StaticDataRuneOptions{}
	for _, o := range opts {
		o(options)
	}

	u := "lol/static-data/v3/runes/" + strconv.FormatInt(runeID, 10)
	u, err := addOptions(u, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := &RuneDTO{}
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// SummonerSpellListDTO contains summoner spell list data.
type SummonerSpellListDTO struct {
	Data    map[string]SummonerSpellDTO `json:"data"`
	Version string                      `json:"version"`
	Type    string                      `json:"type"`
}

// SummonerSpellDTO contains summoner spell data.
type SummonerSpellDTO struct {
	Vars                 []SpellVarsDTO `json:"vars"`
	Image                ImageDTO       `json:"image"`
	CostBurn             string         `json:"costBurn"`
	Cooldown             []int64        `json:"cooldown"`
	EffectBurn           []string       `json:"effectBurn"`
	ID                   int            `json:"id"`
	CooldownBurn         string         `json:"cooldownBurn"`
	Tooltip              string         `json:"tooltip"`
	MaxRank              int            `json:"maxrank"`
	RangeBurn            string         `json:"rangeBurn"`
	Description          string         `json:"description"`
	Effect               [][]int64      `json:"effect"`
	Key                  string         `json:"key"`
	LevelTip             LevelTipDTO    `json:"leveltip"`
	Modes                []string       `json:"modes"`
	Resource             string         `json:"resource"`
	Name                 string         `json:"name"`
	CostType             string         `json:"costType"`
	SanitizedDescription string         `json:"sanitizedDescription"`
	SanitizedTooltip     string         `json:"sanitizedTooltip"`
	Range                interface{}    `json:"range"` // This field is either a List of Integer or the String 'self' for spells that target one's own champion.
	Cost                 []int          `json:"cost"`
	SummonerLevel        int            `json:"summonerLevel"`
}

// StaticDataSpellListOptions specifies the optional parameters for the Static Data spell list service method.
type StaticDataSpellListOptions struct {
	Locale        string   `url:"locale"`
	Version       string   `url:"version"`
	SpellListData []string `url:"spellListData"`
	DataByID      bool     `url:"dataById"`
	Tags          []string `url:"tags"`
}

// StaticDataSpellListOption is a function which modifies the StaticDataSpellListOptions.
type StaticDataSpellListOption func(*StaticDataSpellListOptions)

// SummonerSpells retrieves summoner spell list.
func (s *StaticDataService) SummonerSpells(opts ...StaticDataSpellListOption) (*SummonerSpellListDTO, *http.Response, error) {
	options := &StaticDataSpellListOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/summoner-spells", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	ss := &SummonerSpellListDTO{}
	resp, err := s.client.Do(req, ss)
	if err != nil {
		return nil, resp, err
	}

	return ss, resp, nil
}

// StaticDataSpellOptions specifies the optional parameters for the Static Data spell  service method.
type StaticDataSpellOptions struct {
	Locale    string   `url:"locale"`
	Version   string   `url:"version"`
	SpellData []string `url:"spellData"`
	DataByID  bool     `url:"dataById"`
	Tags      []string `url:"tags"`
}

// StaticDataSpellOption is a function which modifies the StaticDataSpellOptions.
type StaticDataSpellOption func(*StaticDataSpellOptions)

// SummonerSpellByID retrieves a summoner spell by ID.
func (s *StaticDataService) SummonerSpellByID(summonerSpellID int64, opts ...StaticDataSpellOption) (*SummonerSpellDTO, *http.Response, error) {
	options := &StaticDataSpellOptions{}
	for _, o := range opts {
		o(options)
	}

	u := "lol/static-data/v3/summoner-spells/" + strconv.FormatInt(summonerSpellID, 10)
	u, err := addOptions(u, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	ss := &SummonerSpellDTO{}
	resp, err := s.client.Do(req, ss)
	if err != nil {
		return nil, resp, err
	}

	return ss, resp, nil
}

// TarballLinksOptions specifies the optional parameters for the Static Data tarball links service method.
type TarballLinksOptions struct {
	Version string `url:"version"`
}

// TarballLinksOption is a function which modifies the TarballLinksOptions.
type TarballLinksOption func(*TarballLinksOptions)

// TarballLinks retrieves full tarball link.
func (s *StaticDataService) TarballLinks(opts ...TarballLinksOption) (*string, *http.Response, error) {
	options := &TarballLinksOptions{}
	for _, o := range opts {
		o(options)
	}

	u, err := addOptions("lol/static-data/v3/tarball-links", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var tl *string
	resp, err := s.client.Do(req, tl)
	if err != nil {
		return nil, resp, err
	}

	return tl, resp, nil
}

// Versions retrieves a list of valid versions.
func (s *StaticDataService) Versions() ([]string, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "lol/static-data/v3/versions", nil)
	if err != nil {
		return nil, nil, err
	}

	var v []string
	resp, err := s.client.Do(req, &v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
