package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func commandMap(config *Config) error {
	locationResp, err := config.Client.GetLocations(config.Next)
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

func commandMapb(config *Config) error {
	if config.Previous == nil {
		color.Red("Error: No previous page")
		return nil
	}

	locationResp, err := config.Client.GetLocations(config.Previous)
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
