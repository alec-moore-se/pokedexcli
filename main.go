package main

import (
	"github.com/alec-moore-se/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		storageBox:    make(map[string]pokeapi.PokemonStatsReduced),
	}

	startRepl(cfg)
}
