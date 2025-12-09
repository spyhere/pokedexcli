package pokeClient

type LocationsPaginatedResult struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}

type LocationResult struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate,omitempty"`
			Version struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version"`
		} `json:"version_details,omitempty"`
	} `json:"encounter_method_rates,omitempty"`
	GameIndex int `json:"game_index,omitempty"`
	ID        int `json:"id,omitempty"`
	Location  struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"location"`
	Name  string `json:"name,omitempty"`
	Names []struct {
		Language struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"language"`
		Name string `json:"name,omitempty"`
	} `json:"names,omitempty"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance,omitempty"`
				ConditionValues []any `json:"condition_values,omitempty"`
				MaxLevel        int   `json:"max_level,omitempty"`
				Method          struct {
					Name string `json:"name,omitempty"`
					URL  string `json:"url,omitempty"`
				} `json:"method"`
				MinLevel int `json:"min_level,omitempty"`
			} `json:"encounter_details,omitempty"`
			MaxChance int `json:"max_chance,omitempty"`
			Version   struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version"`
		} `json:"version_details,omitempty"`
	} `json:"pokemon_encounters,omitempty"`
}

type PokemonResponse struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden,omitempty"`
		Slot     int  `json:"slot,omitempty"`
	} `json:"abilities,omitempty"`
	BaseExperience int `json:"base_experience,omitempty"`
	Cries          struct {
		Latest string `json:"latest,omitempty"`
		Legacy string `json:"legacy,omitempty"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"forms,omitempty"`
	GameIndices []struct {
		GameIndex int `json:"game_index,omitempty"`
		Version   struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"version"`
	} `json:"game_indices,omitempty"`
	Height                 int    `json:"height,omitempty"`
	HeldItems              []any  `json:"held_items,omitempty"`
	ID                     int    `json:"id,omitempty"`
	IsDefault              bool   `json:"is_default,omitempty"`
	LocationAreaEncounters string `json:"location_area_encounters,omitempty"`
	Moves                  []struct {
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
	} `json:"moves,omitempty"`
	Name          string `json:"name,omitempty"`
	Order         int    `json:"order,omitempty"`
	PastAbilities []struct {
		Abilities []struct {
			Ability  any  `json:"ability,omitempty"`
			IsHidden bool `json:"is_hidden,omitempty"`
			Slot     int  `json:"slot,omitempty"`
		} `json:"abilities,omitempty"`
		Generation struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"generation"`
	} `json:"past_abilities,omitempty"`
	PastTypes []any `json:"past_types,omitempty"`
	Species   struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"species"`
	Sprites struct {
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
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat,omitempty"`
		Effort   int `json:"effort,omitempty"`
		Stat     struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"stat"`
	} `json:"stats,omitempty"`
	Types []struct {
		Slot int `json:"slot,omitempty"`
		Type struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"type"`
	} `json:"types,omitempty"`
	Weight int `json:"weight,omitempty"`
}
