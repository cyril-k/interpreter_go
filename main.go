package main

import (
	"fmt"
	"macaque/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, this is the Macaque programming language\n",
		user.Username)
	fmt.Printf("This is a test of REPL\n")
	repl.Start(os.Stdin, os.Stdout)
}
