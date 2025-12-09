package pokeClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	pokecache "github.com/spyhere/pokedexcli/internal/pokeCache"
)

type Client struct {
	baseUrl string
	cache   pokecache.Cache
}

func Init(baseUrl string, cacheInterval time.Duration) *Client {
	return &Client{
		baseUrl: baseUrl,
		cache:   *pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) GetLocationAreas(url string) (LocationsPaginatedResult, error) {
	if url == "" {
		url = c.baseUrl + "location-area/"
	}
	bytes, err := httpGet(url, &c.cache)
	if err != nil {
		return LocationsPaginatedResult{}, err
	}
	var res LocationsPaginatedResult
	if err := json.Unmarshal(bytes, &res); err != nil {
		return LocationsPaginatedResult{}, err
	}
	return res, nil
}

func (c *Client) GetLocationArea(locationName string) (LocationResult, error) {
	url := c.baseUrl + "location-area/" + locationName
	bytes, err := httpGet(url, &c.cache)
	if err != nil {
		return LocationResult{}, err
	}
	var res LocationResult
	if err := json.Unmarshal(bytes, &res); err != nil {
		return LocationResult{}, err
	}
	return res, nil
}

func (c *Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	url := c.baseUrl + "pokemon/" + pokemonName
	bts, err := httpGet(url, &c.cache)
	if err != nil {
		return PokemonResponse{}, err
	}
	var res PokemonResponse
	bReader := bytes.NewReader(bts)
	decoder := json.NewDecoder(bReader)
	if err := decoder.Decode(&res); err != nil {
		return PokemonResponse{}, err
	}
	return res, nil
}

func httpGet(url string, cache *pokecache.Cache) ([]byte, error) {
	cachedRes, err := cache.Get(url)
	if err == nil {
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
