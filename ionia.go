package ionia

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

const (
	// Default base URL for the Riot API.
	// Must be formatted with a valid region string in order to be used in requests.
	defaultBaseURL = "https://%s.api.riotgames.com/"
	defaultRegion  = "na"

	headerRiotToken            = "X-Riot-Token"
	headerRateLimitType        = "X-Rate-Limit-Type"
	headerRetryAfter           = "Retry-After"
	headerAppRateLimit         = "X-App-Rate-Limit"
	headerMethodRateLimit      = "X-Method-Rate-Limit"
	headerAppRateLimitCount    = "X-App-Rate-Limit-Count"
	headerMethodRateLimitCount = "X-Method-Rate-Limit-Count"
)

// Client manages communication with the Riot API.
type Client struct {
	client *http.Client

	BaseURL *url.URL

	// Reuse a single struct instead of allocating on for each service on the heap.
	common service

	rateMu     sync.Mutex
	rateLimits map[string]Rate

	// Riot API Key.
	apiKey string

	// API Sections.
	ChampionMastery *ChampionMasteryService
	Champion        *ChampionService
	League          *LeagueService
}

type service struct {
	client *Client
}

// ClientOption is a function which modifies the ionia client.
type ClientOption func(*Client)

// WithAPIKey returns a ClientOption which sets the Client's API key.
func WithAPIKey(key string) ClientOption {
	return func(c *Client) {
		c.apiKey = key
	}
}

// WithRegion returns a ClientOption which sets the Client's region to the given value.
func WithRegion(region string) ClientOption {
	return func(c *Client) {
		c.BaseURL, _ = url.Parse(fmt.Sprintf(defaultBaseURL, region))
	}
}

// NewClient creates a new Riot API client.
// Any number of ClientOptions can be passed, and will
// be applied after the default client has been created.
func NewClient(opts ...ClientOption) *Client {
	baseURL, _ := url.Parse(fmt.Sprintf(defaultBaseURL, defaultRegion))

	c := &Client{
		client:     http.DefaultClient,
		BaseURL:    baseURL,
		rateLimits: make(map[string]Rate),
	}
	c.common.client = c
	c.ChampionMastery = (*ChampionMasteryService)(&c.common)
	c.Champion = (*ChampionService)(&c.common)
	c.League = (*LeagueService)(&c.common)

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// NewRequest creates a new API request. A relative URL can be provided in urlStr,
// in which case it is resolved to the BaseURL of the Client. Relative URLs should
// always be specified with out a preceding slash.
func (c *Client) NewRequest(urlStr string) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	url, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(headerRiotToken, c.apiKey)

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error
// if an API error has occured.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	rateMethod := getMethod(req.URL.Path)
	if errResp := c.checkRateLimit(req, rateMethod); errResp != nil {
		return errResp, nil
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse rate limit information.
	appRate, methodRate := parseRates(resp)
	c.rateMu.Lock()
	c.rateLimits["app"] = *appRate
	c.rateLimits[rateMethod] = *methodRate
	c.rateMu.Unlock()

	// All valid Riot API responses should return 200 OK.
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("api returned error: %s %d", resp.Status, resp.StatusCode)
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
	}

	return resp, err
}

func (c *Client) checkRateLimit(req *http.Request, method string) *http.Response {
	c.rateMu.Lock()
	rate := c.rateLimits[method]
	c.rateMu.Unlock()

	for s, count := range rate.Counts {
		if r, ok := rate.Limits[s]; ok {
			if count.Used >= r.Allowed {
				// Create a fake response.
				return &http.Response{
					Status:     http.StatusText(http.StatusForbidden),
					StatusCode: http.StatusForbidden,
					Request:    req,
					Header:     make(http.Header),
					Body:       ioutil.NopCloser(strings.NewReader("")),
				}
			}
		}
	}

	return nil
}

// Looks up the method name of the given path.
// These names are based on the ones given on the Rate Limits page
// of the Riot API Documentation (https://developer.riotgames.com/rate-limiting.html).
//
// The format is the HTTP Request Method followed by a few words to describe the method.
// (e.g. /lol/platform/v3/champions => GET_getAllChampions).
func getMethod(path string) string {
	switch {
	case strings.HasPrefix(path, "/lol/platform/v3/champions"):
		if strings.HasSuffix(path, "/lol/platform/v3/champions") {
			return "GET_getAllChampions"
		}
		return "GET_getChampionById"
	}

	return path
}

// Parses all of the rate limit information returned from an API request.
//
// Each request will contain information about the application limits,
// method limits, and service limits (if applicable).
func parseRates(r *http.Response) (appRate *Rate, methodRate *Rate) {
	// TODO: Figure out how to use rate type and retry after headers
	if rateType := r.Header.Get(headerRateLimitType); rateType != "" {

	}
	if retryAfter := r.Header.Get(headerRetryAfter); retryAfter != "" {

	}

	return &Rate{
			Counts: parseCounts(r.Header.Get(headerAppRateLimitCount)),
			Limits: parseLimits(r.Header.Get(headerAppRateLimit)),
		},
		&Rate{
			Counts: parseCounts(r.Header.Get(headerMethodRateLimitCount)),
			Limits: parseLimits(r.Header.Get(headerMethodRateLimit)),
		}
}

// 100:1,1000:10,60000:600,360000:3600
func parseLimits(limits string) map[int]Limit {
	lims := make(map[int]Limit)

	rates := parseRateString(limits)
	for s, a := range rates {
		lims[s] = Limit{a, s}
	}

	return lims
}

func parseCounts(counts string) map[int]Count {
	cts := make(map[int]Count)

	rates := parseRateString(counts)
	for s, u := range rates {
		cts[s] = Count{u, s}
	}

	return cts
}

func parseRateString(rates string) map[int]int {
	rts := make(map[int]int)

	rs := strings.Split(rates, ",")
	for _, r := range rs {
		rsplit := strings.Split(r, ":")
		if len(rsplit) < 2 {
			continue
		}
		l, _ := strconv.Atoi(rsplit[0])
		r, _ := strconv.Atoi(rsplit[1])
		rts[r] = l
	}

	return rts
}

// Rate represents the current rate limit state of a rate limit.
// This can be either an Application, Method, or Service rate limit.
type Rate struct {
	Limits map[int]Limit
	Counts map[int]Count
}

// Limit represents a limit returned by a request.
// These limits are the Application, Method, and Service rate limits.
type Limit struct {
	Allowed int
	Seconds int
}

// Count represents a rate limit count return by a request.
type Count struct {
	Used    int
	Seconds int
}