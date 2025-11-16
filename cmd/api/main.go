package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/anmol420/bloom-filter-example/internal/db"
	"github.com/anmol420/bloom-filter-example/internal/env"
	"github.com/anmol420/bloom-filter-example/internal/store"
)

func main() {
	port := env.GetStringEnv("PORT")
	dbURI := env.GetStringEnv("DB_URI")
	dbName := env.GetStringEnv("DB_NAME")
	cfg := config{
		port: port,
		db: dbConfig{
			dbURI: dbURI,
			dbName: dbName,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, db, err := db.New(ctx, cfg.db.dbURI, cfg.db.dbName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()
	log.Println("Database Connection Established!")
	if err := store.UserUniqueIndex(ctx, db.Collection("users")); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	store := store.NewMongoStorage(client, dbName)
	app := &application{
		config: cfg,
		store: store,
	}
	if err := app.run(app.mount()); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}