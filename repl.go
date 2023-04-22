package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display 20 new areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the last 20 areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the pokemon within an area",
			callback:    commandExplore,
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

func startRepl(cfg *config) {
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
		words := cleanInput(input)

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := cliCommands[commandName]
		if exists {
			clearTerminal()
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}

	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}
