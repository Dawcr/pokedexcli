package main

import (
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	t.Run("new config should have correct initial state", func(t *testing.T) {
		config := NewConfig(5*time.Second, 5*time.Minute)

		baseURL := "https://pokeapi.co/api/v2/location-area"
		if *config.nextLocationsURL != baseURL {
			t.Errorf("expected Next URL to be %q, got %q", baseURL, *config.nextLocationsURL)
		}

		if config.previousLocationsURL != nil {
			t.Error("expected Previous URL to be nil")
		}
	})
}

func TestGetNextLocation(t *testing.T) {

	t.Run("GetNextLocation should error on last page", func(t *testing.T) {
		config := NewConfig(5*time.Second, 5*time.Minute)
		lastPageURL := "example.com"
		config.nextLocationsURL = nil
		config.previousLocationsURL = &lastPageURL

		expectedMsg := "you're on the last page"
		_, err := config.GetNextLocation()
		if err == nil || err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %v", expectedMsg, err)
		}
	})
}

func TestGetPreviousLocation(t *testing.T) {

	t.Run("GetPreviousLocation should error on first page", func(t *testing.T) {
		config := NewConfig(5*time.Second, 5*time.Minute)

		expectedMsg := "you're on the first page"
		_, err := config.GetPreviousLocation()
		if err == nil || err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %v", expectedMsg, err)
		}
	})
}
