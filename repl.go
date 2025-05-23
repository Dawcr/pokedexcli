package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dawcr/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *pokeapi.Config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		reader.Scan()

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}
		userChoice := input[0]

		cmd, ok := getCommands()[userChoice]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg); err != nil {
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
	callback    func(*pokeapi.Config) error
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
			description: "return 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "return previous 20 locations",
			callback:    commandMapB,
		},
	}
}
