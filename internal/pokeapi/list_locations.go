package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocations(pageUrl *string, cache pokecache.Cache) (LocationArea, error) {
	var areas LocationArea

	url := baseURL + "/location-area?offset=0&limit=20"
	if pageUrl != nil {
		url = *pageUrl
	}

	if data, exists := cache.GetCache(url); exists {
		if err := json.Unmarshal(data, &areas); err != nil {
			return LocationArea{}, err
		}
	} else {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return LocationArea{}, err
		}

		res, err := c.httpClient.Do(req)

		if err != nil {
			return LocationArea{}, err
		}

		if res.StatusCode > 299 {
			return LocationArea{}, fmt.Errorf("there was an issue with the request")
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}

		cache.AddCache(url, data)

		if err := json.Unmarshal(data, &areas); err != nil {
			return LocationArea{}, err
		}
	}

	return areas, nil
}
