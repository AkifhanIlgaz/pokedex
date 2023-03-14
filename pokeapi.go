package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PokeLocation struct {
	LocationName string `json:"name"`
	URL          string `json:"url"`
}

type PokeJSONStruct struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous *string        `json:"previous"`
	Results  []PokeLocation `json:"results"`
}

func getLocations(url string) PokeJSONStruct {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	poke := PokeJSONStruct{}
	err = json.Unmarshal(body, &poke)
	if err != nil {
		fmt.Println(err)
	}

	return poke
}
