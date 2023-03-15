package pokemons

import "github.com/AkifhanIlgaz/pokedex/internal/pokeapi"

type Pokedex struct {
	pokemons map[string]pokeapi.Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{pokemons: map[string]pokeapi.Pokemon{}}
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	p.pokemons[pokemon.Name] = pokemon
}
