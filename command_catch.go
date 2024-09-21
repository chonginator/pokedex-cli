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

	const (
		minimumCatchRateDifficulty = 0.02
		maximumCatchRateDifficulty = 0.95
		maximumBaseExperience = 340
		minimumBaseExperience = 39
	)
	normalizedBaseExperience := 
		(float32(pokemon.BaseExperience - minimumBaseExperience)) /
		(maximumBaseExperience - minimumBaseExperience)
	catchChance := 
		minimumCatchRateDifficulty +
		(maximumCatchRateDifficulty - minimumCatchRateDifficulty) *
		(1 - normalizedBaseExperience)
	isPokemonCaught := rand.Float32() <= catchChance

	fmt.Printf("Throwing a pokeball at %s...\n", pokemon.Name)
	if !isPokemonCaught {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	if _, ok := cfg.pokedex[pokemon.Name]; !ok {
		cfg.pokedex[pokemon.Name] = pokemon
	}

	return nil
}