package main

import (
	"log"

	"github.com/igitur/hoauth/cmd"
)

func main() {
	// No timestamps
	log.SetFlags(0)

	cmd.Execute()
}
