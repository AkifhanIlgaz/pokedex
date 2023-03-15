package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func startPokedex() {
	scanner := bufio.NewScanner(os.Stdin)
	config := Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: nil,
	}

	for {
		color.Set(color.FgCyan)
		fmt.Print("Pokedex > ")
		color.Unset()
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		allCommands := getCommands()

		command, ok := allCommands[commandName]
		if ok {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			color.Red("Error: Unknown command!")
			continue
		}

	}
}

func cleanInput(input string) []string {
	input = strings.ToLower(input)
	return strings.Fields(input)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of previous 20 locations",
			callback:    commandMapb,
		},
	}
}
