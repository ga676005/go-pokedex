package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("please provide the name or id of the pokemon")
	}
	idOrName := params[0]
	response, err := cfg.pokeapiClient.GetPokemon(idOrName)
	if err != nil {
		return err
	}

	successRate := float32(rand.Intn(response.BaseExperience)) / float32(response.BaseExperience)
	isCatch := successRate > 0.7

	fmt.Println("successRate", successRate)

	fmt.Printf("Throwing a Pokeball at %s... \n", response.Name)

	if isCatch {
		fmt.Printf("%s was caught! \n", response.Name)
		cfg.store.Add(response.Name, response)
	} else {
		fmt.Printf("%s escaped! \n", response.Name)
	}

	return nil
}
