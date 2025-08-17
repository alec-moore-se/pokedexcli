package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	scnr := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scnr.Scan()
		userInput := scnr.Text()
		userInputs := cleanInput(userInput)
		command := getCommand(userInputs[0])

	}
}
