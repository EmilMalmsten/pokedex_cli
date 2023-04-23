package main

import (
	"errors"
	"fmt"
	"log"
)

func commandExplore(cfg *config, args ...string) error {
	fmt.Println(args)

	if len(args) != 1 {
		return errors.New("you must provide an area")
	}

	area := args[0]

	locationResp, err := cfg.pokeapiClient.ExploreArea(area)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pokemon found in area:")
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil

}
