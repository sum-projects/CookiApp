package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sum-project/CookiApp/auth-service/cmd/db"
	"log"
)

func main() {
	config, err := LoadConfig("/app")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal("ping connect to db:", err)
	}

	store := db.NewStore(conn)
	server := NewServer(store)

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
