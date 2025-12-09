package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := GetCommands()
	config := &Config{}
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			log.Fatal("failed reading the text")
		}
		words := CleanInput(scanner.Text())
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, ok := commands[words[0]]
		if !ok {
			commands["help"].cb(config, args...)
			continue
		}
		err := command.cb(config, args...)
		if err != nil {
			log.Fatal("unexpected: %w", err)
		}
	}
}
