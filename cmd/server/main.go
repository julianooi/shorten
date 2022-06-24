package main

import (
	"fmt"
	"github.com/julianooi/shorten/server"
	"log"
)

func main() {
	err := start()
	if err != nil {
		log.Fatal(err)
	}
}

func start() error {
	svr := server.NewServer()
	err := svr.Start()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}
