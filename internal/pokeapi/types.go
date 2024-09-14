package pokeapi

type PaginatedResourceList struct{
	Count int 						`json:"count"`
	Next *string 					`json:"next"`
	Previous *string 			`json:"previous"`
	Results []Resource `json:"results"`
}

type Resource struct{
	Name string `json:"name"`
	URL string	`json:"url"`
}