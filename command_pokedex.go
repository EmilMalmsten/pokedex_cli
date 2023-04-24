package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("Your pokedex:")

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any pokemon yet.")
		return nil
	}

	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s", pokemon.Name)
	}

	return nil
}
