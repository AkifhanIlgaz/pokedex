package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(pageUrl *string) (PokeJSONStruct, error) {
	url := baseURL + "/location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("cache is using")
		locationResp := PokeJSONStruct{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return PokeJSONStruct{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeJSONStruct{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeJSONStruct{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return PokeJSONStruct{}, err
	}

	locationResp := PokeJSONStruct{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return PokeJSONStruct{}, err
	}

	c.cache.Add(url, body)

	return locationResp, nil
}
