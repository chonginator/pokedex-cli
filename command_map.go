package main

import (
	"fmt"
)

func commandMapForward(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextPageURL)
	if err != nil {
		return err
	}

	cfg.nextPageURL = locations.Next
	cfg.previousPageURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	if cfg.previousPageURL == nil {
		return fmt.Errorf("no previous results")
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.previousPageURL)
	if err != nil {
		return err
	}

	cfg.nextPageURL = locations.Next
	cfg.previousPageURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}