package main

import (
	"fmt"
	"log"
)

func commandMap(cfg *config, args ...string) error {
	locationAreasResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, area := range locationAreasResp.Results {
		fmt.Println(area.Name)
	}
	cfg.nextLocationAreaURL = locationAreasResp.Next
	cfg.prevLocationAreaURL = locationAreasResp.Previous
	return nil

}

func commandMapb(cfg *config, args ...string) error {
	locationAreasResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}
	for _, area := range locationAreasResp.Results {
		fmt.Println(area.Name)
	}
	cfg.nextLocationAreaURL = locationAreasResp.Next
	cfg.prevLocationAreaURL = locationAreasResp.Previous
	return nil

}
