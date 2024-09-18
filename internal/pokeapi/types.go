package pokeapi

type PaginatedResourceList struct{
	Count int 						`json:"count"`
	Next *string 					`json:"next"`
	Previous *string 			`json:"previous"`
	Results []Resource `json:"results"`
}

type LocationArea struct{
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct{
	Pokemon Resource `json:"pokemon"`
}

type Resource struct{
	Name string `json:"name"`
	URL string	`json:"url"`
}