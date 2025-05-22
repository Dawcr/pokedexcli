package main

import (
	"fmt"

	"github.com/dawcr/pokedexcli/internal/pokeapi"
)

func commandMapf(cfg *Config, userParam string) error {
	return cfg.handleData(cfg.GetNextLocation())
}

func commandMapB(cfg *Config, userParam string) error {
	return cfg.handleData(cfg.GetPreviousLocation())
}

func (cfg *Config) handleData(data pokeapi.LocationArea, err error) error {
	if err != nil {
		return err
	}
	cfg.previousLocationsURL = data.Previous
	cfg.nextLocationsURL = data.Next
	for _, locationData := range data.Results {
		fmt.Println(locationData.Name)
	}
	return nil
}

// GetNextLocation retrieves the next page of location areas from the PokeAPI
// Returns an error if on the last page or if the request fails
func (cfg *Config) GetNextLocation() (pokeapi.LocationArea, error) {
	if cfg.nextLocationsURL == nil && cfg.previousLocationsURL != nil {
		return pokeapi.LocationArea{}, fmt.Errorf("you're on the last page")
	}

	return cfg.pokeapiClient.GetLocation(cfg.nextLocationsURL)
}

// GetPreviousLocation retrieves the previous page of location areas from the PokeAPI
// Returns an error if on the first page or if the request fails
func (cfg *Config) GetPreviousLocation() (pokeapi.LocationArea, error) {
	if cfg.previousLocationsURL == nil {
		return pokeapi.LocationArea{}, fmt.Errorf("you're on the first page")
	}

	return cfg.pokeapiClient.GetLocation(cfg.previousLocationsURL)
}
