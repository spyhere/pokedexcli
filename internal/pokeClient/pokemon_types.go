package pokeClient

type PokemonResponse struct {
	Abilities              []PokeAbilities     `json:"abilities,omitempty"`
	BaseExperience         int                 `json:"base_experience,omitempty"`
	Cries                  PokeCries           `json:"cries"`
	Forms                  []PokeForms         `json:"forms,omitempty"`
	GameIndices            []PokeGameIndices   `json:"game_indices,omitempty"`
	Height                 int                 `json:"height,omitempty"`
	HeldItems              []any               `json:"held_items,omitempty"`
	ID                     int                 `json:"id,omitempty"`
	IsDefault              bool                `json:"is_default,omitempty"`
	LocationAreaEncounters string              `json:"location_area_encounters,omitempty"`
	Moves                  []PokeMoves         `json:"moves,omitempty"`
	Name                   string              `json:"name,omitempty"`
	Order                  int                 `json:"order,omitempty"`
	PastAbilities          []PokePastAbilities `json:"past_abilities,omitempty"`
	PastTypes              []any               `json:"past_types,omitempty"`
	Species                PokeSpecies         `json:"species"`
	Sprites                PokeSprites         `json:"sprites"`
	Stats                  []PokeStats         `json:"stats,omitempty"`
	Types                  []PokeTypes         `json:"types,omitempty"`
	Weight                 int                 `json:"weight,omitempty"`
}

type PokeAbilities struct {
	Ability struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"ability"`
	IsHidden bool `json:"is_hidden,omitempty"`
	Slot     int  `json:"slot,omitempty"`
}

type PokeCries struct {
	Latest string `json:"latest,omitempty"`
	Legacy string `json:"legacy,omitempty"`
}

type PokeForms struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type PokeGameIndices struct {
	GameIndex int `json:"game_index,omitempty"`
	Version   struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"version"`
}

type PokeMoves struct {
	Move struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"move"`
	VersionGroupDetails []struct {
		LevelLearnedAt  int `json:"level_learned_at,omitempty"`
		MoveLearnMethod struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"move_learn_method"`
		Order        int `json:"order,omitempty"`
		VersionGroup struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"version_group"`
	} `json:"version_group_details,omitempty"`
}

type PokePastAbilities struct {
	Abilities []struct {
		Ability  any  `json:"ability,omitempty"`
		IsHidden bool `json:"is_hidden,omitempty"`
		Slot     int  `json:"slot,omitempty"`
	} `json:"abilities,omitempty"`
	Generation struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"generation"`
}

