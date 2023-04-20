package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "runtime"
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
        "map": {
            name: "map",
            description: "Display 20 new areas",
            callback: commandMap,
        },
        "mapb": {
            name: "mapb",
            description: "Display the last 20 areas",
            callback: commandMapb,
        },
    }
}


func clearTerminal() {
    var cmd *exec.Cmd

    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }

    cmd.Stdout = os.Stdout
    cmd.Run()
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

        clearTerminal()
        err := cliCommands[input].callback()
        if err != nil {
            fmt.Println("Error: ", err)
        }
    }
}
