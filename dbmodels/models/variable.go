package models

type Variable struct {
	Id   string `gorm:"primaryKey"`
	Name string
	// ValueId string
	// Label   string
	// RunID   string
}
