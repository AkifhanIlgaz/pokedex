package main

import "fmt"

func commandMap(config *Config) error {
	poke := getLocations(config.Next)
	config.Next = poke.Next
	config.Previous = poke.Previous

	for _, location := range poke.Results {
		fmt.Println(location.LocationName)
	}
	return nil
}
