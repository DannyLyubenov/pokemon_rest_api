package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/DannyLyubenov/pokemon_rest_api/pkg/apis/pokemon"
)

func TestDecideTranslation(t *testing.T) {
	mewtwo, abra, onix, horsea := pokemonMockData()

	// apiLimit is true when we exceed more than 5 API calls to funtranslations.com
	// because this is out of out control, the tests will pass
	mewtwoDesc, mewtwoApiLimit, _ := DecideTranslation(mewtwo)
	mewtwoExpected := "Created by a scientist after years of horrific gene splicing and dna engineering experiments,  it was."
	if mewtwoDesc != mewtwoExpected && mewtwoApiLimit != true {
		t.Fatal("unexpected output:", mewtwo.Name)
	}

	abraDesc, abraApiLimit, _ := DecideTranslation(abra)
	abraExpected := "Using its ability to did read minds,  'twill unfold impending danger and teleport to safety."

	if abraDesc != abraExpected && abraApiLimit != true {
		t.Fatal("unexpected output:", abra.Name)
	}

	onixDesc, onixApiLimit, _ := DecideTranslation(onix)
	onixExpected := "It rapidly bores through the ground at 50 mph by squirming and twisting its massive,Rugged body."
	if onixDesc != onixExpected && onixApiLimit != true {
		t.Fatal("unexpected output:", onix.Name)
	}

	horseaDesc, _, _ := DecideTranslation(horsea)
	horseaExpected := "水面から　勢いよく　墨を　発射して 飛んでいる　虫を 撃ち落とすことがあるという。"
	if horseaDesc != horseaExpected {
		t.Fatal("unexpected output:", horsea.Name)
	}

}

func TestGetBasicInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/v2/pokemon/:name", GetBasicInfo)

	req, err := http.NewRequest(http.MethodGet, "/api/v2/pokemon/ekans", nil)
	if err != nil {
		t.Fatal("Couldn't create request:", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatal("unexpected response")
	}
}

func TestGetFunDescription(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/v2/pokemon/translated/:name", GetBasicInfo)

	req, err := http.NewRequest(http.MethodGet, "/api/v2/pokemon/translated/abra", nil)
	if err != nil {
		t.Fatal("Couldn't create request:", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatal("unexpected response")
	}
}

func pokemonMockData() (pokemon.Stats, pokemon.Stats, pokemon.Stats, pokemon.Stats) {
	mewtwo := pokemon.Stats{
		Name:        "mewtwo",
		IsLegendary: true,
		Desc: []pokemon.Desc{
			{FlavorText: "It was created by\na scientist after\nyears of horrific\fgene splicing and\nDNA engineering\nexperiments.",
				Language: pokemon.Language{
					Name: "en",
				}},
		},
		Habitat: pokemon.Habitat{
			Name: "rare",
		},
	}

	abra := pokemon.Stats{
		Name:        "abra",
		IsLegendary: false,
		Desc: []pokemon.Desc{
			{FlavorText: "Using its ability\nto read minds, it\nwill identify\fimpending danger\nand TELEPORT to\nsafety.",
				Language: pokemon.Language{
					Name: "en",
				}},
		},
		Habitat: pokemon.Habitat{
			Name: "urban",
		},
	}

	onix := pokemon.Stats{
		Name:        "onix",
		IsLegendary: false,
		Desc: []pokemon.Desc{
			{FlavorText: "It was created by\na scientist after\nyears of horrific\fgene splicing and\nDNA engineering\nexperiments.",
				Language: pokemon.Language{
					Name: "en",
				}},
		},
		Habitat: pokemon.Habitat{
			Name: "cave",
		},
	}

	horsea := pokemon.Stats{
		Name:        "horsea",
		IsLegendary: false,
		Desc: []pokemon.Desc{
			{FlavorText: "水面から　勢いよく　墨を　発射して\n飛んでいる　虫を\n撃ち落とすことがあるという。",
				Language: pokemon.Language{
					Name: "ja",
				}},
		},
		Habitat: pokemon.Habitat{
			Name: "sea",
		},
	}

	return mewtwo, abra, onix, horsea
}
