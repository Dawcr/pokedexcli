package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *Config, userParams ...string) error {
	if len(userParams) != 1 {
		return errors.New("you must provide a location name or ID")
	}
	target := userParams[0]

	url := baseURL + "/pokemon/" + target
	data, err := cfg.pokeapiClient.GetMonDetails(&url)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", target)

	// Using a worse version of the original catch rate instead of one based on base experience yield as indicated in the lesson.
	// Regular pokeballs roll a number between 0 and 255 inclusive
	R1 := rand.Intn(256)

	// Estimate capture rate from exp yield to not hammer the API
	captureRate := math.Round(252/(1+math.Exp(0.1*float64(data.BaseExperience-150))) + 3)

	// if base capture rate is less than R1 the mon breaks free (not using R* since the mon has no status effects here)
	// if R2(anon var in this case) is less than or equal to the HP factor F(123), the Pokémon is caught.
	// F is calculated via steps listed in https://www.dragonflycave.com/mechanics/gen-i-capturing assuming 69% hp remaining in this case (hopefully my math is correct)
	if captureRate > float64(R1) && math.Min(float64(rand.Intn(256)), 255) < 123 {
		fmt.Printf("%s was caught!\n", target)
		cfg.pokedex.caughtMons[target] = data
	} else {
		fmt.Printf("%s escaped!\n", target)
	}

	return nil
}
