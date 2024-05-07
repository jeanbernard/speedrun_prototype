package models

import "developer/any/dbmodels/models"

type PlayerNames struct {
	International string `json:"international,omitempty"`
	Japanese      string `json:"japanese,omitempty"`
}

type Player struct {
	Id    string
	Names PlayerNames `json:"names,omitempty"`
}

type CategoryData struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type Category struct {
	Data CategoryData `json:"data,omitempty"`
}

type ValuesContainer struct {
	Values map[string]Values `json:"values,omitempty"`
}

type Values struct {
	Label string `json:"label,omitempty"`
	//Rules string `json:"rules,omitempty"`
}

type VariableData struct {
	Id     string          `json:"id,omitempty"`
	Name   string          `json:"name,omitempty"`
	Values ValuesContainer `json:"values,omitempty"`
}

type Variable struct {
	Data []VariableData `json:"data,omitempty"`
}

type PlayersRun struct {
	Id string `json:"id,omitempty"`
}

type Link struct {
	URI string `json:"uri,omitempty"`
}

type Video struct {
	Links []Link `json:"links,omitempty"`
}

type RunTime struct {
	PrimaryTime float64 `json:"primary_t,omitempty"`
}

type Run struct {
	Id      string       `json:"id,omitempty"`
	Level   string       `json:"level,omitempty"`
	Values  models.JSONB `json:"values,omitempty"`
	Videos  Video        `json:"videos,omitempty"`
	Times   RunTime      `json:"times,omitempty"`
	Players []PlayersRun `json:"players,omitempty"`
}

type Runs struct {
	Run Run `json:"run,omitempty"`
}

type Data struct {
	Category  Category `json:"category,omitempty"`
	Variables Variable `json:"variables,omitempty"`
	Players   Player   `json:"players,omitempty"`
	Runs      []Runs   `json:"runs,omitempty"`
}

type RecordsResponse struct {
	Data []Data `json:"data,omitempty"`
}
