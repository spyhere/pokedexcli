package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := GetCommands()
	config := &Config{}
	Repl(config, scanner, commands)
}
