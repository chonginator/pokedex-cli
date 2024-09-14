package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIResourceList struct{
	Count int 						`json:"count"`
	Next string 					`json:"next"`
	Previous string 			`json:"previous"`
	Results []APIResource `json:"results"`
}

type APIResource struct{
	Name string `json:"name"`
	URL string	`json:"url"`
}

func commandMap(config *Config) error {
	if config.NextURL == nil {
		nextURL := "https://pokeapi.co/api/v2/location?limit=20"
		config.NextURL = &nextURL
	}
	if *config.NextURL == "" {
		return fmt.Errorf("no more results")
	}

	res, err := http.Get(*config.NextURL)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("unexpected non-OK status code: %v", res.StatusCode)
	}

	var locations APIResourceList
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return fmt.Errorf("failed to decode map locations: %w", err)
	}
	config.NextURL = &locations.Next
	config.PreviousURL = &locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(config *Config) error {
	if config.PreviousURL == nil || *config.PreviousURL == "" {
		return fmt.Errorf("no previous results")
	}

	res, err := http.Get(*config.PreviousURL)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("unexpected non-OK status code: %v", res.StatusCode)
	}

	var locations APIResourceList
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return fmt.Errorf("failed to decode map locations: %w", err)
	}
	config.NextURL = &locations.Next
	config.PreviousURL = &locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}