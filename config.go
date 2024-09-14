package main

import "github.com/chonginator/pokedex-cli/pokeapi"

type config struct{
	pokeapiClient pokeapi.Client
	nextPageURL *string
	previousPageURL *string
}
