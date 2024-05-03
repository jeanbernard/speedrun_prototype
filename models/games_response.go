package models

type GamesData struct {
	Id    string            `json:"id,omitempty"`
	Names map[string]string `json:"names,omitempty"`
	//Categories Category          `json:"categories,omitempty"`
	Variables *Variable `json:"variables,omitempty"`
}

type GamesResponse struct {
	GamesData GamesData `json:"data,omitempty"`
}

type GamesBulkResponse struct {
	Data []GamesData
}
