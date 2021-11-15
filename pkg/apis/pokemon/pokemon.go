package pokemon

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DannyLyubenov/pokemon_rest_api/pkg/httprequest"
)

type species struct {
	Species speciesURL `json:"species"`
}

type speciesURL struct {
	URL string `json:"url"`
}

type Stats struct {
	Name        string  `json:"name"`
	IsLegendary bool    `json:"is_legendary"`
	Desc        []Desc  `json:"flavor_text_entries"`
	Habitat     Habitat `json:"habitat"`
}

type Habitat struct {
	Name string `json:"name"`
}

type Desc struct {
	FlavorText string   `json:"flavor_text"`
	Language   Language `json:"language"`
}

type Language struct {
	Name string `json:"name"`
}

func GetStats(pokemonName string) (Stats, error) {
	endpoint, err := getSpecies(strings.ToLower(pokemonName))
	if err != nil {
		return Stats{}, err
	}

	body, err := httprequest.GetBody(endpoint.Species.URL)
	if err != nil {
		return Stats{}, err
	}

	var s Stats
	if err := json.Unmarshal(body, &s); err != nil {
		return Stats{}, err
	}
	return s, nil
}

func getSpecies(pokemonName string) (species, error) {
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	body, err := httprequest.GetBody(endpoint)
	if err != nil {
		return species{}, err
	}

	var s species
	if err := json.Unmarshal(body, &s); err != nil {
		return species{}, err
	}

	return s, nil
}
