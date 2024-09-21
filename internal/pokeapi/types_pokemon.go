package pokeapi

type NamedResource struct {
  Name string `json:"name"`
  URL  string `json:"url"`
}

type Pokemon struct {
  Id                int    `json:"id"`
  Name              string `json:"name"`
  BaseExperience    int    `json:"base_experience"`
  Height            int    `json:"height"`
  IsDefault         bool   `json:"is_default"`
  Order             int    `json:"order"`
  Weight            int    `json:"weight"`
  Abilities         []struct {
    IsHidden bool         `json:"is_hidden"`
    Slot     int          `json:"slot"`
    Ability  NamedResource `json:"ability"`
  } `json:"abilities"`
  Forms []NamedResource `json:"forms"`
  GameIndices []struct {
    GameIndex int          `json:"game_index"`
    Version   NamedResource `json:"version"`
  } `json:"game_indices"`
  HeldItems []struct {
    Item           NamedResource `json:"item"`
    VersionDetails []struct {
      Rarity  int          `json:"rarity"`
      Version NamedResource `json:"version"`
    } `json:"version_details"`
  } `json:"held_items"`
  LocationAreaEncounters string `json:"location_area_encounters"`
  Moves []struct {
    Move                 NamedResource `json:"move"`
    VersionGroupDetails []struct {
      LevelLearnedAt  int          `json:"level_learned_at"`
      VersionGroup    NamedResource `json:"version_group"`
      MoveLearnMethod NamedResource `json:"move_learn_method"`
    } `json:"version_group_details"`
  } `json:"moves"`
  Species NamedResource `json:"species"`
  Sprites struct {
    BackDefault       *string `json:"back_default"`
    BackFemale        *string `json:"back_female"`
    BackShiny         *string `json:"back_shiny"`
    BackShinyFemale   *string `json:"back_shiny_female"`
    FrontDefault      *string `json:"front_default"`
    FrontFemale       *string `json:"front_female"`
    FrontShiny        *string `json:"front_shiny"`
    FrontShinyFemale  *string `json:"front_shiny_female"`
    Other             struct {
      DreamWorld struct {
        FrontDefault *string `json:"front_default"`
        FrontFemale  *string `json:"front_female"`
      } `json:"dream_world"`
      Home struct {
        FrontDefault *string `json:"front_default"`
        FrontFemale  *string `json:"front_female"`
        FrontShiny   *string `json:"front_shiny"`
      } `json:"home"`
      OfficialArtwork struct {
        FrontDefault *string `json:"front_default"`
        FrontShiny   *string `json:"front_shiny"`
      } `json:"official-artwork"`
      Showdown struct {
        BackDefault      *string `json:"back_default"`
        BackFemale       *string `json:"back_female"`
        BackShiny        *string `json:"back_shiny"`
        BackShinyFemale  *string `json:"back_shiny_female"`
        FrontDefault     *string `json:"front_default"`
        FrontFemale      *string `json:"front_female"`
        FrontShiny       *string `json:"front_shiny"`
        FrontShinyFemale *string `json:"front_shiny_female"`
      } `json:"showdown"`
    } `json:"other"`
    Versions struct {
      GenerationI struct {
        RedBlue struct {
          BackDefault  *string `json:"back_default"`
          BackGray     *string `json:"back_gray"`
          FrontDefault *string `json:"front_default"`
          FrontGray    *string `json:"front_gray"`
        } `json:"red-blue"`
        Yellow struct {
          BackDefault  *string `json:"back_default"`
          BackGray     *string `json:"back_gray"`
          FrontDefault *string `json:"front_default"`
          FrontGray    *string `json:"front_gray"`
        } `json:"yellow"`
      } `json:"generation-i"`
      GenerationII struct {
        Crystal struct {
          BackDefault  *string `json:"back_default"`
          BackShiny    *string `json:"back_shiny"`
          FrontDefault *string `json:"front_default"`
          FrontShiny   *string `json:"front_shiny"`
        } `json:"crystal"`
        Gold struct {
          BackDefault  *string `json:"back_default"`
          BackShiny    *string `json:"back_shiny"`
          FrontDefault *string `json:"front_default"`
          FrontShiny   *string `json:"front_shiny"`
        } `json:"gold"`
        Silver struct {
          BackDefault  *string `json:"back_default"`
          BackShiny    *string `json:"back_shiny"`
          FrontDefault *string `json:"front_default"`
          FrontShiny   *string `json:"front_shiny"`
        } `json:"silver"`
      } `json:"generation-ii"`
      GenerationIII struct {
        Emerald struct {
          FrontDefault *string `json:"front_default"`
          FrontShiny   *string `json:"front_shiny"`
        } `json:"emerald"`
        FireRedLeafGreen struct {
          BackDefault  *string `json:"back_default"`
          BackShiny    *string `json:"back_shiny"`
          FrontDefault *string `json:"front_default"`
          FrontShiny   *string `json:"front_shiny"`
        } `json:"firered-leafgreen"`
        RubySapphire struct {
          BackDefault  *string `json:"back_default"`
          BackShiny    *string `json:"back_shiny"`
          FrontDefault *string `json:"front_default"`
          FrontShiny   *string `json:"front_shiny"`
        } `json:"ruby-sapphire"`
      } `json:"generation-iii"`
      GenerationIV struct {
        DiamondPearl struct {
          BackDefault      *string `json:"back_default"`
          BackFemale       *string `json:"back_female"`
          BackShiny        *string `json:"back_shiny"`
          BackShinyFemale  *string `json:"back_shiny_female"`
          FrontDefault     *string `json:"front_default"`
          FrontFemale      *string `json:"front_female"`
          FrontShiny       *string `json:"front_shiny"`
          FrontShinyFemale *string `json:"front_shiny_female"`
        } `json:"diamond-pearl"`
        HeartGoldSoulSilver struct {
          BackDefault      *string `json:"back_default"`
          BackFemale       *string `json:"back_female"`
          BackShiny        *string `json:"back_shiny"`
          BackShinyFemale  *string `json:"back_shiny_female"`
          FrontDefault     *string `json:"front_default"`
          FrontFemale      *string `json:"front_female"`
          FrontShiny       *string `json:"front_shiny"`
          FrontShinyFemale *string `json:"front_shiny_female"`
        } `json:"heartgold-soulsilver"`
        Platinum struct {
          BackDefault      *string `json:"back_default"`
          BackFemale       *string `json:"back_female"`
          BackShiny        *string `json:"back_shiny"`
          BackShinyFemale  *string `json:"back_shiny_female"`
          FrontDefault     *string `json:"front_default"`
          FrontFemale      *string `json:"front_female"`
          FrontShiny       *string `json:"front_shiny"`
          FrontShinyFemale *string `json:"front_shiny_female"`
        } `json:"platinum"`
      } `json:"generation-iv"`
      GenerationV struct {
        BlackWhite struct {
          Animated struct {
            BackDefault      *string `json:"back_default"`
            BackFemale       *string `json:"back_female"`
            BackShiny        *string `json:"back_shiny"`
            BackShinyFemale  *string `json:"back_shiny_female"`
            FrontDefault     *string `json:"front_default"`
            FrontFemale      *string `json:"front_female"`
            FrontShiny       *string `json:"front_shiny"`
            FrontShinyFemale *string `json:"front_shiny_female"`
          } `json:"animated"`
          BackDefault      *string `json:"back_default"`
          BackFemale       *string `json:"back_female"`
          BackShiny        *string `json:"back_shiny"`
          BackShinyFemale  *string `json:"back_shiny_female"`
          FrontDefault     *string `json:"front_default"`
          FrontFemale      *string `json:"front_female"`
          FrontShiny       *string `json:"front_shiny"`
          FrontShinyFemale *string `json:"front_shiny_female"`
        } `json:"black-white"`
      } `json:"generation-v"`
      GenerationVI struct {
        OmegaRubyAlphaSapphire struct {
          FrontDefault     *string `json:"front_default"`
          FrontFemale      *string `json:"front_female"`
          FrontShiny       *string `json:"front_shiny"`
          FrontShinyFemale *string `json:"front_shiny_female"`
        } `json:"omegaruby-alphasapphire"`
        XY struct {
          FrontDefault     *string `json:"front_default"`
          FrontFemale      *string `json:"front_female"`
          FrontShiny       *string `json:"front_shiny"`
          FrontShinyFemale *string `json:"front_shiny_female"`
        } `json:"x-y"`
      } `json:"generation-vi"`
      GenerationVII struct {
        Icons struct {
          FrontDefault *string `json:"front_default"`
          FrontFemale  *string `json:"front_female"`
        } `json:"icons"`
        UltraSunUltraMoon struct {
          FrontDefault     *string `json:"front_default"`
          FrontFemale      *string `json:"front_female"`
          FrontShiny       *string `json:"front_shiny"`
          FrontShinyFemale *string `json:"front_shiny_female"`
        } `json:"ultra-sun-ultra-moon"`
      } `json:"generation-vii"`
      GenerationVIII struct {
        Icons struct {
          FrontDefault *string `json:"front_default"`
          FrontFemale  *string `json:"front_female"`
        } `json:"icons"`
      } `json:"generation-viii"`
    } `json:"versions"`
  } `json:"sprites"`
  Cries struct {
    Latest string `json:"latest"`
    Legacy string `json:"legacy"`
  } `json:"cries"`
  Stats []struct {
		BaseStat int      `json:"base_stat"`
		Effort   int      `json:"effort"`
		Stat     struct{
			Name string
			URL string
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int      `json:"slot"`
		Type struct{
			Name string
			URL string
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct{
			Name string
			URL string
		} `json:"generation"`
		Types      []struct {
			Slot int      `json:"slot"`
			Type struct{
				Name string
				URL string
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
}