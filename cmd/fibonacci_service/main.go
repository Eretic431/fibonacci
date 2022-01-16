package main

import (
	"log"
)

func main() {
	s, closer, err := initServer()
	if err != nil {
		log.Fatalf("could not init server: %v", err)
	}
	closer()
	s.Run()
}
