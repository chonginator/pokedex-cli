package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return nil
	}
	areaName := args[0]

	pokemonList, err := cfg.pokeapiClient.ListPokemon(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonList {
		fmt.Printf("- %v\n", pokemon.Name)
	}
	return nil
}