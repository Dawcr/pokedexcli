package main

import (
	"time"

	"github.com/dawcr/pokedexcli/internal/pokeapi"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// Config holds the pagination URLs for the PokeAPI requests
type Config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string // URL for the next page of results
	previousLocationsURL *string // URL for the previous page of results
}

// NewConfig creates and initializes a new Config with the base URL
func NewConfig(timeout time.Duration, cacheLength time.Duration) *Config {
	URL := baseURL + "/location-area"
	client := pokeapi.NewClient(timeout, cacheLength)
	return &Config{
		pokeapiClient:        client,
		nextLocationsURL:     &URL,
		previousLocationsURL: nil,
	}
}
