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

	// pokecache := pokecache.NewCache(5 * time.Second)
	val, existsInCache := c.cache.Get(url)
	if existsInCache {
		fmt.Println("returning cached response!")
		var locations PaginatedResourceList
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return PaginatedResourceList{}, err
		}
		return locations, nil
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

	if !existsInCache {
		locationsData, err := json.Marshal(locations)
		if err != nil {
			return PaginatedResourceList{}, err
		}
		c.cache.Add(url, locationsData)
	}

	return locations, nil
}