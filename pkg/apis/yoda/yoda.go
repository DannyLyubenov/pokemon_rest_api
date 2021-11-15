package yoda

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/DannyLyubenov/pokemon_rest_api/pkg/apis/pokemon"
	"github.com/DannyLyubenov/pokemon_rest_api/pkg/httprequest"
	"github.com/DannyLyubenov/pokemon_rest_api/pkg/utils"
)

type Yoda struct {
	Success  Success  `json:"success"`
	Contents Contents `json:"contents"`
}

type Success struct {
	Total int `json:"total"`
}

type Contents struct {
	Translated string `json:"translated"`
}

func GetTranslation(stats pokemon.Stats) (Yoda, error) {
	desc := utils.TrimDesc(stats)
	endpoint := fmt.Sprintf("https://api.funtranslations.com/translate/yoda.json?text=%s", url.QueryEscape(desc))
	body, err := httprequest.GetBody(endpoint)
	if err != nil {
		return Yoda{}, err
	}

	var yoda Yoda
	if err := json.Unmarshal(body, &yoda); err != nil {
		return Yoda{}, err
	}
	return yoda, nil
}
