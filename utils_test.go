package main_test

import (
	"strings"
	"testing"

	"github.com/spyhere/pokedexcli"
	"github.com/spyhere/pokedexcli/internal/pokeClient"
)

func TestGetPokemonString(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		pokemon pokeClient.PokemonResponse
		want    string
	}{
		{
			name: "correctly makes the string from pokemon props",
			pokemon: pokeClient.PokemonResponse{
				Name:   "pikachu",
				Height: 15,
				Weight: 20,
				Stats: []struct {
					BaseStat int "json:\"base_stat,omitempty\""
					Effort   int "json:\"effort,omitempty\""
					Stat     struct {
						Name string "json:\"name,omitempty\""
						URL  string "json:\"url,omitempty\""
					} "json:\"stat\""
				}{
					{
						BaseStat: 40,
						Stat: struct {
							Name string "json:\"name,omitempty\""
							URL  string "json:\"url,omitempty\""
						}{
							Name: "attack",
						},
					},
				},
				Types: []struct {
					Slot int "json:\"slot,omitempty\""
					Type struct {
						Name string "json:\"name,omitempty\""
						URL  string "json:\"url,omitempty\""
					} "json:\"type\""
				}{
					{
						Type: struct {
							Name string "json:\"name,omitempty\""
							URL  string "json:\"url,omitempty\""
						}{
							Name: "electric",
						},
					},
				},
			},
			want: `Name: pikachu
Height: 15
Weight: 20
Stats:
 -attack: 40
Types:
 -electric
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.GetPokemonString(tt.pokemon)
			expected := strings.Split(tt.want, "\n")
			received := strings.Split(got, "\n")
			if len(expected) != len(received) {
				t.Fatalf("Received different length:\n--Want--\n%s\n--Got--\n%s\n", tt.want, got)
			}
			for idx, it := range received {
				if it != expected[idx] {
					t.Errorf("got %v, want %v", it, expected[idx])
				}
			}
		})
	}
}
