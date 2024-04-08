package pokeapi

import (
	"net/http"
	"time"

	"github.com/ga676005/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cahce: pokecache.NewCache(cacheInterval),
	}
}

type Client struct {
	httpClient http.Client
	cahce      pokecache.Cache
}
