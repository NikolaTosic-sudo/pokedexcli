package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func (c *Client) ListPokemons(location string, cache pokecache.Cache) (PokemonEncounters, error) {
	var pokemons PokemonEncounters

	url := baseURL + "/location-area/" + location

	if data, exists := cache.GetCache(url); exists {
		if err := json.Unmarshal(data, &pokemons); err != nil {
			return PokemonEncounters{}, err
		}
	} else {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return PokemonEncounters{}, err
		}

		res, err := c.httpClient.Do(req)

		if err != nil {
			return PokemonEncounters{}, err
		}

		if res.StatusCode > 299 {
			return PokemonEncounters{}, fmt.Errorf("there was an issue with the request")
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return PokemonEncounters{}, err
		}

		cache.AddCache(url, data)

		if err := json.Unmarshal(data, &pokemons); err != nil {
			return PokemonEncounters{}, err
		}
	}

	return pokemons, nil
}
