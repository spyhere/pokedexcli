package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	cb          func(*Config, ...string) error
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
		"explore": {
			name:        "explore",
			description: "Let you explore pokemons in specific location area",
			cb:          commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch the pokemon by name",
			cb:          commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect the stats of the pokemon in your possession",
			cb:          commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display pokemons that you have caught",
			cb:          commandPokedex,
		},
	}
}

func commandExit(c *Config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, _ ...string) error {
	commandDescription := ""
	for _, it := range GetCommands() {
		commandDescription += fmt.Sprintf("%s: %s\n", it.name, it.description)
	}
	fmt.Printf(`Welcome to the Pokedex!
Usage:

%s`, commandDescription)
	return nil
}

func commandMap(c *Config, _ ...string) error {
	if len(c.Next) == 0 && len(c.Previous) != 0 {
		fmt.Println("You are on the last page")
		return nil
	}
	res, err := c.Client.GetLocationAreas(c.Next)
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

func commandMapb(c *Config, _ ...string) error {
	if len(c.Previous) == 0 {
		fmt.Println("You are on the first page")
		return nil
	}
	res, err := c.Client.GetLocationAreas(c.Previous)
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

func commandExplore(c *Config, args ...string) error {
	location := args[0]
	if location == "" {
		return fmt.Errorf("Didn't receive any location")
	}
	res, err := c.Client.GetLocationArea(location)
	if err != nil {
		return err
	}
	foundPokemon := ""
	for _, pokemonEncounters := range res.PokemonEncounters {
		foundPokemon += "- " + pokemonEncounters.Pokemon.Name + "\n"
	}
	fmt.Printf("Found Pokemon:\n%s", foundPokemon)
	return nil
}

func commandCatch(c *Config, args ...string) error {
	pokemonName := args[0]
	if pokemonName == "" {
		return fmt.Errorf("Didn't receive any pokemon name")
	}
	res, err := c.Client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	if _, ok := c.Pokedex[res.Name]; ok {
		fmt.Printf("You already have %s in your Pokedex!\n", res.Name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)
	random := rand.Intn(res.BaseExperience)
	catchThreshold := int(math.Min(50.0, float64(res.BaseExperience)*0.8))
	if random <= catchThreshold {
		fmt.Println(res.Name, "was caught!")
		c.Pokedex[res.Name] = res
	} else {
		fmt.Println(res.Name, "escaped!")
	}
	return nil
}

func commandInspect(c *Config, args ...string) error {
	pokemonName := args[0]
	if pokemonName == "" {
		fmt.Printf("No pokemon name were given!")
		return nil
	}
	pokemon, ok := c.Pokedex[pokemonName]
	if !ok {
		fmt.Printf("You haven't caught %s yet!\n", pokemonName)
		return nil
	}
	fmt.Print(GetPokemonString(pokemon))
	return nil
}

func commandPokedex(c *Config, args ...string) error {
	if len(c.Pokedex) == 0 {
		fmt.Println("You haven't caught a single pokemon yet!")
		return nil
	}
	output := ""
	for key := range c.Pokedex {
		output += " -" + key + "\n"
	}
	fmt.Printf("Your Pokedex:\n%s", output)
	return nil
}
