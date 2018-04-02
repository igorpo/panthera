package main

import (
	"fmt"
	"os"
	"os/user"
	"panthera/repl"

	log "github.com/sirupsen/logrus"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("os error creating a user")
	}

	fmt.Printf("Hello %s, welcome to Panthera programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
