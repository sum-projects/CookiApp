package main

import (
	"log"
)

const webPort = "80"

func main() {
	config, err := LoadConfig("/app")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server := NewServer(config)
	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
