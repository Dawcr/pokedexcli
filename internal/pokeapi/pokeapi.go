// Package pokeapi provides functionality for interacting with PokeAPI
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetNextLocation retrieves the next page of location areas from the PokeAPI
// Returns an error if on the last page or if the request fails
func GetNextLocation(cfg *Config) (LocationArea, error) {
	if cfg == nil {
		return LocationArea{}, fmt.Errorf("expected config struct, instead received nil value")
	}
	if cfg.nextLocationsURL == nil && cfg.previousLocationsURL != nil {
		return LocationArea{}, fmt.Errorf("you're on the last page")
	}

	return getLocation(cfg, cfg.nextLocationsURL)
}

// GetPreviousLocation retrieves the previous page of location areas from the PokeAPI
// Returns an error if on the first page or if the request fails
func GetPreviousLocation(cfg *Config) (LocationArea, error) {
	if cfg == nil {
		return LocationArea{}, fmt.Errorf("expected config struct, instead received nil value")
	}
	if cfg.previousLocationsURL == nil {
		return LocationArea{}, fmt.Errorf("you're on the first page")
	}

	return getLocation(cfg, cfg.previousLocationsURL)
}

// getLocation is a helper function to handle the http get request
// It updates the Config with new pagination URLs and returns the location data
// Returns an error if the request fails
func getLocation(cfg *Config, url *string) (LocationArea, error) {

	raw, exists := cfg.cache.Get(*url)
	if !exists {
		req, err := http.NewRequest("GET", *url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		res, err := cfg.pokeapiClient.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationArea{}, fmt.Errorf("request unsuccessful, returned with status code: %v", res.StatusCode)
		}

		// Dump response as []byte for caching
		raw, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}
		cfg.cache.Add(*url, raw)
	}

	var data LocationArea
	if err := json.Unmarshal(raw, &data); err != nil {
		return LocationArea{}, err
	}

	cfg.previousLocationsURL = data.Previous
	cfg.nextLocationsURL = data.Next

	return data, nil
}
