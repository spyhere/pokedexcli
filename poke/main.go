package poke

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spyhere/pokedexcli/internal/pokecache"
)

var API = Api{
	LocationArea: "https://pokeapi.co/api/v2/location-area",
}
var cache = pokecache.NewCache(time.Second * 5)

func Get[T pokeResult](url string) (T, error) {
	var res T
	bytes, err := httpGet(url)
	if err != nil {
		return res, err
	}
	if err := json.Unmarshal(bytes, &res); err != nil {
		return res, fmt.Errorf("couldn't deserialize: %w", err)
	}
	return res, nil
}

func httpGet(url string) ([]byte, error) {
	cachedRes, ok := cache.Get(url)
	if ok != nil {
		return cachedRes, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("couldn't GET %s: %w", url, err)
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("couldn't read response: %w", err)
	}
	cache.Add(url, bytes)
	return bytes, nil
}
