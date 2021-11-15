package utils

import (
	"strings"

	"github.com/DannyLyubenov/pokemon_rest_api/pkg/apis/pokemon"
)

func TrimDesc(stats pokemon.Stats) string {
	desc := getDescription(stats)
	replacer := strings.NewReplacer("\n", " ", "\f", " ")
	return replacer.Replace(desc)
}

func getDescription(stats pokemon.Stats) string {
	var desc string
	for _, item := range stats.Desc {
		if item.Language.Name == "en" {
			desc = item.FlavorText
		}
	}

	if desc != "" {
		return desc
	} else if len(stats.Desc) > 0 {
		return stats.Desc[0].FlavorText
	} else {
		return ""
	}
}
