package main

import (
	"bufio"
	"fmt"
	"log"
)

func Repl(c *Config, scanner *bufio.Scanner, commands map[string]cliCommand) {
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
			commands["help"].cb(c, args...)
			continue
		}
		err := command.cb(c, args...)
		if err != nil {
			log.Fatal(err)
		}
	}
}
