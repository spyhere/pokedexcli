package main

import "github.com/spyhere/pokedexcli/internal/pokeClient"

type Config struct {
	Next     string
	Previous string
	Client   pokeClient.Client
	Pokedex  map[string]pokeClient.PokemonResponse
}
