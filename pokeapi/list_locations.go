package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PaginatedResourceList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PaginatedResourceList{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PaginatedResourceList{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return PaginatedResourceList{}, fmt.Errorf("unexpected non-OK status code: %v", res.StatusCode)
	}

	var locations PaginatedResourceList
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return PaginatedResourceList{}, err
	}

	return locations, nil
}