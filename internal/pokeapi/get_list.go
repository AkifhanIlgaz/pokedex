package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokeJSONStruct, error) {
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := PokeJSONStruct{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokeJSONStruct{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeJSONStruct{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeJSONStruct{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeJSONStruct{}, err
	}

	locationsResp := PokeJSONStruct{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return PokeJSONStruct{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
