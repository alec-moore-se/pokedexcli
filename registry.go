package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand = makeMapCommands()
commands["help"] = cliCommand{
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		}


func makeMapCommands() map[string]cliCommand {
	returner := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return returner
}

func getCommand(s string) cliCommand {
	el, exists := commands[s]
	if exists {
		return el
	}
	return cliCommand{}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	for _, com := range commands {
		fmt.Printf("%s: %s", com.name, com.description)
	}
	return nil
}
