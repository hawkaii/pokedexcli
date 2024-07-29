package pokeapi

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ability"`
		Is_hidden bool `json:"is_hidden"`
		Slot      int  `json:"slot"`
	} `json:"abilities"`
	Base_experience          int    `json:"base_experience"`
	Height                   int    `json:"height"`
	Id                       int    `json:"id"`
	Is_default               bool   `json:"is_default"`
	Location_area_encounters string `json:"location_area_encounters"`
	Name                     string `json:"name"`
	Order                    int    `json:"order"`
	Weight                   int    `json:"weight"`

	Stats []struct {
		Base_stat int `json:"base_stat"`
		Effort    int `json:"effort"`
		Stat      struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
