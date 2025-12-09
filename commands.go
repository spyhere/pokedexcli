package main

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
	}
}
