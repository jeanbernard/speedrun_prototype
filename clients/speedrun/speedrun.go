package speedrun

import (
	"developer/any/models"
	"encoding/json"
	"fmt"
	"io"

	dbmodels "developer/any/dbmodels/models"
	"net/http"
)

const GAME_ID = "m1zjmz60"

func GetGames() (models.GamesResponse, error) {
	var game models.GamesResponse

	resp, err := http.Get(fmt.Sprintf("https://www.speedrun.com/api/v1/games/%v?embed=variables,categories", GAME_ID))
	if err != nil {
		return game, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return game, err
	}

	if err := json.Unmarshal(body, &game); err != nil {
		return game, err
	}

	return game, nil
}

func GetGamesBulk() (models.GamesBulkResponse, error) {
	var game models.GamesBulkResponse

	resp, err := http.Get("https://www.speedrun.com/api/v1/games?_bulk=yes&max=20")
	if err != nil {
		return game, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return game, err
	}

	if err := json.Unmarshal(body, &game); err != nil {
		return game, err
	}

	return game, nil
}

func GetRecords(game dbmodels.Game) (models.RecordsResponse, error) {
	var spr models.RecordsResponse

	resp, err := http.Get(
		fmt.Sprintf(
			"https://www.speedrun.com/api/v1/games/%v/records?embed=variables,category&top=1&skip-empty=true&max=10",
			game.Id))

	if err != nil {
		return spr, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return spr, err
	}

	// Convert response to objects
	if err := json.Unmarshal(body, &spr); err != nil {
		return spr, err
	}

	return spr, err
}
