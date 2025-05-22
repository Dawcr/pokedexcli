package main

import "fmt"

func commandExplore(cfg *Config, userParam string) error {
	url := baseURL + "/location-area/" + userParam
	data, err := cfg.pokeapiClient.GetMons(&url)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", userParam)
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
