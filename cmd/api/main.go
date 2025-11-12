package main

import (
	"fmt"

	"github.com/anmol420/bloom-filter-example/internal/env"
)

func main() {
	port := env.GetStringEnv("PORT")
	fmt.Println(port)
}