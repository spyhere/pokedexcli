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
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			log.Fatal("failed reading the text")
		}
		first_word := CleanInput(scanner.Text())[0]
		command, ok := commands[first_word]
		if !ok {
			commands["help"].cb()
			continue
		}
		command.cb()
	}
}
