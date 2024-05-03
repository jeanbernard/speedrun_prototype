package models

type Category struct {
	Id     string `gorm:"primaryKey"`
	GameID string
	Name   string
	Type   string
}
