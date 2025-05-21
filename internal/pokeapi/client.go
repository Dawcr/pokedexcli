package pokeapi

import (
	"net/http"
	"time"

	"github.com/dawcr/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration, cacheinterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheinterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
