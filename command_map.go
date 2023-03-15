package main

import (
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
