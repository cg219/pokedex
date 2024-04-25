package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/charmbracelet/log"
)

const base = "https://pokeapi.co/api/v2"

type LocationQuery struct{
    offset int
    limit int
}

type LocationResp struct {
    Count int `json:"count"`
    Next string `json:"next"`
    Previous string `json:"previous"`
    Results []struct {
        Name string `json:"name"`
        URL string `json:"url"`
    } `json:"results"`
}

type Location struct {
    Name string
    URL string
}

func extractLocationQuery(rawurl string) LocationQuery {
    lq := LocationQuery{}

    u, err := url.Parse(rawurl)

    if err != nil {
        return lq
    }

    q, err := url.ParseQuery(u.RawQuery)

    if err != nil {
        return lq
    }

    offset, err := strconv.Atoi(q.Get("offset"))

    if err != nil {
        return lq
    }

    limit, err := strconv.Atoi(q.Get("limit"))

    if err != nil {
        return lq
    }

    lq.offset = offset
    lq.limit = limit

    return lq
}

func GetLocation(l LocationQuery) ([]Location, LocationQuery, LocationQuery, error) {
    res, err := http.Get(fmt.Sprintf("%s/location-area?offset=%dlimit=%d", base, l.offset, l.limit))

    if err != nil {
        log.Error(err)
    }

    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)

    if err != nil {
        log.Error(err)
    }

    results := LocationResp{}
    err = json.Unmarshal(body, &results)

    if err != nil {
        log.Error(err)
    }


    nl := extractLocationQuery(results.Next)
    pl := extractLocationQuery(results.Previous)

    locations := []Location{}

    for _, l := range results.Results {
        locations = append(locations, Location{ Name: l.Name, URL: l.URL })
    }
    
    return locations, nl, pl, nil
}
