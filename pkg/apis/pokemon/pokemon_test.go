package pokemon

import "testing"

func TestGetStats(t *testing.T) {
	pokemonName := "ONIX" // testing for name sensitivity
	stats, err := GetStats(pokemonName)
	if err != nil {
		t.Fatal(err)
	}

	if stats.Name != "onix" || stats.Habitat.Name != "cave" {
		t.Fatal("unexpected output:", stats.Name)
	}
}
