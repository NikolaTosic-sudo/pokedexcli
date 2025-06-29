package main

import (
	"fmt"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func MapCommand(c *config, cache pokecache.Cache, p Pokedex, location string) error {

	if c.Next == nil {
		url := "https://pokeapi.co/api/v2/location-area"
		c.Next = &url
	}

	areas, err := c.pokeapiClient.ListLocations(c.Next, cache)

	if err != nil {
		return err
	}

	c.Next = areas.Next

	if areas.Previous != nil {
		c.Previous = areas.Previous
	}

	for _, area := range areas.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}

func MapBackCommand(c *config, cache pokecache.Cache, p Pokedex, location string) error {

	if c.Previous == nil {

		fmt.Print("you're on the first page\n")

		return nil
	}

	areas, err := c.pokeapiClient.ListLocations(c.Previous, cache)

	if err != nil {
		return err
	}

	c.Next = areas.Next

	if areas.Previous != nil {
		c.Previous = areas.Previous
	}

	for _, area := range areas.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}
