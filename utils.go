package main

import (
	"fmt"
	"strings"

	"github.com/spyhere/pokedexcli/internal/pokeClient"
)

func CleanInput(text string) []string {
	return strings.Split(
		strings.ToLower(strings.TrimSpace(text)),
		" ",
	)
}

func GetPokemonString(pokemon pokeClient.PokemonResponse) string {
	pokemonStats := ""
	for _, it := range pokemon.Stats {
		pokemonStats += fmt.Sprintf(" -%s: %v\n", it.Stat.Name, it.BaseStat)
	}
	pokemonTypes := ""
	for _, it := range pokemon.Types {
		pokemonTypes += fmt.Sprintf(" -%s\n", it.Type.Name)
	}
	output := fmt.Sprintf(`Name: %s
Height: %d
Weight: %d
Stats:
%sTypes:
%s`,
		pokemon.Name,
		pokemon.Height,
		pokemon.Weight,
		pokemonStats,
		pokemonTypes,
	)
	return output
}
