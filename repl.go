package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for printPrompt(); scanner.Scan(); printPrompt() {
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := commands[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command. Please try again!")
		}
	}
}

type cliCommand struct{
	name				string
	description string
	callback 		func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of the next 20 locations in the Pokemon world.",
			callback: commandMapForward,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of the previous 20 locations in the Pokemon world.",
			callback: commandMapBack,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Displays a list of all the Pokemon in a given location area.",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon_name>",
			description: "Attempt to catch a Pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon_name",
			description: "View details about a caught Pokemon",
			callback: commandInspect,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays the names of all the Pokemon you've caught",
			callback: commandPokedex,
		},
	}
}

func printPrompt() {
	fmt.Printf("Pokedex > ")
}

func cleanInput(input string) []string {
	inputLower := strings.ToLower(input)
	words := strings.Fields(inputLower)
	return words
}