package ionia

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
)

const (
	baseURLPath = "/test"
)

// Creates a test HTTP server along with an ionia.Client.
// Tests should register handlers on mux which provide mock
// responses for the API method being tested.
func createTestServer() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "FAIL: Client.BaseURL path prefix is not preserved in the request URL:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tDid you accidentally use an absolute endpoint URL rather than relative?")
		http.Error(w, "Client.BaseURL path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	server := httptest.NewServer(apiHandler)

	client = NewClient()
	url, _ := url.Parse(server.URL + baseURLPath + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func TestWithAPIKey(t *testing.T) {
	key := "testkey123"
	client := NewClient(WithAPIKey(key))

	if key != client.apiKey {
		t.Errorf("expected client.apiKey to be %s, got %s", key, client.apiKey)
	}
}

func TestWithRegion(t *testing.T) {
	region := "test"
	expected := fmt.Sprintf(defaultBaseURL, region)
	client := NewClient(WithRegion(region))

	if expected != client.BaseURL.String() {
		t.Errorf("expected url: %s, got: %s", expected, client.BaseURL.String())
	}
}

func TestNewRequest(t *testing.T) {
	apiKey := "testing"
	c := NewClient(WithAPIKey(apiKey))

	req, err := c.NewRequest(http.MethodGet, "https://example.com/test/", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}

	if apiKey != req.Header.Get(headerRiotToken) {
		t.Errorf("expected %s header to be %s, got %s", headerRiotToken, apiKey, req.Header.Get(headerRiotToken))
	}
}

func TestNewRequest_ErrorForNoTrailingSlash(t *testing.T) {
	tt := []struct {
		rawurl    string
		wantError bool
	}{
		{rawurl: "https://example.com/test", wantError: true},
		{rawurl: "https://example.com/test/", wantError: false},
	}
	c := NewClient()
	for _, tc := range tt {
		u, err := url.Parse(tc.rawurl)
		if err != nil {
			t.Fatalf("url.Parse returned unexpected error: %v.", err)
		}
		c.BaseURL = u
		if _, err := c.NewRequest(http.MethodGet, "test", nil); tc.wantError && err == nil {
			t.Fatalf("expected error to be returned")
		} else if !tc.wantError && err != nil {
			t.Fatalf("NewRequest returned unexpected error: %v", err)
		}
	}
}

func TestDo(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	type test struct {
		Test string `json:"test"`
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{ "test": "test string" }`)
	})

	req, _ := client.NewRequest(http.MethodGet, ".", nil)
	t.Logf("created request %v", req)
	body := &test{}

	if _, err := client.Do(req, body); err != nil {
		t.Errorf("failed to Do: %v", err)
	}

	want := &test{"test string"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}

func TestDo_HTTPError(t *testing.T) {
	client, mux, _, teardown := createTestServer()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	})

	req, _ := client.NewRequest(http.MethodGet, ".", nil)
	resp, err := client.Do(req, nil)

	if err == nil {
		t.Fatalf("expected HTTP 400 error, got no error")
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected HTTP 400 error, got %d status code", resp.StatusCode)
	}
}

func TestParseLimits(t *testing.T) {
	tt := []struct {
		name     string
		limits   string
		expected map[int]Limit
	}{
		{
			name:   "Riot Limits Example",
			limits: "100:1,1000:10,60000:600,360000:3600",
			expected: map[int]Limit{
				1:    Limit{100, 1},
				10:   Limit{1000, 10},
				600:  Limit{60000, 600},
				3600: Limit{360000, 3600},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if l := parseLimits(tc.limits); !reflect.DeepEqual(l, tc.expected) {
				t.Errorf("unexpected limits parsed: expected %v, got %v", tc.expected, l)
			}
		})
	}
}

func TestParseCounts(t *testing.T) {
	tt := []struct {
		name     string
		counts   string
		expected map[int]Count
	}{
		{
			name:   "Riot Counts Example",
			counts: "1:1,1:10,1:600,1:3600",
			expected: map[int]Count{
				1:    Count{1, 1},
				10:   Count{1, 10},
				600:  Count{1, 600},
				3600: Count{1, 3600},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if c := parseCounts(tc.counts); !reflect.DeepEqual(c, tc.expected) {
				t.Errorf("unexpected counts parsed: got %v, want %v", c, tc.expected)
			}
		})
	}
}

func TestParseRateString(t *testing.T) {
	tt := []struct {
		name     string
		rate     string
		expected map[int]int
	}{
		{
			name: "Riot Example Rate",
			rate: "1:1,1:10,1:600,1:3600",
			expected: map[int]int{
				1:    1,
				10:   1,
				600:  1,
				3600: 1,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if r := parseRateString(tc.rate); !reflect.DeepEqual(r, tc.expected) {
				t.Errorf("unexpected rate map: got %v, want %v", r, tc.expected)
			}
		})
	}
}

func TestGetMethod(t *testing.T) {
	tt := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Get All Champions",
			path:     "/lol/platform/v3/champions",
			expected: "GET_getAllChampions",
		},
		{
			name:     "Get Champion By ID",
			path:     "/lol/platform/v3/champions/123",
			expected: "GET_getChampionById",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if m := getMethod(tc.path); m != tc.expected {
				t.Errorf("unexpected method returned: got %s, want %s", m, tc.expected)
			}
		})
	}
}

func TestCheckRateLimit(t *testing.T) {
	c := NewClient()

	c.rateLimits = map[string]Rate{
		"okay": Rate{
			Counts: map[int]Count{
				1: Count{9, 1},
			},
			Limits: map[int]Limit{
				1: Limit{10, 1},
			},
		},
		"error": Rate{
			Counts: map[int]Count{
				1: Count{10, 1},
			},
			Limits: map[int]Limit{
				1: Limit{10, 1},
			},
		},
	}

	if resp := c.checkRateLimit(nil, "okay"); resp != nil {
		t.Errorf("unexpected response returned: %v", resp)
	}
	if resp := c.checkRateLimit(nil, "error"); resp == nil {
		t.Errorf("expected response, got: %v", resp)
	}
}
