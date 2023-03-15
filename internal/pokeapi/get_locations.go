package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetLocations(pageUrl *string) (PokeJSONStruct, error) {
	url := baseURL + "/location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}
	res, err := http.Get(url)
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

	return locationResp, nil
}
