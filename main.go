package main

import (
	"time"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokeapi"
	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	c := &config{
		pokeapiClient: pokeClient,
	}
	cache := pokecache.NewCache(5 * time.Second)

	startRepl(c, cache)
}
