package main

import (
	"log"
)

func main() {
	config, err := LoadConfig("/app")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	mailer := NewMail(
		config.Domain,
		config.Host,
		config.Username,
		config.Password,
		config.Encryption,
		config.FromName,
		config.FromAddress,
		config.Port,
	)
	server := NewServer(mailer)

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
