package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	isPokemonCaught := rand.Intn(pokemon.BaseExperience) < 40

	fmt.Printf("Throwing a pokeball at %s...\n", pokemon.Name)
	if !isPokemonCaught {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command")

	cfg.pokedex[pokemon.Name] = pokemon

	return nil
}