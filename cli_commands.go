package main

import (
	"fmt"
	"os"
)

type cliCommand struct{
	name				string
	description string
	callback 		func() error
}

func cliCommands() map[string]cliCommand {
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
	fmt.Printf(
`
Welcome to the Pokedex!
Usage:

`)
	for name, cliCommand := range cliCommands() {
		fmt.Printf(
`
%s: %s
`, name, cliCommand.description,
		)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}