package poke

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var API = api{
	LocationArea: uri{
		url: "https://pokeapi.co/api/v2/location-area",
	},
}

func (u uri) Get(url string) (pokeResult, error) {
	if len(url) == 0 {
		url = u.url
	}
	res, err := http.Get(url)
	if err != nil {
		return pokeResult{}, fmt.Errorf("couldn't GET %s: %w", url, err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var result pokeResult
	if err := decoder.Decode(&result); err != nil {
		return pokeResult{}, fmt.Errorf("couldn't deserialize: %w", err)
	}
	return result, nil
}
