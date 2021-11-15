package app

import (
	"fmt"

	"github.com/DannyLyubenov/pokemon_rest_api/pkg/apis/pokemon"
	"github.com/DannyLyubenov/pokemon_rest_api/pkg/apis/shakespeare"
	"github.com/DannyLyubenov/pokemon_rest_api/pkg/apis/yoda"
	"github.com/DannyLyubenov/pokemon_rest_api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Habitat     string `json:"habitat"`
	IsLegendary bool   `json:"isLegendary"`
	APILimit    bool   `json:"apiLimit"`
}

func GetFunDescription(c *gin.Context) {
	pokemonName := c.Param("name")
	stats, err := pokemon.GetStats(pokemonName)
	if err != nil {
		fmt.Println(err)
	}

	desc, apiLimit, err := DecideTranslation(stats)
	if err != nil {
		fmt.Println(err)
	}

	uResp := &userResponse{
		Name:        stats.Name,
		Desc:        desc,
		Habitat:     stats.Habitat.Name,
		IsLegendary: stats.IsLegendary,
		APILimit:    apiLimit,
	}

	c.JSON(200, uResp)
}

func GetBasicInfo(c *gin.Context) {
	pokemonName := c.Param("name")
	stats, err := pokemon.GetStats(pokemonName)
	if err != nil {
		fmt.Println(err)
	}

	uResp := &userResponse{
		Name:        stats.Name,
		Desc:        utils.TrimDesc(stats),
		Habitat:     stats.Habitat.Name,
		IsLegendary: stats.IsLegendary,
	}

	c.JSON(200, uResp)
}

func DecideTranslation(stats pokemon.Stats) (string, bool, error) {
	if stats.Habitat.Name == "cave" || stats.IsLegendary {
		yoda, err := yoda.GetTranslation(stats)
		if err != nil {
			// return the original description and set the bool to true if we cannot translate the description
			return utils.TrimDesc(stats), true, err
		}
		return yoda.Contents.Translated, false, nil
	} else {
		sh, err := shakespeare.GetTranslation(stats)
		if err != nil {
			return utils.TrimDesc(stats), true, err
		}
		return sh.Contents.Translated, false, nil
	}
}
