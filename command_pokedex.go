package main

import "fmt"

func commandPokedex(cfg *Config, userParams ...string) error {
	fmt.Println("Your Pokedex:")
	for mon := range cfg.pokedex.caughtMons {
		fmt.Println("  - ", mon)
	}
	return nil
}
