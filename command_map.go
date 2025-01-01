package main

import (
	"fmt"

	"github.com/dawcr/pokedexcli/internal/pokeapi"
)

func commandMapf(conf *pokeapi.Config) error {
	data, err := pokeapi.GetNextLocation(conf)
	if err != nil {
		return err
	}
	for _, locationData := range data.Results {
		fmt.Println(locationData.Name)
	}
	return nil
}

func commandMapB(conf *pokeapi.Config) error {
	data, err := pokeapi.GetPreviousLocation(conf)
	if err != nil {
		return err
	}
	for _, locationData := range data.Results {
		fmt.Println(locationData.Name)
	}
	return nil
}
