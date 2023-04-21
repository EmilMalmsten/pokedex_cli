package main

import (
    "fmt"
    "github.com/emilmalmsten/pokedex_cli/internal/pokeapi"
    "log"
)

func commandMap(appd *appData) error {

    locationAreasResp, err := pokeapi.ListLocationAreas(appd.nextLocationAreaUrl)
    if err != nil {
        log.Fatal(err)
    }
    for _, area := range locationAreasResp.Results {
        fmt.Println(area.Name)
    }
    appd.nextLocationAreaUrl = locationAreasResp.Next
    appd.prevLocationAreaUrl = locationAreasResp.Previous
    return nil

}

func commandMapb(appd *appData) error {

    locationAreasResp, err := pokeapi.ListLocationAreas(appd.prevLocationAreaUrl)
    if err != nil {
        log.Fatal(err)
    }
    for _, area := range locationAreasResp.Results {
        fmt.Println(area.Name)
    }
    appd.nextLocationAreaUrl = locationAreasResp.Next
    appd.prevLocationAreaUrl = locationAreasResp.Previous
    return nil

}
