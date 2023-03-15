package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func commandMap(config *Config) error {
	poke := getLocations(config.Next)
	config.Next = poke.Next
	config.Previous = poke.Previous

	for _, location := range poke.Results {
		c := color.Set(color.FgHiWhite, color.Bold)
		c.Println(location.LocationName)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.Previous == nil {
		fmt.Println("No previous page")
		return nil
	}

	poke := getLocations(*config.Previous)
	config.Next = *config.Previous
	config.Previous = poke.Previous

	for _, location := range poke.Results {
		c := color.Set(color.FgHiWhite, color.Bold)
		c.Println(location.LocationName)
	}

	return nil
}

func commandHelp(config *Config) error {
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

func commandExit(config *Config) error {
	os.Exit(1)
	return nil
}
