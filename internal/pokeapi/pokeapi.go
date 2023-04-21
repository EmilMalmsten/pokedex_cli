package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
    baseURL = "https://pokeapi.co/api/v2"
)

type LocationAreasResp struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
    locationAreasResp := LocationAreasResp{}

    endpoint := "/location-area"
    fullURL := baseURL + endpoint
    if pageURL != nil {
        fullURL = *pageURL
    }

    res, err := http.Get(fullURL)
    if err != nil {
        return locationAreasResp, err
    }
    body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
        return locationAreasResp, fmt.Errorf("Response failed with status code: %v", res.StatusCode)
	}
	if err != nil {
        return locationAreasResp, err
	}

    err = json.Unmarshal(body, &locationAreasResp)
    if err != nil {
        return locationAreasResp, err
    }

    return locationAreasResp, nil
}
