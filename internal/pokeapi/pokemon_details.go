package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func (c *Client) PokemonDetails(pokemon string, cache pokecache.Cache) (Pokemon, error) {
	var pokemonDetails Pokemon

	url := baseURL + "/pokemon/" + pokemon

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		fmt.Print("No pokemon with that name")
		return Pokemon{}, err
	}

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("there was an issue with the request")
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if err := json.Unmarshal(data, &pokemonDetails); err != nil {
		return Pokemon{}, err
	}

	return pokemonDetails, nil
}
