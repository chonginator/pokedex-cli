package pokeapi

type PaginatedResourceList struct{
	Count int 						`json:"count"`
	Next *string 					`json:"next"`
	Previous *string 			`json:"previous"`
	Results []struct{
		Name string
		URL string
	} `json:"results"`
}

type LocationArea struct{
	ID int
	Name string
	GameIndex int
	EncounterMethodRates []struct{
		EncounterMethod struct{
			Name string
			URL string
		}
		VersionDetails []struct{
			Rate int
			Version struct{
				Name string
				URL string
			}
		}
	}
	Location struct{
		Name string
		URL string
	}
	Names []struct{
		Name string
		Language struct{
			Name string
			URL string
		}
	}
	PokemonEncounters []struct{
		Pokemon struct{
			Name string
			URL string
		}
		VersionDetails []struct{
			Version struct{
				Name string
				URL string
			}
			MaxChance int
			EncounterDetails []struct{
				MinLevel int
				MaxLevel int
				ConditionValues interface{}
				Chance int
				Method struct{
					Name string
					URL string
				}
			}
		}
	} `json:"pokemon_encounters"`
}
