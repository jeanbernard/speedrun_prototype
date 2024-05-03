package models

type Game struct {
	Id       string `gorm:"primaryKey"`
	Runs     []*Run
	Category []Category
	Name     string
}
