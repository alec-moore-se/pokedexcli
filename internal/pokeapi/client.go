package pokeapi

import (
	"github.com/alec-moore-se/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient  http.Client
	clientCache *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		clientCache: pokecache.NewCache(timeout),
	}
}
