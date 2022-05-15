package main

import (
	"fmt"
	"os"
	"os/user"
	"sprite/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome %s to Sprite, a small language with a humble beginning!\n", user.Username)
	fmt.Printf("Type your commands here.\n")
	// Starting Sprite REPL
	repl.Start(os.Stdin, os.Stdout)
}
