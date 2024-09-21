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

type Pokemon struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	BaseExperience    int                `json:"base_experience"`
	Height            int                `json:"height"`
	IsDefault         bool               `json:"is_default"`
	Order             int                `json:"order"`
	Weight            int                `json:"weight"`
	Abilities         []Ability          `json:"abilities"`
	Forms             []Resource         `json:"forms"`
	GameIndices       []GameIndex        `json:"game_indices"`
	HeldItems         []HeldItem         `json:"held_items"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves             []Move             `json:"moves"`
	Species           Resource           `json:"species"`
	Sprites           Sprites            `json:"sprites"`
	Cries             Cries              `json:"cries"`
	Stats             []Stat             `json:"stats"`
	Types             []Type             `json:"types"`
	PastTypes         []PastType         `json:"past_types"`
}

type Ability struct {
	IsHidden bool     `json:"is_hidden"`
	Slot     int      `json:"slot"`
	Ability  Resource `json:"ability"`
}

type GameIndex struct {
	GameIndex int      `json:"game_index"`
	Version   Resource `json:"version"`
}

type HeldItem struct {
	Item           Resource         `json:"item"`
	VersionDetails []VersionDetail  `json:"version_details"`
}

type VersionDetail struct {
	Rarity  int      `json:"rarity"`
	Version Resource `json:"version"`
}

type Move struct {
	Move                 Resource             `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int      `json:"level_learned_at"`
	VersionGroup    Resource `json:"version_group"`
	MoveLearnMethod Resource `json:"move_learn_method"`
}

type Sprites struct {
	BackDefault       *string `json:"back_default"`
	BackFemale        *string `json:"back_female"`
	BackShiny         *string `json:"back_shiny"`
	BackShinyFemale   *string `json:"back_shiny_female"`
	FrontDefault      *string `json:"front_default"`
	FrontFemale       *string `json:"front_female"`
	FrontShiny        *string `json:"front_shiny"`
	FrontShinyFemale  *string `json:"front_shiny_female"`
	Other             OtherSprites `json:"other"`
	Versions          Versions     `json:"versions"`
}

type OtherSprites struct {
	DreamWorld DreamWorld `json:"dream_world"`
	Home       Home       `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown   Showdown   `json:"showdown"`
}

type DreamWorld struct {
	FrontDefault *string `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

type Home struct {
	FrontDefault *string `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
	FrontShiny   *string `json:"front_shiny"`
}

type OfficialArtwork struct {
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type Showdown struct {
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type Versions struct {
	GenerationI   GenerationI   `json:"generation-i"`
	GenerationII  GenerationII  `json:"generation-ii"`
	GenerationIII GenerationIII `json:"generation-iii"`
	GenerationIV  GenerationIV  `json:"generation-iv"`
	GenerationV   GenerationV   `json:"generation-v"`
	GenerationVI  GenerationVI  `json:"generation-vi"`
	GenerationVII GenerationVII `json:"generation-vii"`
	GenerationVIII GenerationVIII `json:"generation-viii"`
}

type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  Yellow  `json:"yellow"`
}

type RedBlue struct {
	BackDefault  *string `json:"back_default"`
	BackGray     *string `json:"back_gray"`
	FrontDefault *string `json:"front_default"`
	FrontGray    *string `json:"front_gray"`
}

type Yellow struct {
	BackDefault  *string `json:"back_default"`
	BackGray     *string `json:"back_gray"`
	FrontDefault *string `json:"front_default"`
	FrontGray    *string `json:"front_gray"`
}

type GenerationII struct {
	Crystal Crystal `json:"crystal"`
	Gold    Gold    `json:"gold"`
	Silver  Silver  `json:"silver"`
}

type Crystal struct {
	BackDefault  *string `json:"back_default"`
	BackShiny    *string `json:"back_shiny"`
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type Gold struct {
	BackDefault  *string `json:"back_default"`
	BackShiny    *string `json:"back_shiny"`
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type Silver struct {
	BackDefault  *string `json:"back_default"`
	BackShiny    *string `json:"back_shiny"`
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type GenerationIII struct {
	Emerald            Emerald            `json:"emerald"`
	FireRedLeafGreen   FireRedLeafGreen   `json:"firered-leafgreen"`
	RubySapphire       RubySapphire       `json:"ruby-sapphire"`
}

type Emerald struct {
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type FireRedLeafGreen struct {
	BackDefault  *string `json:"back_default"`
	BackShiny    *string `json:"back_shiny"`
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type RubySapphire struct {
	BackDefault  *string `json:"back_default"`
	BackShiny    *string `json:"back_shiny"`
	FrontDefault *string `json:"front_default"`
	FrontShiny   *string `json:"front_shiny"`
}

type GenerationIV struct {
	DiamondPearl       DiamondPearl       `json:"diamond-pearl"`
	HeartGoldSoulSilver HeartGoldSoulSilver `json:"heartgold-soulsilver"`
	Platinum           Platinum           `json:"platinum"`
}

type DiamondPearl struct {
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type HeartGoldSoulSilver struct {
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type Platinum struct {
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type GenerationV struct {
	BlackWhite BlackWhite `json:"black-white"`
}

type BlackWhite struct {
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
}

type GenerationVI struct {
	OmegaRubyAlphaSapphire OmegaRubyAlphaSapphire `json:"omegaruby-alphasapphire"`
	XY                     XY                     `json:"x-y"`
}

type OmegaRubyAlphaSapphire struct {
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type XY struct {
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type GenerationVII struct {
	Icons             Icons             `json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}

type Icons struct {
	FrontDefault *string `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

type UltraSunUltraMoon struct {
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type GenerationVIII struct {
	Icons Icons `json:"icons"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Stat struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	Stat     Resource `json:"stat"`
}

type Type struct {
	Slot int      `json:"slot"`
	Type Resource `json:"type"`
}

type PastType struct {
	Generation Resource `json:"generation"`
	Types      []Type   `json:"types"`
}

type Resource struct{
	Name string `json:"name"`
	URL string	`json:"url"`
}