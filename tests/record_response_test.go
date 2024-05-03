package tests

import (
	"developer/any/models"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getJSONExample() ([]byte, error) {
	jsonFile, err := os.ReadFile("../api/examples/records_example.json")
	if err != nil {
		return []byte{}, err
	}
	return jsonFile, nil
}

func TestJSONParse(t *testing.T) {
	jsonStr, err := getJSONExample()
	if err != nil {
		panic(err)
	}

	spr := models.RecordsResponse{}

	err = json.Unmarshal(jsonStr, &spr)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(spr.Data))

	// Player's id/name in outer data struct
	player := spr.Data[0].Players
	assert.Equal(t, "48gn05yj", player.Id)
	assert.Equal(t, "PyramidK", player.Names.International)
	assert.Equal(t, "", player.Names.Japanese)

	// Category in outer data struct
	category := spr.Data[0].Category.Data
	assert.Equal(t, "PC", category.Name)
	assert.Equal(t, "82466qnd", category.Id)
	assert.Equal(t, "per-game", category.Type)

	// Variables in outer data struct
	assert.Equal(t, 3, len(spr.Data[0].Variables.Data))

	variableNames := []string{"Scenario", "Category", "Difficulty"}
	variableIds := []string{"kn033dn3", "789qw3lw", "jlz197n2"}

	for idx, varValue := range spr.Data[0].Variables.Data {
		assert.Equal(t, variableNames[idx], varValue.Name)
		assert.Equal(t, variableIds[idx], varValue.Id)

		for key, value := range varValue.Values.Values {
			switch key {
			case "5le3736q":
				assert.Equal(t, "Leon A", value.Label)
			case "0q5xjxv1":
				assert.Equal(t, "Claire A", value.Label)
			case "4lx070gl":
				assert.Equal(t, "Leon B", value.Label)
			case "814k0kkl":
				assert.Equal(t, "Claire B", value.Label)
			case "21dk9dpl":
				assert.Equal(t, "Any%", value.Label)
			case "5q8kjzyq":
				assert.Equal(t, "Low%", value.Label)
			case "gq750dy1":
				assert.Equal(t, "Normal", value.Label)
			case "21gnvdml":
				assert.Equal(t, "Hard", value.Label)
			default:
				t.Errorf("Unknown value ID")
			}
		}
	}

	// Speedruns
	assert.Equal(t, 1, len(spr.Data[0].Runs))
	speedrun := spr.Data[0].Runs[0]

	assert.Equal(t, "me3ll79z", speedrun.Run.Id)
	assert.Equal(t, float64(2915), speedrun.Run.Times.PrimaryTime)
	assert.Equal(t, "", speedrun.Run.Level)
	assert.Equal(t, "https://www.youtube.com/watch?v=I4gKsQM34mM", speedrun.Run.Videos.Links[0].URI)

	for key, value := range speedrun.Run.Values {
		switch key {
		case "kn033dn3":
			assert.Equal(t, "5le3736q", value)
		case "789qw3lw":
			assert.Equal(t, "21dk9dpl", value)
		case "jlz197n2":
			assert.Equal(t, "gq750dy1", value)
		}
	}

	assert.Equal(t, 1, len(speedrun.Run.Players))
	assert.Equal(t, "48gn05yj", speedrun.Run.Players[0].Id)
}
