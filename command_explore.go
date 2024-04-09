package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("please provide the name or id of the location")
	}
	idOrName := params[0]
	response, err := cfg.pokeapiClient.GetLocationDetail(idOrName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s \n", response.Name)
	for _, v := range response.PokemonEncounters {
		fmt.Printf("%s \n", v.Pokemon.Name)
	}

	return nil
}
