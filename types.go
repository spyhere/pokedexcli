package main

import "github.com/spyhere/pokedexcli/internal/pokeClient"

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  struct {
		HP             int
		Attack         int
		Defense        int
		SpecialAttack  int
		SpecialDefense int
		Speed          int
	}
	Types []string
}

type Config struct {
	Next     string
	Previous string
	Client   pokeClient.Client
	Pokedex  map[string]Pokemon
}
