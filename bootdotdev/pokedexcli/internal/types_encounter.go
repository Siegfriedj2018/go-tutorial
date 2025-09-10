package internal

type EncounterMethods struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethodRates struct {
	EncounterMethods EncounterMethods 	`json:"encounter_method"`
	EncounterVersion []EncounterVersion `json:"version_details"`
}

type EncounterMethod struct {
	Name string	 `json:"name"`
	Url  *string `json:"url"`
}

type EncounterVersion struct {
	Rate		int 		`json:"rate"`
	Version Version `json:"version"`
}

type Version struct {
	Name string 	`json:"name"`
	Url  *string 	`json:"url"`
}

type Location struct {
	Name string  `json:"name"`
	Url  *string `json:"url"`
}

type Names struct {
	Name    	*string `json:"name"`
	Language	
}
