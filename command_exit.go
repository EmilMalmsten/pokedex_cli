package main

import (
    "fmt"
    "os"
)

func commandExit(appd *appData) error {
    fmt.Println("Exiting...")
    os.Exit(0)
    return nil
}
