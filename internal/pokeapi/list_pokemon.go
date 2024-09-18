package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListPokemon(locationAreaName string) ([]Resource, error) {
	URL := baseURL + "/location-area" + "/" + locationAreaName

	if val, ok := c.cache.Get(URL); ok {
		pokemon := []Resource{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return []Resource{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return []Resource{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []Resource{}, err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	locationArea := LocationArea{}
	err = decoder.Decode(&locationArea)
	if err != nil {
		return []Resource{}, err
	}
	pokemon := []Resource{}
	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		pokemon = append(pokemon, pokemonEncounter.Pokemon)
	}

	pokemonData, err := json.Marshal(pokemon)
	if err != nil {
		return []Resource{}, err
	}
	c.cache.Add(URL, pokemonData)

	return pokemon, nil
}