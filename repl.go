package main

import (
	"fmt"
	"github.com/spyhere/pokedexcli/poke"
	"os"
	"strings"
)

type Config struct {
	Next     string
	Previous string
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
	if len(c.Next) == 0 && len(c.Previous) != 0 {
		fmt.Println("You are on the last page")
		return nil
	}
	if c.Next == "" {
		c.Next = poke.API.LocationArea
	}
	res, err := poke.Get[poke.LocationsPaginatedResult](c.Next)
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
	res, err := poke.Get[poke.LocationsPaginatedResult](c.Previous)
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
