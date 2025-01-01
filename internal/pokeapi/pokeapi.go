// Package pokeapi provides functionality for interacting with PokeAPI
package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Config holds the pagination URLs for the PokeAPI requests
type Config struct {
	Next     *string // URL for the next page of results
	Previous *string // URL for the previous page of results
}

// GetNextLocation retrieves the next page of location areas from the PokeAPI
// Returns an error if on the last page or if the request fails
func GetNextLocation(conf *Config) (LocationArea, error) {
	if conf == nil {
		return LocationArea{}, fmt.Errorf("expected config struct, instead received nil value")
	}
	if conf.Next == nil && conf.Previous != nil {
		return LocationArea{}, fmt.Errorf("you're on the last page")
	}

	return getLocation(conf, conf.Next)
}

// GetPreviousLocation retrieves the previous page of location areas from the PokeAPI
// Returns an error if on the first page or if the request fails
func GetPreviousLocation(conf *Config) (LocationArea, error) {
	if conf == nil {
		return LocationArea{}, fmt.Errorf("expected config struct, instead received nil value")
	}
	if conf.Previous == nil {
		return LocationArea{}, fmt.Errorf("you're on the first page")
	}

	return getLocation(conf, conf.Previous)
}

// getLocation is a helper function to handle the http get request
// It updates the Config with new pagination URLs and returns the location data
// Returns an error if the request fails
func getLocation(conf *Config, url *string) (LocationArea, error) {
	res, err := http.Get(*url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("request unsuccessful, returned with status code: %v", res.StatusCode)
	}

	var data LocationArea
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return LocationArea{}, err
	}

	conf.Previous = data.Previous
	conf.Next = data.Next

	return data, nil
}

// NewConfig creates and initializes a new Config with the base URL
func NewConfig() *Config {
	baseURL := "https://pokeapi.co/api/v2/location-area"
	return &Config{
		Next:     &baseURL,
		Previous: nil,
	}
}
