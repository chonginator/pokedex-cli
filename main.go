package main

import (
	"time"

	"github.com/chonginator/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	pokedex := map[string]pokeapi.Pokemon{}
	cfg := &config{ pokeapiClient: pokeClient, pokedex: pokedex }

	startRepl(cfg)
}
