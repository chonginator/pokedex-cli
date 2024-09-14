package main

import (
	"time"

	"github.com/chonginator/pokedex-cli/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{ pokeapiClient: pokeClient }

	startRepl(cfg)
}
