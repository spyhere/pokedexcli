package main

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	cb          func(*Config) error
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
		"map": {
			name:        "map",
			description: "Show next 20 locations",
			cb:          commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 locations",
			cb:          commandMapb,
		},
	}
}

func CleanInput(text string) []string {
	return strings.Split(
		strings.ToLower(strings.TrimSpace(text)),
		" ",
	)
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	commandDescription := ""
	for _, it := range GetCommands() {
		commandDescription += fmt.Sprintf("%s: %s\n", it.name, it.description)
	}
	fmt.Printf(`Welcome to the Pokedex!
Usage:

%s`, commandDescription)
	return nil
}

func commandMap(c *Config) error {
	if len(c.Next) == 0 {
		fmt.Println("You are on the last page")
		return nil
	}
	res, err := PokeGet(c.Next)
	if err != nil {
		return err
	}
	c.Next = res.Next
	c.Previous = res.Previous
	for _, it := range res.Results {
		fmt.Println(it.Name)
	}
	return nil
}

func commandMapb(c *Config) error {
	if len(c.Previous) == 0 {
		fmt.Println("You are on the first page")
		return nil
	}
	res, err := PokeGet(c.Previous)
	if err != nil {
		return err
	}
	c.Next = res.Next
	c.Previous = res.Previous
	for _, it := range res.Results {
		fmt.Println(it.Name)
	}
	return nil
}
