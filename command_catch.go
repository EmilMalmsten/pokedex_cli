package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func catchPokemon(experienceLevel int) bool {
	// Calculate the catch probability using a power function
	k := 0.001
	p := 1 / (1 + k*float64(experienceLevel))

	// Generate a random number between 0 and 1
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()

	// Compare the random number to the catch probability
	if r < p {
		// The catch is successful
		return true
	} else {
		// The catch failed
		return false
	}
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemon := args[0]

	pokemonResp, err := cfg.pokeapiClient.PokemonInfo(pokemon)
	if err != nil {
		log.Fatal(err)
	}

	catched := catchPokemon(pokemonResp.BaseExperience)
	if catched {
		fmt.Printf("%s was catched successfully!\n", pokemonResp.Name)
		cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("%s escaped!", pokemonResp.Name)
	}

	return nil
}
