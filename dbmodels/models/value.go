package models

type Value struct {
	Id      string `gorm:"primaryKey"`
	Label   string
	LabelId string
}
