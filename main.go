package main

import (
	"github.com/alec-moore-se/pokedexcli/internal/pokeapi"
	"github.com/alec-moore-se/pokedexcli/internal/pokecache"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(10 * time.Second)
	cfg := &config{
		pokeCache:     pokeCache,
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
