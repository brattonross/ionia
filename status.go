package ionia

import "net/http"

// StatusService represents the LOL-Status-V3 API methods.
// https://developer.riotgames.com/api-methods/#lol-status-v3
type StatusService service

// ShardStatus contains shard status data.
type ShardStatus struct {
	Name      string    `json:"name"`
	RegionTag string    `json:"region_tag"`
	HostName  string    `json:"hostname"`
	Services  []Service `json:"services"`
	Slug      string    `json:"slug"`
	Locales   []string  `json:"locales"`
}

// Service contains service data.
type Service struct {
	Status    string     `json:"status"`
	Incidents []Incident `json:"incidents"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
}

// Incident contains incident data.
type Incident struct {
	Active    bool      `json:"active"`
	CreatedAt string    `json:"created_at"`
	ID        int64     `json:"id"`
	Updates   []Message `json:"updates"`
}

// Message contains message data.
type Message struct {
	Severity     string        `json:"severity"`
	Author       string        `json:"author"`
	CreatedAt    string        `json:"created_at"`
	Translations []Translation `json:"translations"`
	UpdatedAt    string        `json:"updated_at"`
	Content      string        `json:"content"`
	ID           string        `json:"id"`
}

// Translation contains translation data.
type Translation struct {
	Locale    string `json:"locale"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}

// ShardData retrieves the League of Legends status for the given shard.
func (s *StatusService) ShardData() (*ShardStatus, *http.Response, error) {
	req, err := s.client.NewRequest("lol/status/v3/shard-data")
	if err != nil {
		return nil, nil, err
	}

	ss := &ShardStatus{}
	resp, err := s.client.Do(req, ss)
	if err != nil {
		return nil, resp, err
	}

	return ss, resp, nil
}
