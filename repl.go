package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for printPrompt(); scanner.Scan(); printPrompt() {
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
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
	callback 		func() error
}

func getCommands() map[string]cliCommand {
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

func printPrompt() {
	fmt.Printf("Pokedex > ")
}

func cleanInput(input string) []string {
	inputLower := strings.ToLower(input)
	words := strings.Fields(inputLower)
	return words
}