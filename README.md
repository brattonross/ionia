# ionia #

[![GoDoc](https://godoc.org/github.com/brattonross/ionia?status.svg)](https://godoc.org/github.com/brattonross/ionia) [![Build Status](https://travis-ci.org/brattonross/ionia.svg?branch=master)](https://travis-ci.org/brattonross/ionia) [![Coverage Status](https://coveralls.io/repos/github/brattonross/ionia/badge.svg?branch=master)](https://coveralls.io/github/brattonross/ionia?branch=master)

ionia is a Go client library for accessing the [Riot API](https://developer.riotgames.com/)

### TODO ###

ionia is still a work in progress.
Things still to do:

* Regional proxies
* Write tests for remaining services
* Add better descriptions to many data types and functions
* Possibly add Riot-defined constants (e.g. region IDs)

### Usage ###
```go
import "github.com/brattonross/ionia"
```

Create a new ionia client, then use the client to access the various services of the Riot API.
Example:

```go
// Note: The default region for the client is NA.
// See below for examples of how to customize the client.
client := ionia.NewClient("my-riot-api-key")

// Retrieve information for a summoner name.
summoner, _, err := client.Summoner.BySummonerName("Doublelift")
```

Another example:

```go
// Modify the client's region using ionia's built in WithRegion function.
client := ionia.NewClient("my-riot-api-key", ionia.WithRegion("EUW1"))
```

The above code will create a new client with the API key "my-riot-api-key", and the region "EUW1".

Example of an API method that takes optional parameters:

```go
client := ionia.NewClient("my-riot-api-key")

// Retrieve a list of champions, and order the data by ID.
champions, _, err := client.StaticData.Champions(func(s *StaticDataChampionsOptions) {
    s.DataByID = true
})
```

### Rate Limiting ###

Riot imposes a rate limit on a per key, per method, and per service basis.
The ionia client aims to respect those limits, and does not perform requests to the Riot API if it detects that a rate limit will be / has been breached. Rate limits imposed by Riot will vary per user.

Learn more about rate limiting at https://developer.riotgames.com/rate-limiting.html.
