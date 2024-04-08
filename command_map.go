package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {
	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	for _, area := range response.Results {
		fmt.Printf(" - %s \n", area.Name)
	}

	cfg.nextLocationAreaURL = response.Next
	cfg.previousLocationAreaURL = response.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}

	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	for _, area := range response.Results {
		fmt.Printf(" - %s \n", area.Name)
	}

	cfg.nextLocationAreaURL = response.Next
	cfg.previousLocationAreaURL = response.Previous

	return nil
}
