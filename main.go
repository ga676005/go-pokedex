package main

import (
	"time"

	"github.com/ga676005/pokedexcli/internal/pokeapi"
	"github.com/ga676005/pokedexcli/internal/store"
)

type config struct {
	pokeapiClient           pokeapi.Client
	previousLocationAreaURL *string
	nextLocationAreaURL     *string
	store                   store.Store
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		store:         store.NewStore(),
	}

	startRepl(&cfg)
}
