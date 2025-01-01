package main

import (
	"time"

	"github.com/dawcr/pokedexcli/internal/pokeapi"
)

func main() {
	configurations := pokeapi.NewConfig(5 * time.Second)
	startRepl(configurations)
}
