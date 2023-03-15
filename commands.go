package main

import (
	"fmt"
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
	color.Magenta("Exploring %s...", args[0])

	locationResp, err := config.Client.GetLocations(args[0])
	if err != nil {
		return err
	}
	for _, pokemon := range locationResp.PokemonEncounters {
		c := color.Set(color.FgHiBlue, color.Bold)
		c.Println(pokemon.Pokemon.Name)
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
