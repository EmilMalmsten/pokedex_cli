package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(locationName string) (Location, error) {
	fullURL := baseURL + "/location-area/" + locationName

	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		locationResp := Location{}
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return locationResp, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Location{}, fmt.Errorf("Response failed with status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return locationResp, err
	}

	c.cache.Add(fullURL, data)

	return locationResp, nil
}
