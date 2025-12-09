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

func PokemonInit(res *pokeClient.PokemonResponse) Pokemon {
	var baseHP int
	var baseAttack int
	var baseDefense int
	var specialAttack int
	var specialDefense int
	var speed int
	for _, it := range res.Stats {
		statName := it.Stat.Name
		baseStat := it.BaseStat
		switch statName {
		case "hp":
			baseHP = baseStat
		case "attack":
			baseAttack = baseStat
		case "defense":
			baseDefense = baseStat
		case "special-attack":
			specialAttack = baseStat
		case "special-defense":
			specialDefense = baseStat
		case "speed":
			speed = baseStat
		}
	}
	var types []string
	for _, it := range res.Types {
		types = append(types, it.Type.Name)
	}
	return Pokemon{
		Name:   res.Name,
		Height: res.Height,
		Weight: res.Weight,
		Stats: struct {
			HP             int
			Attack         int
			Defense        int
			SpecialAttack  int
			SpecialDefense int
			Speed          int
		}{
			HP:             baseHP,
			Attack:         baseAttack,
			Defense:        baseDefense,
			SpecialAttack:  specialAttack,
			SpecialDefense: specialDefense,
			Speed:          speed,
		},
		Types: types,
	}
}

func GetPokemonString(pokemon Pokemon) string {
	pokemonTypes := ""
	for _, it := range pokemon.Types {
		pokemonTypes += " -" + it
	}
	output := fmt.Sprintf(`Name: %s
Height: %d
Weight: %d
Stats:
 -hp: %d
 -attack: %d
 -defense: %d
 -special-attack: %d
 -special-defense: %d
 -speed: %d
Types:
%s
`,
		pokemon.Name,
		pokemon.Height,
		pokemon.Weight,
		pokemon.Stats.HP,
		pokemon.Stats.Attack,
		pokemon.Stats.Defense,
		pokemon.Stats.SpecialAttack,
		pokemon.Stats.SpecialDefense,
		pokemon.Stats.Speed,
		pokemonTypes,
	)
	return output
}
