package pokeapi

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("new config should have correct initial state", func(t *testing.T) {
		config := NewConfig()

		baseURL := "https://pokeapi.co/api/v2/location-area"
		if *config.Next != baseURL {
			t.Errorf("expected Next URL to be %q, got %q", baseURL, *config.Next)
		}

		if config.Previous != nil {
			t.Error("expected Previous URL to be nil")
		}
	})
}

func TestGetNextLocation(t *testing.T) {
	t.Run("GetNextLocation should complain if conf is nil", func(t *testing.T) {
		expectedMsg := "expected config struct, instead received nil value"
		_, err := GetNextLocation(nil)
		if err == nil || err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %v", expectedMsg, err)
		}
	})

	t.Run("GetNextLocation should error on last page", func(t *testing.T) {
		config := NewConfig()
		lastPageURL := "example.com"
		config.Next = nil
		config.Previous = &lastPageURL

		expectedMsg := "you're on the last page"
		_, err := GetNextLocation(config)
		if err == nil || err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %v", expectedMsg, err)
		}
	})
}

func TestGetPreviousLocation(t *testing.T) {
	t.Run("GetPreviousLocation should complain if conf is nil", func(t *testing.T) {
		expectedMsg := "expected config struct, instead received nil value"
		_, err := GetPreviousLocation(nil)
		if err == nil || err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %v", expectedMsg, err)
		}
	})

	t.Run("GetPreviousLocation should error on first page", func(t *testing.T) {
		config := NewConfig()

		expectedMsg := "you're on the first page"
		_, err := GetPreviousLocation(config)
		if err == nil || err.Error() != expectedMsg {
			t.Errorf("expected error message %q, got %v", expectedMsg, err)
		}
	})
}
