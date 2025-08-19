package main

import (
	"fmt"
	"github.com/alec-moore-se/pokedexcli/internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Current  string
	Next     string
	Previous string
}

func makeMapCommands() map[string]cliCommand {
	returner := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets 20 locations areas for current page, increments page",
			callback:    commandMapShow,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets 20 locations areas for previous page, decrements page",
			callback:    commandMapShowBack,
		},
	}
	return returner
}

func getCommand(s string) (cliCommand, error) {
	el, exists := makeMapCommands()[s]
	if exists {
		return el, nil
	}
	return cliCommand{}, fmt.Errorf("Unknown Command")
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	for _, com := range makeMapCommands() {
		fmt.Printf("%s: %s\n", com.name, com.description)
	}
	return nil
}

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = locationsResp.Next
	cfg.Previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "null" {
		return fmt.Errorf("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
