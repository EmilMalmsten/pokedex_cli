package main

import (
    "bufio"
    "fmt"
    "os"
)

type cliCommand struct {
    name string
    description string
    callback func() error
}

func getCliCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
    }
}

func commandHelp() error {
    fmt.Println("here is some help")
    return nil
}

func commandExit() error {
    fmt.Println("Exiting...")
    os.Exit(0)
    return nil
}

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)
    cliCommands := getCliCommands()

    for {

        fmt.Println("")
        fmt.Println("Welcome to the pokedex!")
        fmt.Println("Usage:")
        fmt.Println("")

        for command := range cliCommands {
            fmt.Println(cliCommands[command].name)
        }

        fmt.Println("")
        fmt.Print("Pokedex > ")

        scanner.Scan()
        input := scanner.Text()

        err := cliCommands[input].callback()
        if err != nil {
            fmt.Println("Error: ", err)
        }
    }
}
