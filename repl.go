package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	cb          func() error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			cb:          commandExit,
		},
		"help": {
			name:        "help",
			description: "Read help of Pokedex",
			cb:          commandHelp,
		},
	}
}

func CleanInput(text string) []string {
	return strings.Split(
		strings.ToLower(strings.TrimSpace(text)),
		" ",
	)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commandDescription := ""
	for _, it := range GetCommands() {
		commandDescription += fmt.Sprintf("%s: %s\n", it.name, it.description)
	}
	fmt.Printf(`Welcome to the Pokedex!
Usage:

%s`, commandDescription)
	return nil
}
