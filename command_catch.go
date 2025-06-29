package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func catchCommand(c *config, cache pokecache.Cache, p Pokedex, pokemonName string) error {

	if _, exists := p[pokemonName]; exists {
		fmt.Print("You already caught " + pokemonName)
		return nil
	}

	pokemon, err := c.pokeapiClient.PokemonDetails(pokemonName, cache)

	if err != nil {
		return err
	}

	fmt.Print("Throwing a Pokeball at " + pokemon.Name + "...\n")

	random := rand.Intn(101)

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if random > 30 {
			fmt.Print(pokemon.Name + " was caught!\n")
			p[pokemonName] = pokemon
			return nil
		}
		fmt.Print(pokemon.Name + " escaped!\n")
	}

	return nil
}
