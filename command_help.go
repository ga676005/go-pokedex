package main

import "fmt"

func callbackHelp(cfg *config, params ...string) error {
	fmt.Println("Welcome to pokedex help menu")
	avaliableCommands := getCommands()
	for _, c := range avaliableCommands {
		fmt.Printf(" - %s: %s\n", c.name, c.description)
	}
	fmt.Println()

	return nil
}
