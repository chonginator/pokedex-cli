package pokeapi

import (
	"net/http"
	"time"

	"github.com/chonginator/pokedex-cli/internal/pokecache"
)

type Client struct{
	httpClient http.Client
	cache *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	pokecache := pokecache.NewCache(5 * time.Second)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache,
	}
}