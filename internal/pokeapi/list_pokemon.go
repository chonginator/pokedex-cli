package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationAreaName string) (LocationArea, error) {
	URL := baseURL + "/location-area/" + locationAreaName

	if val, ok := c.cache.Get(URL); ok {
		pokemon := LocationArea{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return LocationArea{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return LocationArea{}, errors.New("location area does not exist")
	}
	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("non-OK status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, nil
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, nil
	}

	c.cache.Add(URL, data)

	return locationArea, nil
}