package main

import (
	"fmt"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func pokedexCommand(c *config, cache pokecache.Cache, p Pokedex, pokemonName string) error {
	if len(p) == 0 {
		fmt.Print("Your pokedex is empty... Go catch some pokemons!\n")
		return nil
	}
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range p {
		fmt.Printf("	- %s\n", pokemon.Name)
	}

	return nil
}
