package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(locationName string) (Location, error) {
	url := BaseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {

		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, body)

	return locationResp, nil
}
