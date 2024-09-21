package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	URL := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(URL); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)

		if err != nil {
			return Pokemon{}, nil
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return Pokemon{}, errors.New("pokemon does not exist")
	}
	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("unexpected non-OK status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, nil
	}

	c.cache.Add(URL, data)

	return pokemon, nil
}