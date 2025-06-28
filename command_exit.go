package main

import (
	"fmt"
	"os"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func exitCommand(c *config, cache pokecache.Cache) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
