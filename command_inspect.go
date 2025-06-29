package main

import (
	"fmt"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func inspectCommand(c *config, cache pokecache.Cache, p Pokedex, pokemonName string) error {
	pokemon, exists := p[pokemonName]

	if !exists {
		fmt.Print("You haven't caught that pokemon!\n")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Base Experience: %v\n", pokemon.BaseExperience)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Print("Abilities:\n")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("	- %s\n", ability.Ability.Name)
	}
	fmt.Print("Forms:\n")
	for _, form := range pokemon.Forms {
		fmt.Printf("	- %s\n", form.Name)
	}
	fmt.Print("Moves:\n")
	for _, move := range pokemon.Moves {
		fmt.Printf("	- %s -- learns at level %v\n", move.Move.Name, move.VersionGroupDetails[0].LevelLearnedAt)
	}
	fmt.Print("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	- %s %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("	- %s\n", t.Type.Name)
	}

	return nil
}
