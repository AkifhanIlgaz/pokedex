package main

import (
	"fmt"

	"github.com/fatih/color"
)

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
