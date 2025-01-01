package pokeapi

import "time"

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// Config holds the pagination URLs for the PokeAPI requests
type Config struct {
	pokeapiClient        Client
	nextLocationsURL     *string // URL for the next page of results
	previousLocationsURL *string // URL for the previous page of results
}

// NewConfig creates and initializes a new Config with the base URL
func NewConfig(timeout time.Duration) *Config {
	URL := baseURL + "/location-area"
	client := NewClient(timeout)
	return &Config{
		pokeapiClient:        client,
		nextLocationsURL:     &URL,
		previousLocationsURL: nil,
	}
}
