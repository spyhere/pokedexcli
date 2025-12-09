package poke

import (
	"encoding/json"
	"log"

	"github.com/spyhere/pokedexcli/internal/pokecache"
)

func getFromCache(cache *pokecache.Cache, key string) (pokeResult, bool) {
	if val, err := cache.Get(key); err == nil {
		var res pokeResult
		if err := json.Unmarshal(val, &res); err != nil {
			log.Fatal("coudln't deserialize: %w", err)
		}
		return res, true
	}
	return pokeResult{}, false
}

func addToCache(cache *pokecache.Cache, key string, value []byte) {
	cache.Add(key, value)
}
