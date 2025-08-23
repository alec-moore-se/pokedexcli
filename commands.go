package main

import (
	"fmt"
	"github.com/alec-moore-se/pokedexcli/internal/pokeapi"
	"math/rand"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config) error {
	locationsRes, err := cfg.pokeapiClient.ListPokemonInLocation(cfg.nextLocationsURL, cfg.additionalPrompts[0])
	if err != nil {
		return fmt.Errorf("Error Occured: %w", err)
	}
	fmt.Printf("Exploring %s\n", cfg.additionalPrompts[0])
	fmt.Println("Found Pokemon: ")
	for _, pokeEnc := range locationsRes.PokemonEncounters {
		fmt.Println(pokeEnc.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config) error {
	pokemonInQuestion := cfg.additionalPrompts[0]
	pokemonRes, err := cfg.pokeapiClient.ListPokemonAttributes(pokemonInQuestion)
	if err != nil {
		return err
	}
	magicNumber := 40
	percent := 100
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonInQuestion)
	percentChanceOfCatching := float64((float64(magicNumber) / float64(pokemonRes.BaseExperience)) * float64(percent))
	actualCatch := rand.Int() % percent
	if percentChanceOfCatching > float64(actualCatch) {
		fmt.Printf("%s was caught!\n", pokemonInQuestion)
		cfg.storageBox[pokemonInQuestion] = pokeapi.PokemonStatsToReduced(pokemonRes)
	} else {
		fmt.Printf("%s escaped!\n", pokemonInQuestion)
	}
	return nil
}

func commandInspect(cfg *config) error {
	poke, ok := cfg.storageBox[cfg.additionalPrompts[0]]
	if ok {
		pokeapi.PrintPokemonStats(poke)
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func commandPokedex(cfg *config) error {
	fmt.Println("Your Pokedex:")
	for _, poke := range cfg.storageBox {
		fmt.Printf("\t-%s\n", poke.Name)
	}
	return nil
}