type PokeSpecies struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type PokeSprites struct {
	BackDefault      string `json:"back_default,omitempty"`
	BackFemale       any    `json:"back_female,omitempty"`
	BackShiny        string `json:"back_shiny,omitempty"`
	BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
	FrontDefault     string `json:"front_default,omitempty"`
	FrontFemale      any    `json:"front_female,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	Other            struct {
		DreamWorld struct {
			FrontDefault string `json:"front_default,omitempty"`
			FrontFemale  any    `json:"front_female,omitempty"`
		} `json:"dream_world"`
		Home struct {
			FrontDefault     string `json:"front_default,omitempty"`
			FrontFemale      any    `json:"front_female,omitempty"`
			FrontShiny       string `json:"front_shiny,omitempty"`
			FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
		} `json:"home"`
		OfficialArtwork struct {
			FrontDefault string `json:"front_default,omitempty"`
			FrontShiny   string `json:"front_shiny,omitempty"`
		} `json:"official-artwork"`
		Showdown struct {
			BackDefault      string `json:"back_default,omitempty"`
			BackFemale       any    `json:"back_female,omitempty"`
			BackShiny        string `json:"back_shiny,omitempty"`
			BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
			FrontDefault     string `json:"front_default,omitempty"`
			FrontFemale      any    `json:"front_female,omitempty"`
			FrontShiny       string `json:"front_shiny,omitempty"`
			FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
		} `json:"showdown"`
	} `json:"other"`
	Versions struct {
		GenerationI struct {
			RedBlue struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackGray         string `json:"back_gray,omitempty"`
				BackTransparent  string `json:"back_transparent,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontGray        string `json:"front_gray,omitempty"`
				FrontTransparent string `json:"front_transparent,omitempty"`
			} `json:"red-blue"`
			Yellow struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackGray         string `json:"back_gray,omitempty"`
				BackTransparent  string `json:"back_transparent,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontGray        string `json:"front_gray,omitempty"`
				FrontTransparent string `json:"front_transparent,omitempty"`
			} `json:"yellow"`
		} `json:"generation-i"`
		GenerationIi struct {
			Crystal struct {
				BackDefault           string `json:"back_default,omitempty"`
				BackShiny             string `json:"back_shiny,omitempty"`
				BackShinyTransparent  string `json:"back_shiny_transparent,omitempty"`
				BackTransparent       string `json:"back_transparent,omitempty"`
				FrontDefault          string `json:"front_default,omitempty"`
				FrontShiny            string `json:"front_shiny,omitempty"`
				FrontShinyTransparent string `json:"front_shiny_transparent,omitempty"`
				FrontTransparent      string `json:"front_transparent,omitempty"`
			} `json:"crystal"`
			Gold struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackShiny        string `json:"back_shiny,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontTransparent string `json:"front_transparent,omitempty"`
			} `json:"gold"`
			Silver struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackShiny        string `json:"back_shiny,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontTransparent string `json:"front_transparent,omitempty"`
			} `json:"silver"`
		} `json:"generation-ii"`
		GenerationIii struct {
			Emerald struct {
				FrontDefault string `json:"front_default,omitempty"`
				FrontShiny   string `json:"front_shiny,omitempty"`
			} `json:"emerald"`
			FireredLeafgreen struct {
				BackDefault  string `json:"back_default,omitempty"`
				BackShiny    string `json:"back_shiny,omitempty"`
				FrontDefault string `json:"front_default,omitempty"`
				FrontShiny   string `json:"front_shiny,omitempty"`
			} `json:"firered-leafgreen"`
			RubySapphire struct {
				BackDefault  string `json:"back_default,omitempty"`
				BackShiny    string `json:"back_shiny,omitempty"`
				FrontDefault string `json:"front_default,omitempty"`
				FrontShiny   string `json:"front_shiny,omitempty"`
			} `json:"ruby-sapphire"`
		} `json:"generation-iii"`
		GenerationIv struct {
			DiamondPearl struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackFemale       any    `json:"back_female,omitempty"`
				BackShiny        string `json:"back_shiny,omitempty"`
				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      any    `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
			} `json:"diamond-pearl"`
			HeartgoldSoulsilver struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackFemale       any    `json:"back_female,omitempty"`
				BackShiny        string `json:"back_shiny,omitempty"`
				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      any    `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
			} `json:"heartgold-soulsilver"`
			Platinum struct {
				BackDefault      string `json:"back_default,omitempty"`
				BackFemale       any    `json:"back_female,omitempty"`
				BackShiny        string `json:"back_shiny,omitempty"`
				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      any    `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
			} `json:"platinum"`
		} `json:"generation-iv"`
		GenerationV struct {
			BlackWhite struct {
				Animated struct {
					BackDefault      string `json:"back_default,omitempty"`
					BackFemale       any    `json:"back_female,omitempty"`
					BackShiny        string `json:"back_shiny,omitempty"`
					BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
					FrontDefault     string `json:"front_default,omitempty"`
					FrontFemale      any    `json:"front_female,omitempty"`
					FrontShiny       string `json:"front_shiny,omitempty"`
					FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
				} `json:"animated"`
				BackDefault      string `json:"back_default,omitempty"`
				BackFemale       any    `json:"back_female,omitempty"`
				BackShiny        string `json:"back_shiny,omitempty"`
				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      any    `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
			} `json:"black-white"`
		} `json:"generation-v"`
		GenerationVi struct {
			OmegarubyAlphasapphire struct {
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      any    `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
			} `json:"omegaruby-alphasapphire"`
			XY struct {
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      any    `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
			} `json:"x-y"`
		} `json:"generation-vi"`
		GenerationVii struct {
			Icons struct {
				FrontDefault string `json:"front_default,omitempty"`
				FrontFemale  any    `json:"front_female,omitempty"`
			} `json:"icons"`
			UltraSunUltraMoon struct {
				FrontDefault     string `json:"front_default,omitempty"`
				FrontFemale      any    `json:"front_female,omitempty"`
				FrontShiny       string `json:"front_shiny,omitempty"`
				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
			} `json:"ultra-sun-ultra-moon"`
		} `json:"generation-vii"`
		GenerationViii struct {
			Icons struct {
				FrontDefault string `json:"front_default,omitempty"`
				FrontFemale  any    `json:"front_female,omitempty"`
			} `json:"icons"`
		} `json:"generation-viii"`
	} `json:"versions"`
}
type PokeStats struct {
	BaseStat int      `json:"base_stat,omitempty"`
	Effort   int      `json:"effort,omitempty"`
	Stat     PokeStat `json:"stat"`
}
type PokeStat struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type PokeTypes struct {
	Slot int      `json:"slot,omitempty"`
	Type PokeType `json:"type"`
}
type PokeType struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
