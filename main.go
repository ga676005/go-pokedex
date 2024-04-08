package main

import (
	"time"

	"github.com/ga676005/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	previousLocationAreaURL *string
	nextLocationAreaURL     *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&cfg)
}
