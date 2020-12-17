package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/rtmoranorg/vanilla/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Welcome to Vanilla, %s\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
