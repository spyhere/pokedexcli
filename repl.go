package main

import (
	"bufio"
	"fmt"
	"log"
)

func Repl(c *Config, scanner *bufio.Scanner, commands map[string]CliCommand) {
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			log.Fatal("failed reading the text")
		}
		words := CleanInput(scanner.Text())
		err := RunCommand(c, words, commands)
		if err != nil {
			log.Fatal(err)
		}
	}
}
