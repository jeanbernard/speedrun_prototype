package tests

import (
	"developer/any/models"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getGamesBulkJSONExample() ([]byte, error) {
	jsonFile, err := os.ReadFile("../api/examples/games_bulk_example.json")
	if err != nil {
		return []byte{}, err
	}
	return jsonFile, nil
}

func TestJSONParseGamesBulk(t *testing.T) {
	jsonStr, err := getGamesBulkJSONExample()
	if err != nil {
		panic(err)
	}

	bulk := models.GamesBulkResponse{}

	err = json.Unmarshal(jsonStr, &bulk)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(bulk.Data))

	game := bulk.Data[0]
	assert.Equal(t, "y65rm041", game.Id)
	assert.Equal(t, " BIRDIE WING -Golf Girls' Story-", game.Names["international"])

	pagination := bulk.Pagination
	assert.Equal(t, 2, pagination.Max)
	assert.Equal(t, 2, pagination.Size)
	assert.Equal(t, 2, pagination.Offset)

	assert.Equal(t, 2, len(pagination.Links))

	prevLink := pagination.Links[0]
	assert.Equal(t, "prev", prevLink.Rel)
	assert.Equal(t, "https://www.speedrun.com/api/v1/games?_bulk=yes&max=2", prevLink.URI)

	nextLink := pagination.Links[1]
	assert.Equal(t, "next", nextLink.Rel)
	assert.Equal(t, "https://www.speedrun.com/api/v1/games?_bulk=yes&max=2&offset=4", nextLink.URI)
}
