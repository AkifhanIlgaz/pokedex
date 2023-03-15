package main

import (
	"time"

	"github.com/AkifhanIlgaz/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	cfg := &Config{
		CaughtPokemons: map[string]pokeapi.Pokemon{},
		Client:         pokeClient,
	}
	
	startPokedex(cfg)
}
