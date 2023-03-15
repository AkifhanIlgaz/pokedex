package main

import (
	"time"

	"github.com/AkifhanIlgaz/pokedex/internal/pokeapi"
	"github.com/AkifhanIlgaz/pokedex/internal/pokemons"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokedex := pokemons.NewPokedex()
	cfg := &Config{
		Pokedex: pokedex,
		Client:  pokeClient,
	}
	startPokedex(cfg)
}
