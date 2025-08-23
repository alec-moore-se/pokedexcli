package main

import (
	"bufio"
	"fmt"
	"github.com/alec-moore-se/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient     pokeapi.Client
	nextLocationsURL  *string
	prevLocationsURL  *string
	additionalPrompts []string
	storageBox        map[string]pokeapi.PokemonStatsReduced
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		cfg.additionalPrompts = words[1:]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Lists pokemons found in specified area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the specified pokemon, from next line input",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Shows the stats of a caught Pokemon, does not if Pokemon is not listed",
			callback:    commandInspect,
		},
	}
}
