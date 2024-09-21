package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	URL := baseURL + "/pokemon/" + name
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}

	if res.StatusCode == 404 {
		return Pokemon{}, errors.New("pokemon does not exist")
	}
	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("unexpected non-OK status code: %v", res.StatusCode)
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var pokemon Pokemon
	err = decoder.Decode(&pokemon)
	if err != nil {
		return Pokemon{}, nil
	}
	return pokemon, nil
}