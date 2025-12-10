package main_test

import (
	"fmt"
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

func TestRunCommand(t *testing.T) {
	commands := map[string]main.CliCommand{
		"help": {
			Name:        "help",
			Description: "help cmd",
			Cb: func(c *main.Config, s ...string) error {
				return fmt.Errorf("help:%s", strings.Join(s, ","))
			},
		},
		"test": {
			Name:        "test",
			Description: "test cmd",
			Cb: func(c *main.Config, s ...string) error {
				return fmt.Errorf("test:%s", strings.Join(s, ","))
			},
		},
	}
	config := &main.Config{}
	tests := []struct {
		name     string
		c        *main.Config
		words    []string
		commands map[string]main.CliCommand
		wantErr  string
	}{
		{
			name:     "Should call test command with given arguments",
			c:        config,
			words:    []string{"test", "arg1", "arg2"},
			commands: commands,
			wantErr:  "test:arg1,arg2",
		},
		{
			name:     "Should call 'help' with given arguments if command not found",
			c:        config,
			words:    []string{"kill", "arg1"},
			commands: commands,
			wantErr:  "help:arg1",
		},
		{
			name:     "Should call 'help' without arguments",
			c:        config,
			words:    []string{"help"},
			commands: commands,
			wantErr:  "help:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := main.RunCommand(tt.c, tt.words, tt.commands)
			if gotErr.Error() != tt.wantErr {
				t.Fatalf("\n--Got--\n%s\n--Want--\n%s\n", gotErr.Error(), tt.wantErr)
			}
		})
	}
}

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "SoMe CrazY Text ",
			expected: []string{"some", "crazy", "text"},
		},
	}
	for _, c := range testCases {
		actual := main.CleanInput(c.input)
		for i := range c.expected {
			if actual[i] != c.expected[i] {
				t.Errorf("not equal: %v, %v", actual[i], c.expected[i])
				t.Fail()
			}
		}
	}
}
