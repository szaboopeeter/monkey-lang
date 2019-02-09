package main

import (
	"fmt"
	"monkey-lang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Monkey-lang!\n", user.Username)
	fmt.Printf("Feel free to type commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
