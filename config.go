package main

import (
	"github.com/AkifhanIlgaz/pokedex/internal/pokeapi"
	"github.com/AkifhanIlgaz/pokedex/internal/pokemons"
)

type Config struct {
	Pokedex  pokemons.Pokedex
	Client   pokeapi.Client
	Next     *string
	Previous *string
}
