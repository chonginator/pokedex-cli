package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for printPrompt(); scanner.Scan(); printPrompt() {
		input := scanner.Text()
		command, ok := cliCommands()[input]
		if !ok {
			fmt.Println("Unexpected command. Please try again!")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Printf("Error - %s: %v", command.name, err)
		}
	}
}

func printPrompt() {
	fmt.Printf("Pokedex > ")
}
