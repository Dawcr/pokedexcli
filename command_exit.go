package main

import (
	"fmt"
	"os"

	"github.com/dawcr/pokedexcli/internal/pokeapi"
)

func commandExit(conf *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
