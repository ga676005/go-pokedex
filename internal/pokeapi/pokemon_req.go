package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(idOrName string) (PokemonResponse, error) {
	endPoint := "/pokemon/" + idOrName
	fullURL := baseURL + endPoint

	rawData, ok := c.cahce.Get(fullURL)
	if ok {
		pokemonResponse := PokemonResponse{}
		json.Unmarshal(rawData, &pokemonResponse)
		return pokemonResponse, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return PokemonResponse{}, fmt.Errorf("bad status code %v", response.StatusCode)
	}

	rawData, err = io.ReadAll(response.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemonResponse := PokemonResponse{}
	err = json.Unmarshal(rawData, &pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cahce.Add(fullURL, rawData)

	return pokemonResponse, nil
}
