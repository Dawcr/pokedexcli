package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, userParams ...string) error {
	if len(userParams) != 1 {
		return errors.New("you must provide a location name or ID")
	}

	url := baseURL + "/location-area/" + userParams[0]

	data, err := cfg.pokeapiClient.GetMons(&url)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", userParams)
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
