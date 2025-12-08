package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokeResult struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}

const (
	locationAPI string = "https://pokeapi.co/api/v2/location-area"
)

func PokeGet(url string) (PokeResult, error) {
	if len(url) == 0 {
		url = locationAPI
	}
	res, err := http.Get(url)
	if err != nil {
		return PokeResult{}, fmt.Errorf("couldn't GET %s: %w", url, err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var result PokeResult
	if err := decoder.Decode(&result); err != nil {
		return PokeResult{}, fmt.Errorf("couldn't deserialize: %w", err)
	}
	return result, nil
}
