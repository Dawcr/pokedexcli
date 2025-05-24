// Package pokeapi provides functionality for interacting with PokeAPI
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Get a list of locations
func (c *Client) GetLocation(url *string) (LocationArea, error) {
	var data LocationArea
	if err := c.GetData(url, &data); err != nil {
		return LocationArea{}, err
	}

	return data, nil
}

// Get a list of pokemon found at a location
func (c *Client) GetMons(url *string) (LocationMons, error) {
	var data LocationMons
	if err := c.GetData(url, &data); err != nil {
		return LocationMons{}, err
	}

	return data, nil
}

func (c *Client) GetMonDetails(url *string) (Pokemon, error) {
	var data Pokemon
	if err := c.GetData(url, &data); err != nil {
		return Pokemon{}, err
	}

	return data, nil
}

// GetData is a helper function to handle the http get request
// Returns an error if the request fails
func (c *Client) GetData(url *string, data interface{}) error {
	raw, exists := c.cache.Get(*url)
	if !exists {
		req, err := http.NewRequest("GET", *url, nil)
		if err != nil {
			return err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("request unsuccessful, returned with status code: %v", res.StatusCode)
		}

		// Dump response as []byte for caching
		raw, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		c.cache.Add(*url, raw)
	}

	if err := json.Unmarshal(raw, data); err != nil {
		return err
	}

	return nil
}
