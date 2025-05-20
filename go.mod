module github.com/dawcr/pokedexcli

go 1.23.3

require github.com/dawcr/pokedexcli/internal/pokeapi v1.0.0

replace github.com/dawcr/pokedexcli/internal/pokeapi => ./internal/pokeapi

replace github.com/dawcr/pokedexcli/internal/pokecache => ./internal/pokecache
