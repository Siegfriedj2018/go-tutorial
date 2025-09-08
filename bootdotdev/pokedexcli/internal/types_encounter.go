package internal

type EncounterMethod struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}
type EncounterMethodRates struct {
	EncounterMethod string `json:"encounter_method"`
	VersionDetails  string `json:"version_details"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Language string `json:"language"`
	Name     string `json:"name"`
}
type PokemonEncounters struct {
	Pokemon        string `json:"pokemon"`
	VersionDetails string `json:"version_details"`
}