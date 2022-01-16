package main

import (
	"github.com/xlab/closer"
	"log"
)

func main() {
	s, cleanup, err := initServer()
	if err != nil {
		log.Fatalf("could not init server: %v", err)
	}
	closer.Bind(func() {
		s.Log.Infof("Stopping server")
		cleanup()
	})
	s.Run()
}
