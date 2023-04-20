package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	status := client.Ping(context.Background())
	if status.Err() != nil {
		log.Fatal("ping connect to redis", status.Err())
	}

	store := db.NewStore(conn, client)
	server := NewServer(store, config)

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
