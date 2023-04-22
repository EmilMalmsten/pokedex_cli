package main

import (
	"time"

	"github.com/emilmalmsten/pokedex_cli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	pokeClient := pokeapi.NewClient(time.Minute * 45)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
