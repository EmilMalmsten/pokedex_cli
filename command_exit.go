package main

import (
    "fmt"
    "os"
)

func commandExit() error {
    fmt.Println("Exiting...")
    os.Exit(0)
    return nil
}
