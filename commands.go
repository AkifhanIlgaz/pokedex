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
		c.Println(location.Name)
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
		c.Println(location.Name)
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
	if rand > 30 {
		color.Magenta("%s escaped!", pokemonName)
		return nil

	}

	color.Magenta("%s was caught!", pokemonName)
	color.Yellow("You may now inspect it with the inspect command")
	config.CaughtPokemons[pokemonName] = pokemon
	return nil
}

func commandInspect(config *Config, args ...string) error {
	pokemonName := args[0]

	if pokemon, ok := config.CaughtPokemons[pokemonName]; ok {
		keyColor := color.New(color.FgMagenta, color.Bold)
		valueColor := color.New(color.FgHiWhite, color.Bold)

		keyColor.Print("Name: ")
		valueColor.Println(pokemon.Name)

		keyColor.Print("Height: ")
		valueColor.Println(pokemon.Height)

		keyColor.Print("Weight: ")
		valueColor.Println(pokemon.Weight)

		color.Yellow("Stats:")
		for _, statics := range pokemon.Stats {
			valueColor.Print("  - ")
			keyColor.Printf("%s: ", statics.Stat.Name)
			valueColor.Println(statics.BaseStat)

		}
		color.Yellow("Types:")
		for _, t := range pokemon.Types {
			valueColor.Print("  - ")
			keyColor.Println(t.Type.Name)
		}
	} else {
		fmt.Println("You have not caught that pokemon")
	}

	return nil
}

func commandPokedex(config *Config, args ...string) error {
	if len(config.CaughtPokemons) == 0 {
		color.Red("You didn't catch any pokemons")
		return nil
	}

	color.HiMagenta("Your Pokedex:")

	for pokemonName := range config.CaughtPokemons {
		dashColor := color.New(color.FgHiWhite, color.Bold)
		dashColor.Print("  - ")
		color.Magenta(pokemonName)
	}

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
