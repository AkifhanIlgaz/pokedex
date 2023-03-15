package pokeapi

type PokeLocation struct {
	LocationName string `json:"name"`
	URL          string `json:"url"`
}

type PokeJSONStruct struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []PokeLocation `json:"results"`
}
