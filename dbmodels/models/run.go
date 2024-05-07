package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONB Interface for JSONB Field of yourTableName Table
type JSONB map[string]interface{}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

type Run struct {
	Id         string `gorm:"primaryKey"`
	GameID     string
	CategoryID string
	PlayerId   string
	Values     JSONB `gorm:"type:json"`
	//Variables  []Variable
	Level    *string
	VideoURI string
	Runtime  float64
}
