package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, params ...string) error {

	names := cfg.store.GetAllPokemonNames()
	if len(names) == 0 {
		return errors.New("you have not caught any pokemon")
	}

	fmt.Println()
	fmt.Println("Your pokedex: ")
	for _, v := range names {
		fmt.Printf("  - %s", v)
	}
	fmt.Println()

	return nil
}
