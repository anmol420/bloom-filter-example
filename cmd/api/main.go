package main

import (
	"log"
	"os"

	"github.com/anmol420/bloom-filter-example/internal/env"
)

func main() {
	port := env.GetStringEnv("PORT")
	cfg := config{
		port: port,
	}
	app := &application{
		config: cfg,
	}
	if err := app.run(app.mount()); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}