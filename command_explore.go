package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]

	pokemonList, err := cfg.pokeapiClient.ListPokemon(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonList {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}