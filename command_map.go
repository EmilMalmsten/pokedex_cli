package main

import (
    "fmt"
    "github.com/emilmalmsten/pokedex_cli/internal/pokeapi"
    "log"
)

func commandMap() error {

    locationAreasResp, err := pokeapi.ListLocationAreas()
    if err != nil {
        log.Fatal(err)
    }
    for _, area := range locationAreasResp.Results {
        fmt.Println(area.Name)
    }
    return nil

}

func commandMapb() error {

    fmt.Println("Last 20 areas")
    return nil

}
