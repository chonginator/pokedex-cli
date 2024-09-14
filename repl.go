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
		command, exists := commands[commandName]
		if exists {
			err := command.callback(cfg)
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
	callback 		func(*config) error
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
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
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