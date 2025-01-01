package main

import (
	"fmt"

	"github.com/dawcr/pokedexcli/internal/pokeapi"
)

func commandHelp(conf *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
