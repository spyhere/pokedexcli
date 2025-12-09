package main

import (
	"bufio"
	"os"
	"time"

	"github.com/spyhere/pokedexcli/internal/pokeclient"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := GetCommands()
	config := &Config{
		Client: *pokeclient.Init(
			"https://pokeapi.co/api/v2/",
			time.Second*10,
		),
	}
	Repl(config, scanner, commands)
}
