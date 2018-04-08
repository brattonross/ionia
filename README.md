# ionia #

[![GoDoc](https://godoc.org/github.com/brattonross/ionia?status.svg)](https://godoc.org/github.com/brattonross/ionia) [Build Status](https://travis-ci.org/brattonross/ionia.svg?branch=master)](https://travis-ci.org/brattonross/ionia) [![Coverage Status](https://coveralls.io/repos/github/brattonross/ionia/badge.svg?branch=master)](https://coveralls.io/github/brattonross/ionia?branch=master)

ionia is a Go client library for accessing the [Riot API](https://developer.riotgames.com/)

This project is inspired by Google's [go-github](https://github.com/google/go-github).
Credit for many of the techniques used in this library goes to the go-github team.

### TODO ###

ionia is still a work in progress.
Things still to do:

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
client := ionia.NewClient(ionia.WithAPIKey("my-riot-api-key"))

// Retrieve information for a summoner name.
summoner, _, err := client.Summoner.BySummonerName("Doublelift")
```

ionia takes a functional approach to parameters in many places. This means that instead of passing a data structure to the method, you instead pass a function which takes that data structure and modifies it. An example of this can be seen above when we pass "ionia.WithAPIKey"; this is a function which takes a string, and then modifies the client to set the associated API Key to that string.

Here's another example:

```go
// Modify the client's region using ionia's built in WithRegion function.
client := ionia.NewClient(
    ionia.WithAPIKey("my-riot-api-key"),
    ionia.WithRegion("EUW1")
)
```

The above code will create a new client with the API key "my-riot-api-key", and the region "EUW1".
Taking this functional approach means that code that you write will often be cleaner, and you can even write your own functional options to pass in.

Here is an example of an API method that takes optional parameters:

```go
client := ionia.NewClient(ionia.WithAPIKey("my-riot-api-key"))

// Retrieve a list of champions, and order the data by ID.
champions, _, err := client.StaticData.Champions(func(s *StaticDataChampionsOptions) {
    s.DataByID = true
})
```

### Rate Limiting ###

Riot imposes a rate limit on a per key, per method, and per service basis.
The ionia client aims to respect those limits, and does not perform requests to the Riot API if it detects that a rate limit will be / has been breached. Rate limits imposed by Riot will vary per user.

Learn more about rate limiting at https://developer.riotgames.com/rate-limiting.html.