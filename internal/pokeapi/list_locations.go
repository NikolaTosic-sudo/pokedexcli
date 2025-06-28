package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

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

	var areas LocationArea

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	if err := json.Unmarshal(data, &areas); err != nil {
		return LocationArea{}, err
	}

	return areas, nil
}
