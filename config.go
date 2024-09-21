package main

import "github.com/chonginator/pokedex-cli/internal/pokeapi"

type config struct{
	pokeapiClient pokeapi.Client
	pokedex map[string]pokeapi.Pokemon
	nextPageURL *string
	previousPageURL *string
}
