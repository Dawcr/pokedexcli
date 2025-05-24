package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		reader.Scan()

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}
		userChoice := input[0]

		userParams := []string{}
		if len(input) > 1 {
			userParams = input[1:]
		}

		cmd, ok := getCommands()[userChoice]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg, userParams...); err != nil {
			fmt.Printf("%s command returned with error: %s\n", cmd.name, err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Return 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Return previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location name>",
			description: "List all pokemon found in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
	}
}
