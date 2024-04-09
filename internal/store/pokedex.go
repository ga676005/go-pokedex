package store

import "github.com/ga676005/pokedexcli/internal/pokeapi"

type Pokemon = pokeapi.PokemonResponse

type Store struct {
	pokedex map[string]Pokemon
}

func NewStore() Store {
	return Store{
		pokedex: make(map[string]Pokemon),
	}
}

func (s *Store) Add(name string, pokemon Pokemon) {
	s.pokedex[name] = pokemon
}

func (s *Store) Get(name string) (Pokemon, bool) {
	p, exist := s.pokedex[name]
	return p, exist
}

func (s *Store) GetAllPokemonNames() []string {
	names := make([]string, len(s.pokedex))
	i := 0
	for k, _ := range s.pokedex {
		names[i] = k
		i++
	}
	return names
}
