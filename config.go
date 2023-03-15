package main

import "github.com/AkifhanIlgaz/pokedex/internal/pokeapi"

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}
