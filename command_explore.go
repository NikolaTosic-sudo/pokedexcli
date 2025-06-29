package main

import (
	"fmt"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func exploreCommand(c *config, cache pokecache.Cache, p Pokedex, location string) error {

	pokemons, err := c.pokeapiClient.ListPokemons(location, cache)

	if err != nil {
		return err
	}

	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}

	return nil
}
