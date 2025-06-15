package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/lucianocorreia/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! This is the Monkey programing language \n", user.Username)
	fmt.Printf("Type in commands to get started. Use Ctrl+C to exit.\n")
	repl.Start(os.Stdin, os.Stdout)
}
