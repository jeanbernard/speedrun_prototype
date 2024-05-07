package models

type GamesData struct {
	Id    string            `json:"id,omitempty"`
	Names map[string]string `json:"names,omitempty"`
}

type GamesResponse struct {
	GamesData GamesData `json:"data,omitempty"`
}

type PaginationLinks struct {
	Rel string `json:"rel,omitempty"`
	URI string `json:"uri,omitempty"`
}

type Pagination struct {
	Offset int               `json:"offset,omitempty"`
	Max    int               `json:"max,omitempty"`
	Size   int               `json:"size,omitempty"`
	Links  []PaginationLinks `json:"links,omitempty"`
}

type GamesBulkResponse struct {
	Data       []GamesData
	Pagination Pagination
}
