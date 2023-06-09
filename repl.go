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
	callback    func(*Config, ...string) error
}

func startPokedex(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)

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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		allCommands := getCommands()

		command, ok := allCommands[commandName]
		if ok {
			err := command.callback(config, args...)
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
		"explore": {
			name:        "explore <area_name>",
			description: "Display the names of Pokemons in the given location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try to catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "See details of the pokemon if you have already caught it",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemons",
			callback:    commandPokedex,
		},
	}
}
