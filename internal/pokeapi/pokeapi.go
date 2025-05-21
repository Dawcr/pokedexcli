// Package pokeapi provides functionality for interacting with PokeAPI
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// getLocation is a helper function to handle the http get request
// It updates the Config with new pagination URLs and returns the location data
// Returns an error if the request fails
func (c *Client) GetLocation(url *string) (LocationArea, error) {

	raw, exists := c.cache.Get(*url)
	if !exists {
		req, err := http.NewRequest("GET", *url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		res, err := c.httpClient.Do(req)
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
		c.cache.Add(*url, raw)
	}

	var data LocationArea
	if err := json.Unmarshal(raw, &data); err != nil {
		return LocationArea{}, err
	}

	return data, nil
}
