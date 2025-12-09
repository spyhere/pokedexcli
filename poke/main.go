package poke

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spyhere/pokedexcli/internal/pokecache"
)

var API = api{
	LocationArea: uri{
		url: "https://pokeapi.co/api/v2/location-area",
	},
}
var cache = pokecache.NewCache(time.Second * 5)

func (u uri) Get(url string) (pokeResult, error) {
	if len(url) == 0 {
		url = u.url
	}
	cachedRes, ok := getFromCache(cache, url)
	if ok {
		return cachedRes, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return pokeResult{}, fmt.Errorf("couldn't GET %s: %w", url, err)
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeResult{}, fmt.Errorf("couldn't read response: %w", err)
	}
	addToCache(cache, url, bytes)

	var result pokeResult
	if err := json.Unmarshal(bytes, &result); err != nil {
		return pokeResult{}, fmt.Errorf("couldn't deserialize: %w", err)
	}
	return result, nil
}
