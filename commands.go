package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/fatih/color"
)

func commandMap(config *Config, args ...string) error {
	locationResp, err := config.Client.ListLocations(config.Next)
	if err != nil {
		return err
	}

	config.Next = locationResp.Next
	config.Previous = locationResp.Previous

	for _, location := range locationResp.Results {
		c := color.Set(color.FgHiBlue, color.Bold)
		c.Println(location.LocationName)
	}

	return nil
}

func commandMapb(config *Config, args ...string) error {
	if config.Previous == nil {
		color.Red("Error: No previous page")
		return nil
	}

	locationResp, err := config.Client.ListLocations(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locationResp.Next
	config.Previous = locationResp.Previous

	for _, location := range locationResp.Results {
		c := color.Set(color.FgHiBlue, color.Bold)
		c.Println(location.LocationName)
	}

	return nil
}

func commandExplore(config *Config, args ...string) error {
	locationName := args[0]
	color.Magenta("Exploring %s...", locationName)

	locationResp, err := config.Client.GetLocations(locationName)
	if err != nil {
		return err
	}
	for _, pokemon := range locationResp.PokemonEncounters {
		c := color.Set(color.FgHiBlue, color.Bold)
		c.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *Config, args ...string) error {
	pokemonName := args[0]
	color.Magenta("Throwing a Pokeball at %s...", pokemonName)

	pokemon, err := config.Client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	rand := rand.Intn(pokemon.BaseExperience)
	if rand > 10 {
		color.Magenta("%s escaped!", pokemonName)
		return nil

	}
	
	color.Magenta("%s was caught!", pokemonName)
	config.Pokedex.Add(pokemon)
	return nil
}

func commandHelp(config *Config, args ...string) error {
	fmt.Println()
	color.HiMagenta("Welcome to Pokedex!")
	fmt.Println()
	color.HiBlue("Usage:")

	for _, command := range getCommands() {
		c := color.New(color.FgHiYellow, color.Bold)
		c.Printf("%-5s: ", command.name)
		fmt.Printf("%s\n", command.description)
	}

	fmt.Println()
	return nil
}

func commandExit(config *Config, args ...string) error {
	os.Exit(1)
	return nil
}
