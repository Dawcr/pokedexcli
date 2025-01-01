package pokeapi

// LocationArea represents the response structure from the PokeAPI location-area endpoint
type LocationArea struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

// Location represents Results elements from the response received from the PokeAPI location-area endpoint
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
