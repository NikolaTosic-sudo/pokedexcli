package main

import (
	"time"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	c := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(c)
}
