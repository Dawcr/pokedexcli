package main

import (
	"errors"
	"fmt"

	"github.com/dawcr/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *Config, userParams ...string) error {
	if len(userParams) != 1 {
		return errors.New("you must provide a pokemon to inspect")
	}
	mon := userParams[0]

	entry, ok := cfg.pokedex.caughtMons[mon]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	printPokemonStats(entry)

	return nil
}

func printPokemonStats(entry pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", entry.Name)
	fmt.Printf("Height: %d\n", entry.Height)
	fmt.Printf("Weight: %d\n", entry.Weight)
	fmt.Print("Stats:\n")
	for _, stat := range entry.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, monType := range entry.Types {
		fmt.Printf("  -%s\n", monType.Type.Name)
	}
}
