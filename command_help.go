package main

import (
	"fmt"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func helpCommand(c *config, cache pokecache.Cache, p Pokedex, location string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
